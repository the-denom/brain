package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"google.golang.org/api/option"
)

type tallyStats struct {
	// numVotes is the number of votes cast
	numVotes sdk.Dec
	// numMembers is the number of members who voted
	numMembers sdk.Dec
	// numEligibleMembers is the number of members who are eligible to vote
	numEligibleMembers sdk.Dec
	// numEligibleVotes is the number of votes we can include in the tally
	numEligibleVotes sdk.Dec
	// totalVotingPower is the total weight of all votes (including higher-weight individuals, called Guardians)
	totalVotingPower sdk.Dec
}

// Tally iterates over the votes and updates the tally of a proposal based on one-member, one vote
func (k Keeper) Tally(ctx sdk.Context, proposal govtypes.Proposal) (passes bool, burnDeposits bool, tallyResults govtypes.TallyResult) {
	results := make(map[govtypes.VoteOption]sdk.Dec)
	results[govtypes.OptionYes] = sdk.ZeroDec()
	results[govtypes.OptionNo] = sdk.ZeroDec()
	results[govtypes.OptionAbstain] = sdk.ZeroDec()
	results[govtypes.OptionNoWithVeto] = sdk.ZeroDec()

	stats := tallyStats{
		numVotes:           sdk.ZeroDec(),
		numMembers:         sdk.ZeroDec(),
		numEligibleMembers: sdk.ZeroDec(),
		numEligibleVotes:   sdk.ZeroDec(),
		totalVotingPower:   sdk.ZeroDec(),
	}

	k.IterateVotes(ctx, proposal.ProposalId, func(vote govtypes.Vote) bool {
		cl := k.Logger(ctx).With("voterAddress", vote.Voter, "proposal", proposal.ProposalId)

		stats.numVotes = stats.numVotes.Add(sdk.NewDec(1))

		// Validate this voter
		voterAddress := sdk.MustAccAddressFromBech32(vote.Voter)
		isMember, isEligibleToVote, err := k.ensureMembershipAndVotingPower(voterAddress)
		if err != nil {
			cl.Error("could not fetch member details: %s", err)
			return false
		}

		// The vote is ignored if the voter is not a member of the electorate
		if !isMember {
			cl.Error("voter is not a member of the electorate, ignored")
			return false
		}
		stats.numMembers = stats.numMembers.Add(sdk.NewDec(1))

		// The vote is ignored if the member is not eligible to vote
		if !isEligibleToVote {
			cl.Error("electorate member is not eligible to vote, ignored")
			return false
		}
		stats.numEligibleMembers = stats.numEligibleMembers.Add(sdk.NewDec(1))

		// The vote is ignored if the weighting is not 1 for a single option
		// (democratic voting does not support split votes)
		err = k.ensureValidVoteWeighting(vote.Options)
		if err != nil {
			cl.Error("vote has invalid weightings: %s, ignored", err)
			return false
		}

		// Get the vote power
		power := k.getMemberVotingPower(ctx, voterAddress)
		stats.totalVotingPower = stats.totalVotingPower.Add(power)

		// Update the results
		votersChoice := k.getVoterChoice(vote.Options)
		results[votersChoice] = results[votersChoice].Add(power)

		return false
	})

	return true, false, govtypes.TallyResult{}
}

func (k Keeper) ensureMembershipAndVotingPower(voter sdk.AccAddress) (isMember bool, isEligibleToVote bool, err error) {

}

func (k Keeper) ensureValidVoteWeighting(options govtypes.WeightedVoteOptions) error {
	totalWeight := sdk.NewDec(0)
	for _, option := range options {
		// Cannot have any other weighting besides 0 or 1
		if !option.Weight.IsZero() && !option.Weight.Equal(sdk.NewDec(1)) {
			return fmt.Errorf("option %s's weight is invalid: %s", option.Option, option.Weight.String())
		}
		totalWeight = totalWeight.Add(option.Weight)
	}

	// Cannot have a total weight of more than 1
	if !totalWeight.Equal(sdk.NewDec(1)) {
		return fmt.Errorf("vote is spoilt, total weighting of %s exceeds 1", option.Weight.String())
	}

	return nil
}

func (k Keeper) getVoterChoice(options govtypes.WeightedVoteOptions) govtypes.VoteOption {
	for _, option := range options {
		if option.Weight.Equal(sdk.NewDec(1)) {
			return option.Option
		}
	}
	return govtypes.OptionEmpty
}

func (k Keeper) getMemberVotingPower(ctx sdk.Context, address sdk.AccAddress) sdk.Dec {

}
