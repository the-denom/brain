package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateStatus = "update_status"

var _ sdk.Msg = &MsgUpdateStatus{}

func NewMsgUpdateStatus(creator string, address string, status MembershipStatus) *MsgUpdateStatus {
	// Make sure status is

	return &MsgUpdateStatus{
		Creator: creator,
		Address: address,
		Status:  status,
	}
}

func (msg *MsgUpdateStatus) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStatus) Type() string {
	return TypeMsgUpdateStatus
}

func (msg *MsgUpdateStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStatus) ValidateBasic() error {
	// Creator address must be valid
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// Target address must be valid
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	// Status must be valid
	if !msg.Status.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidMembershipStatus, "status must be one of the following: %s", GetAllShortFormMembershipStatusesAsString())
	}

	return nil
}
