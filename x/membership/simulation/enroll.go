package simulation

import (
	"math/rand"

	"github.com/cdbo/brain/x/membership/keeper"
	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgEnroll(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.MembershipKeeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEnroll{
			MemberAddress: simAccount.Address.String(),
		}

		// TODO: Handling the Enroll simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Enroll simulation not implemented"), nil, nil
	}
}
