package keeper

import (
	"github.com/cdbo/brain/x/membership/types"
)

type msgServer struct {
	MembershipKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper MembershipKeeper) types.MsgServer {
	return &msgServer{MembershipKeeper: keeper}
}

var _ types.MsgServer = msgServer{}
