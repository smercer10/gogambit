// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestSetOccupancy tests the SetOccupancy function.
func TestSetOccupancy(t *testing.T) {
	testCases := []struct {
		mask   Bitboard
		bits   int
		idx    int
		expect Bitboard
	}{
		{MaskRelBishopOcc(A1), BishopRelOccBits[A1], 3, 0x40200},
		{MaskRelRookOcc(G6), RookRelOccBits[G6], 3965, 0x402e4040004000},
		{MaskRelRookOcc(D2), RookRelOccBits[D2], 455, 0x80808001600},
		{MaskRelBishopOcc(C3), BishopRelOccBits[C3], 1, 0x200},
	}

	for _, tc := range testCases {
		if result := SetOccupancy(tc.mask, tc.bits, tc.idx); result != tc.expect {
			t.Errorf("SetOccupancy failed for mask = 0x%x, bits = %d, idx = %d: expect 0x%x, got 0x%x",
				tc.mask, tc.bits, tc.idx, tc.expect, result)
		}
	}
}

// TestGetQueenAttacks tests the GetQueenAttacks function.
func TestGetQueenAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		occ    Bitboard
		expect Bitboard
	}{
		{D4, 0x80000100200, 0x80412a1cf71c0a08},
		{A1, 0x0, 0x81412111090503fe},
		{G7, 0x8008200000000000, 0xe0b8e04040404040},
		{C4, 0x400040000404, 0x4424150e7b0e1520},
	}

	InitSliderAtt(Bishop)
	InitSliderAtt(Rook)

	for _, tc := range testCases {
		if result := GetQueenAttacks(tc.sq, tc.occ); result != tc.expect {
			t.Errorf("GetQueenAttacks failed for sq = %s, occ = 0x%x: expect 0x%x, got 0x%x",
				Squares[tc.sq], tc.occ, tc.expect, result)
		}
	}
}
