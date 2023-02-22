package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cdbo/brain/testutil/keeper"
	"github.com/cdbo/brain/testutil/sample"
	"github.com/cdbo/brain/x/membership/keeper"
	"github.com/cdbo/brain/x/membership/types"
	"github.com/cdbo/brain/x/membership/types/mocks"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type MsgServerTestSuite struct {
	suite.Suite
	t       *testing.T
	ctxS    sdk.Context
	ctx     context.Context
	msgSrvr types.MsgServer
	ak      types.AccountKeeper
}

func (s *MsgServerTestSuite) SetupTest() {
	s.ak = mocks.NewAccountKeeper(s.t)
	k, ctxS := keepertest.NewMembershipKeeperWithAccountKeeper(s.t, mocks.NewAccountKeeper(s.t))
	s.msgSrvr = keeper.NewMsgServerImpl(*k)
	s.ctxS = ctxS
	s.ctx = sdk.WrapSDKContext(ctxS)
}

func (s *MsgServerTestSuite) TestUpdateStatus_InvalidTargetAddress() {
	msg := types.NewMsgUpdateStatus(sample.AccAddress(), "invalid target", types.MembershipStatus_MemberElectorate)
	res, err := s.msgSrvr.UpdateStatus(s.ctx, msg)
	s.Require().Error(err)
	s.Require().Nil(res)
}

func (s *MsgServerTestSuite) TestUpdateStatus_MemberDoesNotExist() {
	msg := types.NewMsgUpdateStatus(sample.AccAddress(), sample.AccAddress(), types.MembershipStatus_MemberElectorate)

	// TODO: figure out how to mock the call to GetMember

	res, err := s.msgSrvr.UpdateStatus(s.ctx, msg)
	s.Require().Error(err)
	s.Require().Nil(res)
}

func TestMsgServerTestSuite(t *testing.T) {
	s := new(MsgServerTestSuite)
	s.t = t
	suite.Run(t, s)
}
