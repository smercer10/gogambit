// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

func TestMaskBishopOccupancy(t *testing.T) {
	testCases := []struct {
		sq     int
		expect b.Bitboard
	}{
		{b.H8, 0x40201008040200},
		{b.A3, 0x10080402000200},
		{b.E5, 0x44280028440200},
		{b.D8, 0x14224000000000},
	}

	for _, tc := range testCases {
		if result := MaskBishopOccupancy(tc.sq); result != tc.expect {
			t.Errorf("MaskBishopOccupancy failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}
