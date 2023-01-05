package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnroll = "enroll"

var _ sdk.Msg = &MsgEnroll{}

func NewMsgEnroll(creator string, nickname string) *MsgEnroll {
	return &MsgEnroll{
		Creator:  creator,
		Nickname: nickname,
	}
}

func (msg *MsgEnroll) Route() string {
	return RouterKey
}

func (msg *MsgEnroll) Type() string {
	return TypeMsgEnroll
}

func (msg *MsgEnroll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
