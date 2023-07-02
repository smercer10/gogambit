// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
	"testing"
)

// TestMaskPawnAttacks tests the MaskPawnAttacks function.
func TestMaskPawnAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		side   int
		expect Bitboard
	}{
		{H4, Black, 0x400000},
		{B3, White, 0x5000000},
		{E8, White, 0x0},
		{A1, Black, 0x0},
	}

	for _, tc := range testCases {
		if result := MaskPawnAttacks(tc.sq, tc.side); result != tc.expect {
			t.Errorf("MaskPawnAttacks failed for sq = %d, side = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.side, tc.expect, result)
		}
	}
}
