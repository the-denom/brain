package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
