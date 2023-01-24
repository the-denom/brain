package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "membership"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_membership"

	MemberKey      = "Member/value/"
	MemberCountKey = "Member/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// MakeMemberAddressKey creates a composite key pointing to a Member's Denom account
func MakeMemberAddressKey(address sdk.AccAddress) []byte {
	// Combine MemberKey and address.ToString()
	combined := strings.Join([]string{MemberKey, address.String()}, "")
	return []byte(combined)
}
