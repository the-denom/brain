package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NicknameMaxLength is the maximum number of characters allowed for
// a member's nickname
const NicknameMaxLength = 30

// NewMemberAccountWithDefaultMemberStatus creats a new member account with a
// default member status of Electorate.
func NewMemberAccountWithDefaultMemberStatus(baseAccount *authtypes.BaseAccount, nickname string) *Member {
	acc := &Member{
		BaseAccount: baseAccount,
		Status:      MembershipStatus_MemberElectorate,
		Nickname:    nickname,
	}
	return acc
}
