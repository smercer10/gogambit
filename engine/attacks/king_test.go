// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestMaskKingAttacks tests the MaskKingAttacks function.
func TestMaskKingAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		expect Bitboard
	}{
		{H8, 0x40c0000000000000},
		{A1, 0x302},
		{E5, 0x382838000000},
		{A4, 0x302030000},
	}

	for _, tc := range testCases {
		if result := MaskKingAttacks(tc.sq); result != tc.expect {
			t.Errorf("MaskKingAttacks failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}
