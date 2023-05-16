package keeper_test

import (
	"checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	"testing"

	sdktypes "github.com/cosmos/cosmos-sdk/types"

	"checkers/x/checkers/testutil"
)

func TestKeeper_CollectWager(t *testing.T) {
	keeper, goCtx, ctrl, escrow := testutil.SetupKeeperWithMocks(t)
	ctx := sdktypes.UnwrapSDKContext(goCtx)
	defer ctrl.Finish()
	escrow.ExpectPay(goCtx, alice, 45).Times(1).Return(nil)
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Wager:     45,
		MoveCount: 0,
	})
	require.Nil(t, err)
}
