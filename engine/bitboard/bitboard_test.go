// Package bitboard provides the bitboard type and relevant utilities.
package bitboard

import (
	. "gogambit/engine/globals"
	"testing"
)

// TestSetBit tests the SetBit function.
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

// TestClearBit tests the ClearBit function.
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

// TestGetBit tests the GetBit function.
func TestGetBit(t *testing.T) {
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
		if result := bb.GetBit(tc.sq); result != tc.expect {
			t.Errorf("GetBit failed for bb = 0x%x, sq = %d: expect %t, got %t",
				bb, tc.sq, tc.expect, result)
		}
	}
}

// TestCountBits tests the CountBits function.
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

// TestGetLeastSignificantBit tests the GetLeastSignificantBit function.
func TestGetLeastSignificantBit(t *testing.T) {
	testCases := []struct {
		bb     Bitboard
		expect int
	}{
		{0x0, -1},
		{0x1, 0},
		{0x80000000800002, 1},
		{0x880000000000000, 55},
	}

	for _, tc := range testCases {
		if result := tc.bb.GetLeastSignificantBit(); result != tc.expect {
			t.Errorf("GetLeastSignificantBit failed for bb = 0x%x: expect %d, got %d",
				tc.bb, tc.expect, result)
		}
	}
}

// TestParseFEN tests the ParseFEN function.
func TestParseFEN(t *testing.T) {
	testCases := []struct {
		fen             string
		whiteOcc        Bitboard
		blackOcc        Bitboard
		sideToMove      int
		castlingRights  int
		enPassantSquare int
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

		if SideBitboards[White] != tc.whiteOcc {
			t.Errorf("ParseFEN failed: expect whiteOcc = 0x%x, got 0x%x",
				tc.whiteOcc, SideBitboards[White])
		}

		if SideBitboards[Black] != tc.blackOcc {
			t.Errorf("ParseFEN failed: expect blackOcc = 0x%x, got 0x%x",
				tc.blackOcc, SideBitboards[Black])
		}

		if SideToMove != tc.sideToMove {
			t.Errorf("ParseFEN failed: expect sideToMove = %s, got %s",
				Sides[tc.sideToMove], Sides[SideToMove])
		}

		if CastlingRights != tc.castlingRights {
			t.Errorf("ParseFEN failed: expect castlingRights = 0b%b, got 0b%b",
				tc.castlingRights, CastlingRights)
		}

		if EnPassantSquare != tc.enPassantSquare {
			t.Errorf("ParseFEN failed: expect enPassantSquare = %s, got %s",
				Squares[tc.enPassantSquare], Squares[EnPassantSquare])
		}
	}
}
