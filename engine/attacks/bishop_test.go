// Package attacks provides attack generation utilities.
package attacks

import (
	b "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
	"testing"
)

// TestMaskBishopOccupancy tests the MaskBishopOccupancy function.
func TestMaskBishopOccupancy(t *testing.T) {
	testCases := []struct {
		sq     int
		expect b.Bitboard
	}{
		{H8, 0x40201008040200},
		{A3, 0x10080402000200},
		{E5, 0x44280028440200},
		{D8, 0x14224000000000},
	}

	for _, tc := range testCases {
		if result := MaskBishopOccupancy(tc.sq); result != tc.expect {
			t.Errorf("MaskBishopOccupancy failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}
