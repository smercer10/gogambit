// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskPawnAttacks tests the MaskPawnAttacks function.
func TestMaskPawnAttacks(t *testing.T) {
	testCases := []struct {
		sq       int
		side     int
		expected b.Bitboard
	}{
		{b.H4, Black, 0x400000},
		{b.B3, White, 0x5000000},
		{b.E8, White, 0x0},
		{b.A1, Black, 0x0},
	}

	for _, tc := range testCases {
		if result := MaskPawnAttacks(tc.sq, tc.side); result != tc.expected {
			t.Errorf("MaskPawnAttacks failed for sq = %d, side = %d: expected 0x%x, got 0x%x",
				tc.sq, tc.side, tc.expected, result)
		}
	}
}
