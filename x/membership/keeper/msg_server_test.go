package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cdbo/brain/testutil/keeper"
	"github.com/cdbo/brain/x/membership/keeper"
	"github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MembershipKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
