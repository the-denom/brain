package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
//
//go:generate mockery --name=AccountKeeper --output=mocks
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI

	// Check if an account exists in the store.
	HasAccount(sdk.Context, sdk.AccAddress) bool
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI
	// Set an account in the store.
	SetAccount(sdk.Context, types.AccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type GovKeeper interface {
	// IterateActiveProposalsQueue iterates over the proposals in the active proposal queue
	// and performs a callback function
	IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes.Proposal) (stop bool))
	// DeleteDeposits deletes all the deposits on a specific proposal without refunding them
	DeleteDeposits(ctx sdk.Context, proposalID uint64)
	// RefundDeposits refunds and deletes all the deposits on a specific proposal
	RefundDeposits(ctx sdk.Context, proposalID uint64)
	// Router returns the gov Keeper's Router
	Router() govtypes.Router
	// SetProposal set a proposal to store
	SetProposal(ctx sdk.Context, proposal govtypes.Proposal)
	// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
	RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	// AfterProposalVotingPeriodEnded - call hook if registered
	AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64)
}
