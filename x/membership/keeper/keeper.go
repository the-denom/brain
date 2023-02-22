package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// KeeperI is the interface contract that x/membership's keeper implements
type KeeperI interface {
	GetMember(ctx sdk.Context, address sdk.AccAddress) (member types.Member, found bool)
	AppendMember(ctx sdk.Context, address sdk.AccAddress, newMember types.Member)
	IsMember(ctx sdk.Context, address sdk.AccAddress) bool
	GetMemberCount(ctx sdk.Context) uint64
	SetMemberCount(ctx sdk.Context, count uint64)
	UpdateMemberStatus(ctx sdk.Context, target sdk.AccAddress, s types.MembershipStatus)
}

//go:generate mockery --name=KeeperI --output=mocks
type Keeper struct {
	KeeperI
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	accountKeeper types.AccountKeeper

	paramstore paramtypes.Subspace
}

// Make sure Keeper implements the KeeperI interface
var _ KeeperI = &Keeper{}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper,

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

func (k Keeper) UpdateMemberStatus(ctx sdk.Context, target sdk.AccAddress, s types.MembershipStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemberKey))
	key := types.MakeMemberAddressKey(target)

	// Fetch the member and update status
	m, _ := k.GetMember(ctx, target)
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
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
