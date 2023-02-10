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

	suite.msgSrv.CreateGame(suite.ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
}

func (suite *MsgSrvTestSuite) TestCreateGame() {
	createResponse, err := suite.msgSrv.CreateGame(suite.ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(suite.T(), err)
	require.EqualValues(suite.T(), types.MsgCreateGameResponse{
		GameIndex: "2", // TODO: update with a proper value when updated
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
		NextId:        3,
		FifoHeadIndex: "1",
		FifoTailIndex: "2",
	}, systemInfo)
	game1, found1 := suite.k.GetStoredGame(sdk.UnwrapSDKContext(suite.ctx), "1")
	require.True(suite.T(), found1)
	require.EqualValues(suite.T(), types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
}

func (suite *MsgSrvTestSuite) TestCreate1GameEmitted() {
	ctx := sdk.UnwrapSDKContext(suite.ctx)
	suite.msgSrv.CreateGame(suite.ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.NotNil(suite.T(), ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(suite.T(), events, 1)
	require.EqualValues(suite.T(), sdk.StringEvents([]sdk.StringEvent{{
		Type: "checkers.checkers.EventCreateGame",
		Attributes: []sdk.Attribute{
			{Key: "black", Value: "\"cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g\""}, // by suite
			{Key: "creator", Value: "\"cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3\""},
			{Key: "game_index", Value: "\"1\""},
			{Key: "red", Value: "\"cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7\""},
			{Key: "black", Value: "\"cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g\""},
			{Key: "creator", Value: "\"cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3\""}, // by func
			{Key: "game_index", Value: "\"2\""},
			{Key: "red", Value: "\"cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7\""},
		}},
	}).Flatten(), events)
}

func (suite *MsgSrvTestSuite) TestPlayMove() {
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	require.Nil(suite.T(), err)
	require.EqualValues(suite.T(), types.MsgPlayMoveResponse{
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *MsgSrvTestSuite) TestPlayMoveGameNotFound() {
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "2",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	require.Nil(suite.T(), playMoveResponse)
	require.Equal(suite.T(), "2: game by id not found", err.Error())
}

func (suite *MsgSrvTestSuite) TestPlayMoveSameBlackRed() {
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	require.Nil(suite.T(), err)
	require.EqualValues(suite.T(), types.MsgPlayMoveResponse{
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *MsgSrvTestSuite) TestPlayMoveSavedGame() {
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	ctx := sdk.UnwrapSDKContext(suite.ctx)
	systemInfo, found := suite.k.GetSystemInfo(ctx)
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "1",
		FifoTailIndex: "1",
	}, systemInfo)
	game1, found := suite.k.GetStoredGame(ctx, "1")
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       bob,
		Red:         carol,
		MoveCount:   1,
		BeforeIndex: "-1",
		AfterIndex:  "-1",
	}, game1)
}

func (suite *MsgSrvTestSuite) TestPlayMoveNotPlayer() {
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   alice,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	require.Nil(suite.T(), playMoveResponse)
	require.Equal(suite.T(), alice+": message creator is not a player", err.Error())
}

func (suite *MsgSrvTestSuite) TestPlayMoveCannotParseGame() {
	ctx := sdk.UnwrapSDKContext(suite.ctx)
	storedGame, _ := suite.k.GetStoredGame(ctx, "1")
	storedGame.Board = "not a board"
	suite.k.SetStoredGame(ctx, storedGame)
	defer func() {
		r := recover()
		require.NotNil(suite.T(), r, "The code did not panic")
		require.Equal(suite.T(), r, "game cannot be parsed: invalid board string: not a board")
	}()
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
}

func (suite *MsgSrvTestSuite) TestPlayMoveWrongOutOfTurn() {
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	require.Nil(suite.T(), playMoveResponse)
	require.Equal(suite.T(), "{red}: player tried to play out of turn", err.Error())
}

func (suite *MsgSrvTestSuite) TestPlayMoveWrongPieceAtDestination() {
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     0,
		ToX:       0,
		ToY:       1,
	})
	require.Nil(suite.T(), playMoveResponse)
	require.Equal(suite.T(), "Already piece at destination position: {0 1}: wrong move", err.Error())
}

func (suite *MsgSrvTestSuite) TestPlayMove2() {
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	require.Nil(suite.T(), err)
	require.EqualValues(suite.T(), types.MsgPlayMoveResponse{
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *MsgSrvTestSuite) TestPlayMove2SavedGame() {
	ctx := sdk.UnwrapSDKContext(suite.ctx)
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	systemInfo, found := suite.k.GetSystemInfo(ctx)
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "1",
		FifoTailIndex: "1",
	}, systemInfo)
	game1, found := suite.k.GetStoredGame(ctx, "1")
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   2,
		BeforeIndex: "-1",
		AfterIndex:  "-1",
	}, game1)
}

func (suite *MsgSrvTestSuite) TestPlayMove3() {
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	playMoveResponse, err := suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     2,
		FromY:     3,
		ToX:       0,
		ToY:       5,
	})
	require.Nil(suite.T(), err)
	require.EqualValues(suite.T(), types.MsgPlayMoveResponse{
		CapturedX: 1,
		CapturedY: 4,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *MsgSrvTestSuite) TestPlayMove3SavedGame() {
	ctx := sdk.UnwrapSDKContext(suite.ctx)
	suite.msgSrv.PlayMove(ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgSrv.PlayMove(ctx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	suite.msgSrv.PlayMove(ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     2,
		FromY:     3,
		ToX:       0,
		ToY:       5,
	})
	systemInfo, found := suite.k.GetSystemInfo(ctx)
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "1",
		FifoTailIndex: "1",
	}, systemInfo)
	game1, found := suite.k.GetStoredGame(ctx, "1")
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|********|********|b*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       bob,
		Red:         carol,
		MoveCount:   3,
		BeforeIndex: "-1",
		AfterIndex:  "-1",
	}, game1)
}

func (suite *MsgSrvTestSuite) TestPlayMove2Emitted() {
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	ctx := sdk.UnwrapSDKContext(suite.ctx)
	require.NotNil(suite.T(), ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(suite.T(), events, 2)
	event := events[1]
	require.Equal(suite.T(), "checkers.checkers.EventMove", event.Type)
	require.EqualValues(suite.T(), []sdk.Attribute{
		{Key: "captured_x", Value: "\"-1\""},
		{Key: "captured_y", Value: "\"-1\""},
		{Key: "creator", Value: "\"cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7\""},
		{Key: "game_index", Value: "\"1\""},
		{Key: "winner", Value: "\"*\""},
	}, event.Attributes[5:])
}

func (suite *MsgSrvTestSuite) TestRejectGameByRedOneMoveRemovedGame() {
	suite.msgSrv.PlayMove(suite.ctx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgSrv.RejectGame(suite.ctx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	systemInfo, found := suite.k.GetSystemInfo(sdk.UnwrapSDKContext(suite.ctx))
	require.True(suite.T(), found)
	require.EqualValues(suite.T(), types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	_, found = suite.k.GetStoredGame(sdk.UnwrapSDKContext(suite.ctx), "1")
	require.False(suite.T(), found)
}

func (suite *MsgSrvTestSuite) TestCreate3GamesHasSavedFifo() {
	msgSrvr := suite.msgSrv
	context := suite.ctx
	keeper := suite.k
	t := suite.T()

	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	systemInfo2, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        4,
		FifoHeadIndex: "1",
		FifoTailIndex: "3",
	}, systemInfo2)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found := keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "3",
	}, game2)

	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	systemInfo3, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        5,
		FifoHeadIndex: "1",
		FifoTailIndex: "4",
	}, systemInfo3)
	game1, found = keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found = keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "3",
	}, game2)
	game3, found := keeper.GetStoredGame(ctx, "3")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       carol,
		Red:         alice,
		MoveCount:   uint64(0),
		BeforeIndex: "2",
		AfterIndex:  "4",
	}, game3)
}

func (suite *MsgSrvTestSuite) TestPlayMove2Games2MovesHasSavedFifo() {
	msgSrvr := suite.msgSrv
	context := suite.ctx
	keeper := suite.k
	t := suite.T()

	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	msgSrvr.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	msgSrvr.PlayMove(context, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "2",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	systemInfo1, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        3,
		FifoHeadIndex: "1",
		FifoTailIndex: "2",
	}, systemInfo1)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(1),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found := keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       carol,
		Red:         alice,
		MoveCount:   uint64(1),
		BeforeIndex: "1",
		AfterIndex:  "-1",
	}, game2)
}

func (suite *MsgSrvTestSuite) TestRejectMiddleGameHasSavedFifo() {
	msgSrvr := suite.msgSrv
	context := suite.ctx
	keeper := suite.k
	t := suite.T()

	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	msgSrvr.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "2",
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        4,
		FifoHeadIndex: "1",
		FifoTailIndex: "3",
	}, systemInfo)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "3",
	}, game1)
	game3, found := keeper.GetStoredGame(ctx, "3")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       alice,
		Red:         bob,
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "-1",
	}, game3)
}
