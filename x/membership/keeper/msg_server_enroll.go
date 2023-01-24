package keeper

import (
	"context"

	"github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k msgServer) Enroll(goCtx context.Context, msg *types.MsgEnroll) (*types.MsgEnrollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ak := k.accountKeeper

	// Must have a valid address
	enrollee, err := sdk.AccAddressFromBech32(msg.MemberAddress)
	if err != nil {
		return nil, err
	}

	// Must have a valid nickname length (if set)
	if len(msg.Nickname) > types.NicknameMaxLength {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "nickname too long")
	}

	// Must not already have a member account
	if k.IsMember(ctx, enrollee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "account has already been enrolled")
	}

	// Get or create a base account
	var baseAccount = ak.GetAccount(ctx, enrollee)
	if baseAccount == nil {
		// Create a base baseAccount
		baseAccount = ak.NewAccountWithAddress(ctx, enrollee)
		// Ensure account type is correct
		if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
		}
		// Save the base account
		ak.SetAccount(ctx, baseAccount)
	}

	// Create a member account
	memberAccount := types.NewMemberAccountWithDefaultMemberStatus(
		baseAccount.(*authtypes.BaseAccount),
		msg.Nickname,
	)
	// Save it to the store
	k.SetMember(ctx, enrollee, *memberAccount)

	// Publish events
	err = ctx.EventManager().EmitTypedEvents(
		// A new member was enrolled
		&types.EventMemberEnrolled{MemberAddress: enrollee.String()},
		// A member's citizenship status has changed
		&types.EventMemberStatusChanged{
			MemberAddress:  enrollee.String(),
			Status:         types.MembershipStatus_MemberElectorate,
			PreviousStatus: types.MembershipStatus_MemberStatusEmpty,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgEnrollResponse{}, nil
}
