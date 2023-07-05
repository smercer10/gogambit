// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestMaskRelevantRookOccupancy tests the MaskRelevantRookOccupancy function.
func TestMaskRelevantRookOccupancy(t *testing.T) {
	testCases := []struct {
		sq     int
		expect Bitboard
	}{
		{H8, 0x7e80808080808000},
		{A3, 0x10101017e0100},
		{E5, 0x10106e10101000},
		{D8, 0x7608080808080800},
	}

	for _, tc := range testCases {
		if result := MaskRelevantRookOccupancy(tc.sq); result != tc.expect {
			t.Errorf("MaskRelevantRookOccupancy failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}

// TestGenRookAttacksOnTheFly tests the GenRookAttacksOnTheFly function.
func TestGenRookAttacksOnTheFly(t *testing.T) {
	testCases := []struct {
		sq       int
		blockers Bitboard
		expect   Bitboard
	}{
		{D4, 0x822000000, 0x836080808},
		{E3, 0x10000000880000, 0x10101010e81010},
		{A1, 0x100010060, 0x1013e},
		{H8, 0x0, 0x7f80808080808080},
	}

	for _, tc := range testCases {
		if result := GenRookAttacksOnTheFly(tc.sq, tc.blockers); result != tc.expect {
			t.Errorf("GenRookAttacksOnTheFly failed for sq = %d, blockers = 0x%x: expect 0x%x, got 0x%x",
				tc.sq, tc.blockers, tc.expect, result)
		}
	}
}
