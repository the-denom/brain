package keeper

import (
	"context"

	"github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Status(goCtx context.Context, req *types.QueryStatusRequest) (*types.QueryStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Must have a valid address
	memberAddress, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	// Get the member's account
	member, found := k.GetMember(ctx, memberAddress)
	if !found {
		// Member does not exist
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "member not found")
	}

	return &types.QueryStatusResponse{
		Member: &member,
	}, nil
}
