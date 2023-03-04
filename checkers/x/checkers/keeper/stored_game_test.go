package keeper_test

import (
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "checkers/testutil/keeper"
	"checkers/testutil/nullify"
	"checkers/x/checkers"
	"checkers/x/checkers/keeper"
	"checkers/x/checkers/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredGame(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredGame {
	items := make([]types.StoredGame, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredGame(ctx, items[i])
	}
	return items
}

func TestStoredGameGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredGameRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredGame(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredGame(ctx)),
	)
}

func TestForfeitUnplayed(t *testing.T) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	msgSrv := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	msgSrv.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	game1, found := k.GetStoredGame(ctx, "1")
	require.True(t, found)
	game1.Deadline = types.FormatDeadline(ctx.BlockTime().Add(time.Duration(-1)))
	k.SetStoredGame(ctx, game1)
	k.ForfeitExpiredGames(ctx)

	_, found = k.GetStoredGame(ctx, "1")
	require.False(t, found)

	systemInfo, found := k.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[1]
	require.EqualValues(t, sdk.StringEvent{
		Type: "checkers.checkers.EventForfeitGame",
		Attributes: []sdk.Attribute{
			{Key: "board", Value: "\"*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*\""},
			{Key: "game_index", Value: "\"1\""},
			{Key: "winner", Value: "\"*\""},
		},
	}, event)
}

func TestForfeitPlayedTwice(t *testing.T) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	msgSrv := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	msgSrv.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	msgSrv.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	msgSrv.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})

	game1, found := k.GetStoredGame(ctx, "1")
	require.True(t, found)
	oldDeadline := types.FormatDeadline(ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	k.SetStoredGame(ctx, game1)
	k.ForfeitExpiredGames(ctx)

	game1, found = k.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(2),
		BeforeIndex: "-1",
		AfterIndex:  "-1",
		Deadline:    oldDeadline,
		Winner:      "r",
	}, game1)

	systemInfo, found := k.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 3)
	event := events[1]
	require.EqualValues(t, sdk.StringEvent{
		Type: "checkers.checkers.EventForfeitGame",
		Attributes: []sdk.Attribute{
			{Key: "board", Value: "\"*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*\""},
			{Key: "game_index", Value: "\"1\""},
			{Key: "winner", Value: "\"r\""},
		},
	}, event)
}
