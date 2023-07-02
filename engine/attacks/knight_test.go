// Package attacks provides attack generation utilities.
package attacks

import (
	b "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
	"testing"
)

// TestMaskKnightAttacks tests the MaskKnightAttacks function.
func TestMaskKnightAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		expect b.Bitboard
	}{
		{H4, 0x402000204000},
		{A1, 0x20400},
		{E5, 0x28440044280000},
		{B7, 0x800080500000000},
	}

	for _, tc := range testCases {
		if result := MaskKnightAttacks(tc.sq); result != tc.expect {
			t.Errorf("MaskKnightAttacks failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}
