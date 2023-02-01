package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/membership module sentinel errors
var (
	ErrInvalidMembershipStatus          = sdkerrors.Register(ModuleName, 2, "invalid membership status")
	ErrMemberNotFound                   = sdkerrors.Register(ModuleName, 3, "member not found")
	ErrMembershipStatusChangeNotAllowed = sdkerrors.Register(ModuleName, 4, "membership status change not allowed")
)
