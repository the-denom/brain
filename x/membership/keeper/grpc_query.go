package keeper

import (
	"github.com/cdbo/brain/x/membership/types"
)

var _ types.QueryServer = Keeper{}
