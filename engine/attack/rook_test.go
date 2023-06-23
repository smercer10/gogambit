// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskRookOccupancy tests the MaskRookOccupancy function.
func TestMaskRookOccupancy(t *testing.T) {
	testCases := []struct {
		sq     int
		expect b.Bitboard
	}{
		{b.H8, 0x7e80808080808000},
		{b.A3, 0x10101017e0100},
		{b.E5, 0x10106e10101000},
		{b.D8, 0x7608080808080800},
	}

	for _, tc := range testCases {
		if result := MaskRookOccupancy(tc.sq); result != tc.expect {
			t.Errorf("MaskRookOccupancy failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}
