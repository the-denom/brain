package keeper_test

import (
	"testing"

	testkeeper "github.com/cdbo/brain/testutil/keeper"
	"github.com/cdbo/brain/x/membership/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MembershipKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
