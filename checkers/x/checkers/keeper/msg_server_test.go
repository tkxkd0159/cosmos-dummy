package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	keepertest "checkers/testutil/keeper"
	"checkers/x/checkers"
	"checkers/x/checkers/keeper"
	"checkers/x/checkers/testutil"
	"checkers/x/checkers/types"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

func TestMsgSrvTestSuite(t *testing.T) {
	suite.Run(t, new(MsgSrvTestSuite))
}

type MsgSrvTestSuite struct {
	suite.Suite
	k      keeper.Keeper
	msgSrv types.MsgServer
	ctx    context.Context
}

func (suite *MsgSrvTestSuite) SetupTest() {
	k, ctx := keepertest.CheckersKeeper(suite.T())
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	suite.k = *k
	suite.msgSrv = keeper.NewMsgServerImpl(*k)
	suite.ctx = sdk.WrapSDKContext(ctx)
}

//func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
//	k, ctx := keepertest.CheckersKeeper(t)
//	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
//	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
//}

func (suite *MsgSrvTestSuite) TestCreateGame() {
	createResponse, err := suite.msgSrv.CreateGame(suite.ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(suite.T(), err)
	require.EqualValues(suite.T(), types.MsgCreateGameResponse{
		GameIndex: "1", // TODO: update with a proper value when updated
	}, *createResponse)
}

func (suite *MsgSrvTestSuite) TestCreate1GameHasSaved() {
	suite.msgSrv.CreateGame(suite.ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	systemInfo, found := suite.k.GetSystemInfo(sdk.UnwrapSDKContext(suite.ctx))
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.SystemInfo{
		NextId: 2,
	}, systemInfo)
	game1, found1 := suite.k.GetStoredGame(sdk.UnwrapSDKContext(suite.ctx), "1")
	require.True(suite.T(), found1)
	require.EqualValues(suite.T(), types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, game1)
}
