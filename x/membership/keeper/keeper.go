package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// KeeperI is the interface contract that x/membership's keeper implements
type KeeperI interface {
	// GetMember returns the member associated with the given address, and whether the member was found or not.
	GetMember(ctx sdk.Context, address sdk.AccAddress) (member types.Member, found bool)
	// AppendMember adds a new member to the store, associated with the given address.
	AppendMember(ctx sdk.Context, address sdk.AccAddress, newMember types.Member)
	// IsMember checks if a member with the given address exists in the store.
	IsMember(ctx sdk.Context, address sdk.AccAddress) bool
	// GetMemberCount returns the total number of members in the store.
	GetMemberCount(ctx sdk.Context) uint64
	// SetMemberCount sets the total number of members in the store.
	SetMemberCount(ctx sdk.Context, count uint64)
	// UpdateMemberStatus updates the membership status of a member associated with the given address.
	UpdateMemberStatus(ctx sdk.Context, target sdk.AccAddress, s types.MembershipStatus) error
	// IterateActiveProposalsQueue cycle through proposals that have ended their voting period
	IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes.Proposal) (stop bool))
	// DeleteDeposits deletes all the deposits on a specific proposal without refunding them
	DeleteDeposits(ctx sdk.Context, proposalID uint64)
	// RefundDeposits refunds and deletes all the deposits on a specific proposal
	RefundDeposits(ctx sdk.Context, proposalID uint64)
	// ExecuteProposalHandler executes the proposal's completion handler
	ExecuteProposalHandler(cacheCtx sdk.Context, proposalRoute string, content govtypes.Content) error
	// SetProposal writes the updated proposal to the store
	SetProposal(ctx sdk.Context, proposal govtypes.Proposal)
	// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
	RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	// AfterProposalVotingPeriodEnded - call hook if registered
	AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64)
	// Tally iterates over the votes and updates the tally of a proposal based on one-member, one vote
	Tally(ctx sdk.Context, proposal govtypes.Proposal) (passes bool, burnDeposits bool, tallyResults govtypes.TallyResult)
}

//go:generate mockery --name=KeeperI --output=mocks
type Keeper struct {
	KeeperI
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	accountKeeper types.AccountKeeper
	govKeeper     types.GovKeeper

	paramstore paramtypes.Subspace
}

// Make sure Keeper implements the KeeperI interface
var _ KeeperI = &Keeper{}

// NewKeeper creates a new instance of the Keeper struct.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper,
	govKeeper types.GovKeeper,

	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountKeeper,
		govKeeper:     govKeeper,

		paramstore: ps,
	}
}

func (k Keeper) GetMember(ctx sdk.Context, address sdk.AccAddress) (member types.Member, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	key := types.MakeMemberAddressKey(address)

	var b []byte
	if b = store.Get(key); b == nil {
		return member, false
	}

	k.cdc.Unmarshal(b, &member)
	return member, true
}

func (k Keeper) AppendMember(ctx sdk.Context, address sdk.AccAddress, newMember types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	memberCount := k.GetMemberCount(ctx)
	key := types.MakeMemberAddressKey(address)

	// Marshal and Set
	memberData := k.cdc.MustMarshal(&newMember)
	store.Set(key, memberData)

	// Bump member count
	k.SetMemberCount(ctx, memberCount+1)
}

func (k Keeper) IsMember(ctx sdk.Context, address sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	key := types.MakeMemberAddressKey(address)
	return store.Has(key)
}

func (k Keeper) GetMemberCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberCountKey))
	byteKey := []byte(types.MemberKey)
	bz := store.Get(byteKey)

	// Nil result means zero members
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetMemberCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberCountKey))
	bz := sdk.Uint64ToBigEndian(count)
	store.Set([]byte(types.MemberCountKey), bz)
}

func (k Keeper) UpdateMemberStatus(ctx sdk.Context, target sdk.AccAddress, s types.MembershipStatus) error {
	// Fetch the member
	m, found := k.GetMember(ctx, target)

	// Member must exist
	if !(found) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "member not found: %s", target.String())
	}

	// Must be a valid status transition
	if !m.Status.CanTransitionTo(s) {
		return sdkerrors.Wrapf(types.ErrMembershipStatusChangeNotAllowed, "transition %s is not allowed", s.DescribeTransition(s))
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	key := types.MakeMemberAddressKey(target)

	// Update the member's status
	m.Status = s

	// Marshal and Set
	memberData := k.cdc.MustMarshal(&m)
	store.Set(key, memberData)

	// Publish an update event
	ctx.EventManager().EmitTypedEvent(
		// A member's citizenship status has changed
		&types.EventMemberStatusChanged{
			MemberAddress: target.String(),
			// TODO: Change this
			Operator:       "",
			Status:         types.MembershipStatus_MemberElectorate,
			PreviousStatus: types.MembershipStatus_MemberStatusEmpty,
		},
	)

	return nil
}

// IterateActiveProposalsQueue cycle through proposals that have ended their voting period
func (k Keeper) IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes.Proposal) (stop bool)) {
	// Pass-through to Gov keeper
	k.govKeeper.IterateActiveProposalsQueue(ctx, endTime, cb)
}

// DeleteDeposits deletes all the deposits on a specific proposal without refunding them
func (k Keeper) DeleteDeposits(ctx sdk.Context, proposalID uint64) {
	k.govKeeper.DeleteDeposits(ctx, proposalID)
}

// RefundDeposits refunds and deletes all the deposits on a specific proposal
func (k Keeper) RefundDeposits(ctx sdk.Context, proposalID uint64) {
	k.govKeeper.RefundDeposits(ctx, proposalID)
}

// ExecuteProposalHandler executes the proposal's completion handler
func (k Keeper) ExecuteProposalHandler(cacheCtx sdk.Context, proposalRoute string, content govtypes.Content) error {
	handler := k.govKeeper.Router().GetRoute(proposalRoute)
	err := handler(cacheCtx, content)
	return err
}

// SetProposal writes the updated proposal to the store
func (k Keeper) SetProposal(ctx sdk.Context, proposal govtypes.Proposal) {
	k.govKeeper.SetProposal(ctx, proposal)
}

// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
func (k Keeper) RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	k.govKeeper.RemoveFromActiveProposalQueue(ctx, proposalID, endTime)
}

// AfterProposalVotingPeriodEnded - call hook if registered
func (k Keeper) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
	k.govKeeper.AfterProposalVotingPeriodEnded(ctx, proposalID)
}

// Tally iterates over the votes and updates the tally of a proposal based on one-member, one vote
func (k Keeper) Tally(ctx sdk.Context, proposal govtypes.Proposal) (passes bool, burnDeposits bool, tallyResults govtypes.TallyResult) {
	return true, false, govtypes.TallyResult{}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
