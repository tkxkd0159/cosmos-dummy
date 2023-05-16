package testutil

import (
	"checkers/x/checkers"
	"context"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"checkers/x/checkers/keeper"
	"checkers/x/checkers/types"
)

func CheckersKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	return checkersKeeperWithMocks(t, nil)
}

func checkersKeeperWithMocks(t testing.TB, bank *MockBankEscrowKeeper) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CheckersParams",
	)
	k := keeper.NewKeeper(
		bank,
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func SetupMsgServerWithOneGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *MockBankEscrowKeeper) {
	k, goCtx, ctrl, bankMock := SetupKeeperWithMocks(t)
	server := keeper.NewMsgServerImpl(k)
	server.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: Alice,
		Black:   Bob,
		Red:     Carol,
		Wager:   45,
	})
	return server, k, goCtx, ctrl, bankMock
}

func SetupKeeperWithMocks(t testing.TB) (keeper.Keeper, context.Context, *gomock.Controller, *MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := NewMockBankEscrowKeeper(ctrl)
	k, ctx := checkersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	goCtx := sdk.WrapSDKContext(ctx)
	return *k, goCtx, ctrl, bankMock
}
