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
		mask         Bitboard
		maskBitCount int
		idx          int
		expect       Bitboard
	}{
		{MaskRelevantBishopOccupancy(A1), BishopRelevantOccupancyBitCounts[A1], 3, 0x40200},
		{MaskRelevantRookOccupancy(G6), RookRelevantOccupancyBitCounts[G6], 3965, 0x402e4040004000},
		{MaskRelevantRookOccupancy(D2), RookRelevantOccupancyBitCounts[D2], 455, 0x80808001600},
		{MaskRelevantBishopOccupancy(C3), BishopRelevantOccupancyBitCounts[C3], 1, 0x200},
	}

	for _, tc := range testCases {
		if result := SetOccupancy(tc.mask, tc.maskBitCount, tc.idx); result != tc.expect {
			t.Errorf("SetOccupancy failed for mask = 0x%x, maskBitCount = %d, idx = %d: expect 0x%x, got 0x%x",
				tc.mask, tc.maskBitCount, tc.idx, tc.expect, result)
		}
	}
}

// TestGetQueenAttacks tests the GetQueenAttacks function.
func TestGetQueenAttacks(t *testing.T) {
	testCases := []struct {
		sq        int
		occupancy Bitboard
		expect    Bitboard
	}{
		{D4, 0x80000100200, 0x80412a1cf71c0a08},
		{A1, 0x0, 0x81412111090503fe},
		{G7, 0x8008200000000000, 0xe0b8e04040404040},
		{C4, 0x400040000404, 0x4424150e7b0e1520},
	}

	InitSliderAttacks(Bishop)
	InitSliderAttacks(Rook)

	for _, tc := range testCases {
		if result := GetQueenAttacks(tc.sq, tc.occupancy); result != tc.expect {
			t.Errorf("GetQueenAttacks failed for sq = %d, occupancy = 0x%x: expect 0x%x, got 0x%x",
				tc.sq, tc.occupancy, tc.expect, result)
		}
	}
}
