package keeper

import (
	"foochain/x/foochain/types"
)

var _ types.QueryServer = Keeper{}
