package checkers

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatalf("checkers object should be non-nil")
	}
	if c.Pieces == nil {
		t.Errorf("expected pieces to be non-nil")
	}
	if c.Winner() != NoPlayer {
		t.Errorf("expected winner to be no-player")
	}
	if !c.TurnIs(BlackPlayer) {
		t.Errorf("expected starting turn to be black")
	}
	// Confirm pieces are all at starting positions
	for pos := range Usable {
		if (pos.Y >= 0 && pos.Y < 3) && (!c.PieceAt(pos) || c.Pieces[pos].Player != BlackPlayer) {
			t.Errorf("expected black piece at %v", pos)
		}
		if (pos.Y >= BoardDim-3 && pos.Y < BoardDim) && (!c.PieceAt(pos) || c.Pieces[pos].Player != RedPlayer) {
			t.Errorf("expected red piece at %v", pos)
		}
	}
}

func TestWinner(t *testing.T) {
	// Test no initial winner (assumes correct game setup)
	game := New()
	if game.Winner() != NoPlayer {
		t.Errorf("expected no initial game winner")
	}
	// Test winner is unaffected after removing a piece
	for pos := range game.Pieces {
		delete(game.Pieces, pos)
		break // Only delete one
	}
	if game.Winner() != NoPlayer {
		t.Errorf("expected no winner after removal of a piece")
	}
	// Test removing all red makes black win
	for pos, piece := range game.Pieces {
		if piece.Player == RedPlayer {
			delete(game.Pieces, pos)
		}
	}
	if game.Winner() != BlackPlayer {
		t.Errorf("expected black to win with no red pieces")
	}
	// Test removing all pieces yields no winner
	for pos := range game.Pieces {
		delete(game.Pieces, pos)
	}
	if game.Winner() != NoPlayer {
		t.Errorf("expected no winner with an empty board")
	}
	// Try again removing all black
	game = New()
	for pos, piece := range game.Pieces {
		if piece.Player == BlackPlayer {
			delete(game.Pieces, pos)
		}
	}
	if game.Winner() != RedPlayer {
		t.Errorf("expected red to win with no black pieces")
	}
}

func TestValidMove(t *testing.T) {
	game := New()
	src := Pos{3, 2}
	dst := Pos{4, 3}
	if !game.ValidMove(src, dst) {
		t.Errorf("expected %v to %v to be valid", src, dst)
	}
	dst = Pos{3, 3}
	if game.ValidMove(src, dst) {
		t.Errorf("expected %v to %v to be invalid", src, dst)
	}
	src = Pos{2, 1}
	dst = Pos{3, 2}
	if game.ValidMove(src, dst) {
		t.Errorf("expected %v to %v to be invalid, destination occupied", src, dst)
	}
	src = Pos{2, 3}
	dst = Pos{3, 4}
	if game.ValidMove(src, dst) {
		t.Errorf("expected %v to %v to be invalid, no piece at source", src, dst)
	}
	src = Pos{1, 4}
	dst = Pos{2, 3}
	game.Pieces[src] = Piece{Player: BlackPlayer, King: true}
	if !game.ValidMove(src, dst) {
		t.Errorf("expected %v to %v to be valid, kings can move backwards", src, dst)
	}
}

func TestValidJump(t *testing.T) {
	game := New()
	src := Pos{3, 2}
	dst := Pos{5, 4}
	capLoc := Capture(src, dst)
	if game.ValidJump(src, dst) {
		t.Errorf("expected jump %v to %v to be invalid, no piece to jump", src, dst)
	}
	capturePiece := Piece{Player: RedPlayer, King: false}
	game.Pieces[capLoc] = capturePiece
	if !game.ValidJump(src, dst) {
		t.Errorf("expected jump %v to %v to be valid", src, dst)
	}
	capturePiece.Player = BlackPlayer
	game.Pieces[capLoc] = capturePiece
	if game.ValidJump(src, dst) {
		t.Errorf("expected jump %v to %v to be invalid, can't jump own piece", src, dst)
	}
	jumpPiece := game.Pieces[src]
	delete(game.Pieces, src)
	jumpPiece.King = true
	game.Pieces[dst] = jumpPiece
	capturePiece.Player = RedPlayer
	game.Pieces[capLoc] = capturePiece
	tmp := src
	src = dst
	dst = tmp
	if !game.ValidJump(src, dst) {
		t.Errorf("expected jump %v to %v to be valid, kings can jump backwards", src, dst)
	}
	capturePiece.Player = BlackPlayer
	game.Pieces[capLoc] = capturePiece
	if game.ValidJump(src, dst) {
		t.Errorf("expected jump %v to %v to be invalid, can't jump own piece", src, dst)
	}
}

func TestMove(t *testing.T) {
	game := New()
	src := Pos{3, 2}
	dst := Pos{4, 3}
	_, err := game.Move(src, dst)
	if err != nil {
		t.Errorf("expected move to return no failure: %v", err)
	}
	if !game.PieceAt(dst) || game.PieceAt(src) {
		t.Errorf("expected piece to move from %v to %v", src, dst)
	}
}

func TestJump(t *testing.T) {
	game := New()
	src := Pos{3, 2}
	dst := Pos{5, 4}
	capLoc := Capture(src, dst)
	game.Pieces[capLoc] = Piece{Player: RedPlayer, King: false}
	_, err := game.Move(src, dst)
	if err != nil {
		t.Errorf("expected move to return no failure: %v", err)
	}
	if !game.PieceAt(dst) || game.PieceAt(src) || game.PieceAt(capLoc) {
		t.Errorf("expected jump from %v to %v and removal of %v", src, dst, capLoc)
	}
}

func TestInvalidJump(t *testing.T) {
	game := New()
	src := Pos{3, 2}
	dst := Pos{5, 4}
	capLoc := Capture(src, dst)
	game.Pieces[capLoc] = Piece{Player: BlackPlayer, King: false}
	_, err := game.Move(src, dst)
	if err == nil {
		t.Errorf("expected jump to fail with error")
	}
	if game.PieceAt(dst) || !game.PieceAt(src) || !game.PieceAt(capLoc) {
		t.Errorf("expected jump to have no effect", src, dst, capLoc)
	}
}

func TestKingPiece(t *testing.T) {
	game := New()
	redDst := Pos{0, 7}
	redPiece := game.Pieces[redDst]
	game.kingPiece(redDst)
	if redPiece.King != game.Pieces[redDst].King {
		t.Errorf("expected no change in king status of red piece at %v", redDst)
	}
	blkDst := Pos{1, 0}
	blackPiece := game.Pieces[blkDst]
	game.kingPiece(blkDst)
	if blackPiece.King != game.Pieces[blkDst].King {
		t.Errorf("expected no change in king status of black piece at %v", blkDst)
	}
	game.Pieces[redDst] = blackPiece
	game.Pieces[blkDst] = redPiece
	tmpDst := redDst
	redDst = blkDst
	blkDst = tmpDst
	game.kingPiece(redDst)
	if !game.Pieces[redDst].King {
		t.Errorf("expected red piece at %v to be a king", redDst)
	}
	game.kingPiece(blkDst)
	if !game.Pieces[blkDst].King {
		t.Errorf("expected black piece at %v to be a king", blkDst)
	}
}

func TestJumpPossibleFrom(t *testing.T) {
	game := New()
	src := Pos{3, 2}
	if game.jumpPossibleFrom(src) {
		t.Errorf("expected no jump possible from %v", src)
	}
	game.Pieces[Pos{2, 3}] = Piece{RedPlayer, false}
	if !game.jumpPossibleFrom(src) {
		t.Errorf("expected possible jump from %v", src)
	}
}

func TestJumpPossibleFromKing(t *testing.T) {
	game := New()
	src := Pos{2, 5}
	game.Pieces[src] = Piece{BlackPlayer, true}
	if game.jumpPossibleFrom(src) {
		t.Errorf("expected no jump possible for king from %v", src)
	}
	game.Pieces[Pos{3, 4}] = Piece{RedPlayer, false}
	if !game.jumpPossibleFrom(src) {
		t.Errorf("expected possible jump for king from %v", src)
	}
}

func TestPlayerHasMove(t *testing.T) {
	game := New()
	for loc := range game.Pieces {
		delete(game.Pieces, loc)
	}
	game.Pieces[Pos{0, 3}] = Piece{BlackPlayer, false}
	game.Pieces[Pos{1, 4}] = Piece{RedPlayer, false}
	game.Pieces[Pos{2, 5}] = Piece{RedPlayer, false}
	if game.playerHasMove(BlackPlayer) {
		t.Errorf("expected red player to have no move")
	}
	if !game.playerHasMove(RedPlayer) {
		t.Errorf("expected black player to have a move")
	}
}

func TestMovePossibleFrom(t *testing.T) {
	game := New()
	for loc := range game.Pieces {
		delete(game.Pieces, loc)
	}
	blockedSrc := Pos{0, 3}
	okSrc := Pos{1, 4}
	okKingSrc := Pos{2, 5}
	game.Pieces[blockedSrc] = Piece{BlackPlayer, false}
	game.Pieces[okSrc] = Piece{RedPlayer, false}
	game.Pieces[okKingSrc] = Piece{RedPlayer, true}
	game.Pieces[Pos{3, 4}] = Piece{RedPlayer, false}
	if game.movePossibleFrom(blockedSrc) {
		t.Errorf("expected no possible move from %v", blockedSrc)
	}
	if !game.movePossibleFrom(okSrc) {
		t.Errorf("expected possbile move for normal piece from %v", okSrc)
	}
	if !game.movePossibleFrom(okKingSrc) {
		t.Errorf("expected possible move for kinged piece from %v", okKingSrc)
	}
}

func TestUpdateTurnNormalMove(t *testing.T) {
	game := New()
	origTurn := game.Turn
	src := Pos{3, 2}
	dst := Pos{4, 3}
	game.Move(src, dst)
	if origTurn == game.Turn || game.Turn != RedPlayer {
		t.Errorf("expected turn to change from black to red after move from %v to %v", src, dst)
	}
	src = Pos{2, 5}
	dst = Pos{3, 4}
	game.Move(src, dst)
	if origTurn != game.Turn || game.Turn != BlackPlayer {
		t.Errorf("expected turn to change from red to black after move from %v to %v", src, dst)
	}
}

func TestUpdateTurnNormalJump(t *testing.T) {
	game := New()
	origTurn := game.Turn
	game.Pieces[Pos{2, 3}] = Piece{RedPlayer, false}
	src := Pos{3, 2}
	dst := Pos{1, 4}
	game.Move(src, dst)
	if origTurn == game.Turn || game.Turn != RedPlayer {
		t.Errorf("expected turn to change from black to red after jump from %v to %v", src, dst)
	}
	src = Pos{0, 5}
	dst = Pos{2, 3}
	game.Move(src, dst)
	if origTurn != game.Turn || game.Turn != BlackPlayer {
		t.Errorf("expected turn to change from black to red after jump from %v to %v", src, dst)
	}
}

func TestUpdateTurnJumpContinuation(t *testing.T) {
	game := New()
	origTurn := game.Turn
	game.Pieces[Pos{2, 3}] = Piece{RedPlayer, false}
	delete(game.Pieces, Pos{3, 6})
	src := Pos{3, 2}
	dst := Pos{1, 4}
	game.Move(src, dst)
	if origTurn != game.Turn || game.Turn != BlackPlayer {
		t.Errorf("expected no turn change after jump from %v to %v", src, dst)
	}
	src = Pos{1, 4}
	dst = Pos{3, 6}
	game.Move(src, dst)
	if origTurn == game.Turn || game.Turn != RedPlayer {
		t.Errorf("expected turn to change from black to red after jump from %v to %v", src, dst)
	}
}

func TestUpdateTurnNoMove(t *testing.T) {
	game := New()
	game.Turn = RedPlayer
	for loc := range game.Pieces {
		delete(game.Pieces, loc)
	}
	blkSrc := Pos{0, 3}
	redSrc := Pos{0, 5}
	blockLoc := Pos{1, 4}
	unblockLoc := Pos{2, 3}
	game.Pieces[blkSrc] = Piece{BlackPlayer, false}
	game.Pieces[redSrc] = Piece{RedPlayer, false}
	game.Pieces[Pos{2, 5}] = Piece{RedPlayer, false}
	game.Move(redSrc, blockLoc)
	if game.Turn != RedPlayer {
		t.Errorf("expected turn to remain on red since black has no move")
	}
	game.Move(blockLoc, unblockLoc)
	if game.Turn != BlackPlayer {
		t.Errorf("expected turn to change since black now has move")
	}
}

func TestUpdateTurnNoKingJumpContinuation(t *testing.T) {
	game := New()
	for loc := range game.Pieces {
		delete(game.Pieces, loc)
	}
	src := Pos{4, 5}
	dst := Pos{2, 7} // dst2 := Pos{0, 5}
	game.Pieces[Pos{3, 6}] = Piece{RedPlayer, false}
	game.Pieces[Pos{1, 6}] = Piece{RedPlayer, false}
	game.Pieces[src] = Piece{BlackPlayer, false}
	game.Move(src, dst)
	if game.Turn != RedPlayer {
		t.Errorf("expected turn to be red: new kings can't continue to jump")
	}
}

func TestUpdateTurnKingJumpContinuation(t *testing.T) {
	game := New()
	for loc := range game.Pieces {
		delete(game.Pieces, loc)
	}
	src := Pos{4, 5}
	dst1 := Pos{2, 7}
	dst2 := Pos{0, 5}
	game.Pieces[Pos{3, 6}] = Piece{RedPlayer, false}
	game.Pieces[Pos{1, 6}] = Piece{RedPlayer, false}
	game.Pieces[Pos{7, 6}] = Piece{RedPlayer, false}
	game.Pieces[src] = Piece{BlackPlayer, true}
	game.Move(src, dst1)
	if game.Turn == RedPlayer {
		t.Errorf("expected turn to be black: non-new kings can continue to jump")
	}
	game.Move(dst1, dst2)
	if game.Turn != RedPlayer {
		t.Errorf("expected turn to be red, king jump continuation is over")
	}
}

func TestMustJump(t *testing.T) {
	game := New()
	for loc := range game.Pieces {
		delete(game.Pieces, loc)
	}
	noJmpPos, jmpPos := Pos{1, 2}, Pos{3, 2}
	jumpedPos := Pos{4, 3}
	dstPos := Pos{0, 3}
	game.Pieces[noJmpPos], game.Pieces[jmpPos] = Piece{BlackPlayer, false}, Piece{BlackPlayer, false}
	game.Pieces[jumpedPos] = Piece{RedPlayer, false}
	capture, err := game.Move(noJmpPos, dstPos)
	if err == nil || game.Turn == RedPlayer || capture != NoPos {
		t.Error("should not allow non-jump when a jump is possible")
	}
}

func TestString(t *testing.T) {
	game := New()
	expected := "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"
	actual := game.String()
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	expected = "*B*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*R|r*r*r*r*"
	game.Pieces[Pos{1, 0}] = Piece{BlackPlayer, true}
	game.Pieces[Pos{7, 6}] = Piece{RedPlayer, true}
	actual = game.String()
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestParse(t *testing.T) {
	expected := New()
	actual, err := Parse(expected.String())
	if err != nil {
		t.Fatal(fmt.Sprintf("expected successful parse, instead got error: %v", err))
	}
	if expected.String() != actual.String() {
		t.Errorf("parsed game not equal to game: expected %v, got %v", expected, actual)
	}
}
