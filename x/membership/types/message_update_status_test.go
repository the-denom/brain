package types

import (
	"testing"

	"github.com/cdbo/brain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateStatus_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateStatus
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgUpdateStatus{
				Creator: "invalid_address",
				Address: sample.AccAddress(),
				Status:  MembershipStatus_MemberExpulsed,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid creator address",
			msg: MsgUpdateStatus{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				Status:  MembershipStatus_MemberExpulsed,
			},
		},
		{
			name: "invalid target address",
			msg: MsgUpdateStatus{
				Creator: sample.AccAddress(),
				Address: "invalid_address",
				Status:  MembershipStatus_MemberExpulsed,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid creator address",
			msg: MsgUpdateStatus{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				Status:  MembershipStatus_MemberExpulsed,
			},
		},
		{
			name: "invalid target status: empty",
			msg: MsgUpdateStatus{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				Status:  MembershipStatus_MemberStatusEmpty,
			},
			err: ErrInvalidMembershipStatus,
		},
		{
			name: "invalid target status: out of range",
			msg: MsgUpdateStatus{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				Status:  MembershipStatus(100),
			},
			err: ErrInvalidMembershipStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
