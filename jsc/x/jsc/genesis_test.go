package jsc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "jsc/testutil/keeper"
	"jsc/testutil/nullify"
	"jsc/x/jsc"
	"jsc/x/jsc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JscKeeper(t)
	jsc.InitGenesis(ctx, *k, genesisState)
	got := jsc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
