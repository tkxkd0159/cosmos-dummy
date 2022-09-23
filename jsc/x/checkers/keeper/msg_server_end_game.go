package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"jsc/x/checkers/types"
)

func (k msgServer) EndGame(goCtx context.Context, msg *types.MsgEndGame) (*types.MsgEndGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEndGameResponse{}, nil
}
