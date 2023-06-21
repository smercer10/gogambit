// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskKnightAttacks tests the MaskKnightAttacks function.
func TestMaskKnightAttacks(t *testing.T) {
	testCases := []struct {
		sq       int
		expected b.Bitboard
	}{
		{b.H4, 0x402000204000},
		{b.A1, 0x20400},
		{b.E5, 0x28440044280000},
		{b.B7, 0x800080500000000},
	}

	for _, tc := range testCases {
		if result := MaskKnightAttacks(tc.sq); result != tc.expected {
			t.Errorf("MaskKnightAttacks failed for sq = %d: expected 0x%x, got 0x%x",
				tc.sq, tc.expected, result)
		}
	}
}
