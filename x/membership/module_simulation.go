package membership

import (
	"math/rand"

	"github.com/cdbo/brain/testutil/sample"
	membershipsimulation "github.com/cdbo/brain/x/membership/simulation"
	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = membershipsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgEnroll = "op_weight_msg_enroll"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnroll int = 100

	opWeightMsgUpdateStatus = "op_weight_msg_update_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	membershipGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&membershipGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgEnroll int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnroll, &weightMsgEnroll, nil,
		func(_ *rand.Rand) {
			weightMsgEnroll = defaultWeightMsgEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnroll,
		membershipsimulation.SimulateMsgEnroll(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateStatus, &weightMsgUpdateStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStatus = defaultWeightMsgUpdateStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStatus,
		membershipsimulation.SimulateMsgUpdateStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
