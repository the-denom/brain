package keeper

import (
	"context"

	"github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Enroll(goCtx context.Context, msg *types.MsgEnroll) (*types.MsgEnrollResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	_ = sdk.UnwrapSDKContext(goCtx)
	/*
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

		// Must not be a member already
		if ak.HasAccount(ctx, enrollee) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", msg.MemberAddress)
		}

		// Create a base account
		baseAccount := ak.NewAccountWithAddress(ctx, enrollee)
		// Ensure account type is correct
		if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
		}
		// Create a member account
		memberAccount := types.NewMemberAccountWithDefaultMemberStatus(baseAccount.(*authtypes.BaseAccount), msg.Nickname)

		// Save it to the store
		ak.SetAccount(ctx, &memberAccount)
	*/

	return &types.MsgEnrollResponse{}, nil

}
