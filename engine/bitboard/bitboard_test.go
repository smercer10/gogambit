// Package bitboard provides the bitboard type and relevant utilities.
package bitboard

import (
	. "gogambit/engine/globals"
	"testing"
)

// TestSetBit tests the SetBit method.
func TestSetBit(t *testing.T) {
	bb := Bitboard(0x0)

	bb = bb.SetBit(A1)

	if bb != 0x1 {
		t.Errorf("SetBit failed: expect 0x1, got 0x%x", bb)
	}

	bb = bb.SetBit(H8)

	if bb != 0x8000000000000001 {
		t.Errorf("SetBit failed: expect 0x8000000000000001, got 0x%x", bb)
	}
}

// TestClearBit tests the ClearBit method.
func TestClearBit(t *testing.T) {
	bb := Bitboard(0x8000000000000001)

	bb = bb.ClearBit(A1)

	if bb != 0x8000000000000000 {
		t.Errorf("ClearBit failed: expect 0x8000000000000000, got 0x%x", bb)
	}

	bb = bb.ClearBit(H8)

	if bb != 0x0 {
		t.Errorf("ClearBit failed: expect 0x0, got 0x%x", bb)
	}
}

// TestIsSet tests the IsSet method.
func TestIsSet(t *testing.T) {
	bb := Bitboard(0x8000000000000001)

	testCases := []struct {
		sq     int
		expect bool
	}{
		{A1, true},
		{H8, true},
		{A2, false},
		{H7, false},
	}

	for _, tc := range testCases {
		if result := bb.IsSet(tc.sq); result != tc.expect {
			t.Errorf("IsSet failed for bb = 0x%x, sq = %s: expect %t, got %t",
				bb, Squares[tc.sq], tc.expect, result)
		}
	}
}

// TestCountBits tests the CountBits method.
func TestCountBits(t *testing.T) {
	testCases := []struct {
		bb     Bitboard
		expect int
	}{
		{0x0, 0},
		{0x1, 1},
		{0x7f80808080808080, 14},
		{0x8f7080800, 10},
	}

	for _, tc := range testCases {
		if result := tc.bb.CountBits(); result != tc.expect {
			t.Errorf("CountBits failed for bb = 0x%x: expect %d, got %d",
				tc.bb, tc.expect, result)
		}
	}
}

// TestGetLSB tests the GetLSB method.
func TestGetLSB(t *testing.T) {
	testCases := []struct {
		bb     Bitboard
		expect int
	}{
		{0x0, -1},
		{0x1, A1},
		{0x80000000800002, B1},
		{0x880000000000000, H7},
	}

	for _, tc := range testCases {
		if result := tc.bb.GetLSB(); result != tc.expect {
			t.Errorf("GetLSB failed for bb = 0x%x: expect %s, got %s",
				tc.bb, Squares[tc.expect], Squares[result])
		}
	}
}

// TestParseFEN tests the ParseFEN function.
func TestParseFEN(t *testing.T) {
	testCases := []struct {
		fen         string
		whiteOcc    Bitboard
		blackOcc    Bitboard
		sideToMove  int
		castRights  int
		enPassantSq int
	}{
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			0xffff, 0xffff000000000000, White, 0b1111, NA,
		},
		{
			"4k3/8/8/8/8/8/5PPP/4K3 b - - 0 55",
			0xe010, 0x1000000000000000, Black, 0b000, NA,
		},
		{
			"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/6P1/PP2PP1P/RNBQKBNR b KQkq - 0 3",
			0xc40b3ff, 0xffe7100800000000, Black, 0b1111, NA,
		},
		{
			"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			0x1000efff, 0xffff000000000000, Black, 0b1111, E3,
		},
	}

	for _, tc := range testCases {
		ParseFEN(tc.fen)

		if SideOcc[White] != tc.whiteOcc {
			t.Errorf("ParseFEN failed: expect whiteOcc = 0x%x, got 0x%x",
				tc.whiteOcc, SideOcc[White])
		}

		if SideOcc[Black] != tc.blackOcc {
			t.Errorf("ParseFEN failed: expect blackOcc = 0x%x, got 0x%x",
				tc.blackOcc, SideOcc[Black])
		}

		if SideToMove != tc.sideToMove {
			t.Errorf("ParseFEN failed: expect sideToMove = %s, got %s",
				Sides[tc.sideToMove], Sides[SideToMove])
		}

		if CastRights != tc.castRights {
			t.Errorf("ParseFEN failed: expect castRights = 0b%b, got 0b%b",
				tc.castRights, CastRights)
		}

		if EnPassantSq != tc.enPassantSq {
			t.Errorf("ParseFEN failed: expect enPassantSq = %s, got %s",
				Squares[tc.enPassantSq], Squares[EnPassantSq])
		}
	}
}

// TestCopyBoardAndTakeBack tests the CopyBoard and TakeBack functions.
func TestCopyBoardAndTakeBack(t *testing.T) {
	ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	CopyBoard()

	ParseFEN("8/8/8/8/8/8/8/8 b - - 0 1")

	TakeBack()

	if SideOcc[White] != 0xffff {
		t.Errorf("CopyBoard and TakeBack failed: expect whiteOcc = 0xffff, got 0x%x",
			SideOcc[White])
	}

	if SideToMove != White {
		t.Errorf("CopyBoard and TakeBack failed: expect sideToMove = White, got %s",
			Sides[SideToMove])
	}

	if CastRights != 0b1111 {
		t.Errorf("CopyBoard and TakeBack failed: expect castRights = 0b1111, got 0b%b",
			CastRights)
	}
}
