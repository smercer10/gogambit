// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestSetOccupancy tests the SetOccupancy function.
func TestSetOccupancy(t *testing.T) {
	testCases := []struct {
		mask   Bitboard
		idx    int
		expect Bitboard
	}{
		{MaskRelevantBishopOccupancy(D4), 2999, 0x400042200},
		{MaskRelevantRookOccupancy(G6), 3965, 0x24040004000},
		{MaskRelevantRookOccupancy(A1), 455, 0xe},
		{MaskRelevantBishopOccupancy(C3), 1, 0x200},
	}

	for _, tc := range testCases {
		if result := SetOccupancy(tc.mask, tc.idx); result != tc.expect {
			t.Errorf("SetOccupancy failed for mask = 0x%x, idx = %d: expect 0x%x, got 0x%x",
				tc.mask, tc.idx, tc.expect, result)
		}
	}
}
