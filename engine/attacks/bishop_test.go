// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
	"testing"
)

// TestMaskBishopOccupancy tests the MaskBishopOccupancy function.
func TestMaskBishopOccupancy(t *testing.T) {
	testCases := []struct {
		sq     int
		expect Bitboard
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

// TestGenBishopAttacksOnTheFly tests the GenBishopAttacksOnTheFly function.
func TestGenBishopAttacksOnTheFly(t *testing.T) {
	testCases := []struct {
		sq      int
		blocker Bitboard
		expect  Bitboard
	}{
		{D4, 0x40020000140000, 0x40221400140000},
		{E3, 0x20020000000, 0x20428002844},
		{A1, 0x40000, 0x40200},
		{H8, 0x0, 0x40201008040201},
	}

	for _, tc := range testCases {
		if result := GenBishopAttacksOnTheFly(tc.sq, tc.blocker); result != tc.expect {
			t.Errorf("GenBishopAttacksOnTheFly failed for sq = %d, blocker = 0x%x: expect 0x%x, got 0x%x",
				tc.sq, tc.blocker, tc.expect, result)
		}
	}
}