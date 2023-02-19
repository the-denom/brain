package keeper

import (
	"context"

	"github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateStatus(goCtx context.Context, msg *types.MsgUpdateStatus) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Target member must have a valid address
	target, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Get the target member account
	targetAccount, found := k.GetMember(ctx, target)

	// Target member must exist
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "no such member: %s", target.String())
	}

	// Must be a valid status transition
	if !targetAccount.Status.CanTransitionTo(msg.Status) {
		return nil, sdkerrors.Wrapf(types.ErrMembershipStatusChangeNotAllowed, "transition %s is not allowed", targetAccount.Status.DescribeTransition(msg.Status))
	}

	// Execute the status update
	k.UpdateMemberStatus(ctx, targetAccount.GetAddress(), msg.Status)

	return &types.MsgUpdateStatusResponse{}, nil
}
