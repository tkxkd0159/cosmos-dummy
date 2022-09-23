package checkers

import (
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

type StoreGame struct {
	Creator string
	Index   string // The unique id that identifies this game
	Game    string // The serialized board
	Turn    string // "red" or "black"
	Red     string
	Black   string
}

func SerializeCreator(creator string) sdktypes.AccAddress {
	var (
		accAddr sdktypes.AccAddress
		err     error
	)

	accAddr, err = sdktypes.AccAddressFromBech32(creator)
	if err != nil {
		panic("Cannot convert bech32-style creator name to []byte")
	}
	return accAddr
}

func DeserializeCreator(c sdktypes.AccAddress) string {
	return c.String()
}
