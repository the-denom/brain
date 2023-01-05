package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnroll = "enroll"

var _ sdk.Msg = &MsgEnroll{}

func NewMsgEnroll(memberAddress string, nickname string) *MsgEnroll {
	return &MsgEnroll{
		MemberAddress: memberAddress,
		Nickname:      nickname,
	}
}

func (msg *MsgEnroll) Route() string {
	return RouterKey
}

func (msg *MsgEnroll) Type() string {
	return TypeMsgEnroll
}

func (msg *MsgEnroll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.MemberAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnroll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnroll) ValidateBasic() error {
	// Member address must be valid
	if _, err := sdk.AccAddressFromBech32(msg.MemberAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid member address: %s", err)
	}
	// Nickname is optional but can only be a max of 30 characters long
	if len(msg.Nickname) > NicknameMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("member nickname longer than the %d character limit", NicknameMaxLength)
	}
	return nil
}
