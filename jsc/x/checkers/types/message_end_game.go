package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEndGame = "end_game"

var _ sdk.Msg = &MsgEndGame{}

func NewMsgEndGame(creator string, sig string) *MsgEndGame {
	return &MsgEndGame{
		Creator: creator,
		Sig:     sig,
	}
}

func (msg *MsgEndGame) Route() string {
	return RouterKey
}

func (msg *MsgEndGame) Type() string {
	return TypeMsgEndGame
}

func (msg *MsgEndGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEndGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEndGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
