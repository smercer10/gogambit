// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskKingAttacks tests the MaskKingAttacks function.
func TestMaskKingAttacks(t *testing.T) {
	testCases := []struct {
		sq       int
		expected b.Bitboard
	}{
		{b.H8, 0x40c0000000000000},
		{b.A1, 0x302},
		{b.E5, 0x382838000000},
		{b.A4, 0x302030000},
	}

	for _, tc := range testCases {
		if result := MaskKingAttacks(tc.sq); result != tc.expected {
			t.Errorf("MaskKingAttacks failed for sq = %d: expected 0x%x, got 0x%x",
				tc.sq, tc.expected, result)
		}
	}
}
