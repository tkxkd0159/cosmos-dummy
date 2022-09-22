package checkers

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	BoardDim = 8
	RED      = "red"
	BLACK    = "black"
	RowSep   = "|"
)

type Player struct {
	Color string
}

type Piece struct {
	Player Player
	King   bool
}

var PieceStrings = map[Player]string{
	RedPlayer:   "r",
	BlackPlayer: "b",
	NoPlayer:    "*",
}

var NoPiece = Piece{NoPlayer, false}

var StringPieces = map[string]Piece{
	"r": {RedPlayer, false},
	"b": {BlackPlayer, false},
	"R": {RedPlayer, true},
	"B": {BlackPlayer, true},
	"*": NoPiece,
}

type Pos struct {
	X int
	Y int
}

var NoPos = Pos{-1, -1}

var BlackPlayer = Player{BLACK}
var RedPlayer = Player{RED}
var NoPlayer = Player{
	Color: "NO_PLAYER",
}

var Players = map[string]Player{
	RED:   RedPlayer,
	BLACK: BlackPlayer,
}

var Opponents = map[Player]Player{
	BlackPlayer: RedPlayer,
	RedPlayer:   BlackPlayer,
}

var Usable = map[Pos]bool{}
var Moves = map[Player]map[Pos]map[Pos]bool{}
var Jumps = map[Player]map[Pos]map[Pos]Pos{}
var KingMoves = map[Pos]map[Pos]bool{}
var KingJumps = map[Pos]map[Pos]Pos{}

func Capture(src, dst Pos) Pos {
	return Pos{(src.X + dst.X) / 2, (src.Y + dst.Y) / 2}
}

// init is an initialization logic for state machine before receive any requests from Tendermint Core
func init() {

	// Initialize usable spaces
	for y := 0; y < BoardDim; y++ {
		for x := (y + 1) % 2; x < BoardDim; x += 2 {
			Usable[Pos{X: x, Y: y}] = true
		}
	}

	// Initialize deep maps
	for _, p := range Players {
		Moves[p] = map[Pos]map[Pos]bool{}
		Jumps[p] = map[Pos]map[Pos]Pos{}
	}

	// Compute possible moves, jumps and captures
	for pos := range Usable {
		KingMoves[pos] = map[Pos]bool{}
		KingJumps[pos] = map[Pos]Pos{}
		var directions = []int{1, -1}
		for i, player := range []Player{BlackPlayer, RedPlayer} {
			Moves[player][pos] = map[Pos]bool{}
			Jumps[player][pos] = map[Pos]Pos{}
			movOff := 1
			jmpOff := 2
			for _, direction := range directions {
				mov := Pos{pos.X + (movOff * direction), pos.Y + (movOff * directions[i])}
				if Usable[mov] {
					Moves[player][pos][mov] = true
					KingMoves[pos][mov] = true
				}
				jmp := Pos{pos.X + (jmpOff * direction), pos.Y + (jmpOff * directions[i])}
				if Usable[jmp] {
					capturePos := Capture(pos, jmp)
					Jumps[player][pos][jmp] = capturePos
					KingJumps[pos][jmp] = capturePos
				}
			}
		}
	}
}

type Game struct {
	Pieces map[Pos]Piece
	Turn   Player
}

func New() *Game {
	pieces := make(map[Pos]Piece)
	game := &Game{pieces, BlackPlayer}
	game.addInitialPieces()
	return game
}

func (game *Game) addInitialPieces() {
	for pos := range Usable {
		if pos.Y >= 0 && pos.Y < 3 {
			game.Pieces[pos] = Piece{BlackPlayer, false}
		}
		if pos.Y >= BoardDim-3 && pos.Y < BoardDim {
			game.Pieces[pos] = Piece{RedPlayer, false}
		}
	}
}

func (game *Game) PieceAt(pos Pos) bool {
	_, ok := game.Pieces[pos]
	return ok
}

func (game *Game) TurnIs(player Player) bool {
	return game.Turn == player
}

func (game *Game) Winner() Player {
	redCount := 0
	blackCount := 0
	for _, piece := range game.Pieces {
		switch {
		case piece.Player == BlackPlayer:
			blackCount += 1
		case piece.Player == RedPlayer:
			redCount += 1
		}
	}
	if blackCount > 0 && redCount <= 0 {
		return BlackPlayer
	} else if redCount > 0 && blackCount <= 0 {
		return RedPlayer
	}
	return NoPlayer
}

// ValidMove acts like a CheckTx
func (game *Game) ValidMove(src, dst Pos) bool {
	if !game.PieceAt(src) || game.PieceAt(dst) {
		return false
	}
	piece := game.Pieces[src]
	if (!piece.King && Moves[piece.Player][src][dst]) || (piece.King && KingMoves[src][dst]) {
		return !game.playerHasJump(piece.Player)
	}
	return game.ValidJump(src, dst)
}

func (game *Game) ValidJump(src, dst Pos) bool {
	if !game.PieceAt(src) || game.PieceAt(dst) {
		return false
	}
	piece := game.Pieces[src]
	if !piece.King {
		capLoc, jumpOk := Jumps[piece.Player][src][dst]
		return jumpOk && game.PieceAt(capLoc) && game.Pieces[capLoc].Player == Opponents[piece.Player]
	} else {
		capLoc, kingJumpOk := KingJumps[src][dst]
		return kingJumpOk && game.PieceAt(capLoc) && game.Pieces[capLoc].Player == Opponents[piece.Player]
	}
}

func (game *Game) kingPiece(dst Pos) {
	if !game.PieceAt(dst) {
		return
	}
	piece := game.Pieces[dst]
	if (dst.Y == 0 && piece.Player == RedPlayer) ||
		(dst.Y == BoardDim-1 && piece.Player == BlackPlayer) {
		piece.King = true
		game.Pieces[dst] = piece
	}
}

func (game *Game) updateTurn(dst Pos, jumped bool) {
	opponent := Opponents[game.Turn]
	if (!jumped || !game.jumpPossibleFrom(dst)) && game.playerHasMove(opponent) {
		game.Turn = opponent
	}
}

func (game *Game) jumpPossibleFrom(src Pos) bool {
	if !game.PieceAt(src) {
		return false
	}
	piece := game.Pieces[src]
	if !piece.King {
		// enumerate all player jumps and return true if one is valid
		for dst := range Jumps[piece.Player][src] {
			if game.ValidJump(src, dst) {
				return true
			}
		}
	} else {
		// enumerate all king jumps and return true if one is valid
		for dst := range KingJumps[src] {
			if game.ValidJump(src, dst) {
				return true
			}
		}
	}
	return false
}

func (game *Game) movePossibleFrom(src Pos) bool {
	if !game.PieceAt(src) {
		return false
	}
	piece := game.Pieces[src]
	if !piece.King {
		for dst := range Moves[piece.Player][src] {
			if game.ValidMove(src, dst) {
				return true
			}
		}
	} else {
		for dst := range KingMoves[src] {
			if game.ValidMove(src, dst) {
				return true
			}
		}
	}
	return false
}

func (game *Game) playerHasMove(player Player) bool {
	for loc, piece := range game.Pieces {
		if piece.Player == player && (game.movePossibleFrom(loc) || game.jumpPossibleFrom(loc)) {
			return true
		}
	}
	return false
}

func (game *Game) playerHasJump(player Player) bool {
	for loc, piece := range game.Pieces {
		if piece.Player == player && game.jumpPossibleFrom(loc) {
			return true
		}
	}
	return false
}

// Move acts like a DeliverTx
func (game *Game) Move(src, dst Pos) (captured Pos, err error) {
	captured = NoPos
	err = nil
	if !game.PieceAt(src) {
		return NoPos, errors.New(fmt.Sprintf("No piece at source position: %v", src))
	}
	if game.PieceAt(dst) {
		return NoPos, errors.New(fmt.Sprintf("Already piece at destination position: %v", dst))
	}
	if !game.TurnIs(game.Pieces[src].Player) {
		return NoPos, errors.New(fmt.Sprintf("Not %v's turn", game.Pieces[src].Player))
	}
	if !game.ValidMove(src, dst) {
		return NoPos, errors.New(fmt.Sprintf("Invalid move: %v to %v", src, dst))
	}
	if game.ValidJump(src, dst) {
		game.Pieces[dst] = game.Pieces[src]
		delete(game.Pieces, src)
		captured = Capture(src, dst)
		delete(game.Pieces, captured)
	} else {
		game.Pieces[dst] = game.Pieces[src]
		delete(game.Pieces, src)
	}
	game.updateTurn(dst, captured != NoPos)
	game.kingPiece(dst)
	return
}

// String is querying the game state. Serialize the board without any effort.
// It does not save player's Turn
func (game *Game) String() string {
	var buf bytes.Buffer
	for y := 0; y < BoardDim; y++ {
		for x := 0; x < BoardDim; x++ {
			pos := Pos{x, y}
			if game.PieceAt(pos) {
				piece := game.Pieces[pos]
				val := PieceStrings[piece.Player]
				if piece.King {
					val = strings.ToUpper(val)
				}
				buf.WriteString(val)
			} else {
				buf.WriteString(PieceStrings[NoPlayer])
			}
		}
		if y < (BoardDim - 1) {
			buf.WriteString(RowSep)
		}
	}
	return buf.String()
}

func ParsePiece(s string) (Piece, bool) {
	piece, ok := StringPieces[s]
	return piece, ok
}

// Parse re-instantiate the board state out of its serialized form
func Parse(s string) (*Game, error) {
	if len(s) != BoardDim*BoardDim+(BoardDim-1) {
		return nil, errors.New(fmt.Sprintf("invalid board string: %v", s))
	}
	pieces := make(map[Pos]Piece)
	result := &Game{pieces, BlackPlayer}
	for y, row := range strings.Split(s, RowSep) {
		for x, c := range strings.Split(row, "") {
			if x >= BoardDim || y >= BoardDim {
				return nil, errors.New(fmt.Sprintf("invalid board, piece out of bounds: %v, %v", x, y))
			}
			if piece, ok := ParsePiece(c); !ok {
				return nil, errors.New(fmt.Sprintf("invalid board, invalid piece at %v, %v", x, y))
			} else if piece != NoPiece {
				result.Pieces[Pos{x, y}] = piece
			}
		}
	}
	return result, nil
}
