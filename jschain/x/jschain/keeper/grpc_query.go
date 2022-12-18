package keeper

import (
	"jschain/x/jschain/types"
)

var _ types.QueryServer = Keeper{}
