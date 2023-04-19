package membership

import (
	"fmt"
	"time"

	"github.com/cdbo/brain/x/membership/keeper"
	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	logger := keeper.Logger(ctx)

	// iterate through proposals that have ended their voting period
	keeper.IterateActiveProposalsQueue(ctx, ctx.BlockHeader().Time, func(proposal govtypes.Proposal) (stop bool) {
		var tagValue, logMsg string

		// Calculate the tally
		passes, burnDeposits, tallyResults := keeper.Tally(ctx, proposal)

		if burnDeposits {
			keeper.DeleteDeposits(ctx, proposal.ProposalId)
		} else {
			keeper.RefundDeposits(ctx, proposal.ProposalId)
		}

		if passes {
			cacheCtx, writeCache := ctx.CacheContext()

			err := keeper.ExecuteProposalHandler(cacheCtx, proposal.ProposalRoute(), proposal.GetContent())
			if err == nil {
				proposal.Status = govtypes.StatusPassed
				tagValue = govtypes.AttributeValueProposalPassed
				logMsg = "passed"

				// The cached context is created with a new EventManager. However, since
				// the proposal handler execution was successful, we want to track/keep
				// any events emitted, so we re-emit to "merge" the events into the
				// original Context's EventManager.
				ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

				// write state to the underlying multi-store
				writeCache()
			} else {
				proposal.Status = govtypes.StatusFailed
				tagValue = govtypes.AttributeValueProposalFailed
				logMsg = fmt.Sprintf("passed, but failed on execution: %s", err)
			}
		} else {
			proposal.Status = govtypes.StatusRejected
			tagValue = govtypes.AttributeValueProposalRejected
			logMsg = "rejected"
		}

		proposal.FinalTallyResult = tallyResults

		// Write the updated proposal to the store
		keeper.SetProposal(ctx, proposal)
		// Remove it from the active queue
		keeper.RemoveFromActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)

		// Hook: when proposal become active
		keeper.AfterProposalVotingPeriodEnded(ctx, proposal.ProposalId)

		logger.Info(
			"proposal tallied",
			"proposal", proposal.ProposalId,
			"title", proposal.GetTitle(),
			"result", logMsg,
		)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				govtypes.EventTypeActiveProposal,
				sdk.NewAttribute(govtypes.AttributeKeyProposalID, fmt.Sprintf("%d", proposal.ProposalId)),
				sdk.NewAttribute(govtypes.AttributeKeyProposalResult, tagValue),
			),
		)

		// Don't stop looping
		return false
	})
}
