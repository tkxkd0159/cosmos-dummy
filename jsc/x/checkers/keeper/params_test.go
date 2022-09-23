package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "jsc/testutil/keeper"
	"jsc/x/checkers/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CheckersKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
