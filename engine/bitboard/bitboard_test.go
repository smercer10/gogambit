// Package bitboard provides bitboard utilities.
package bitboard

import "testing"

// TestSetBit tests the SetBit function.
func TestSetBit(t *testing.T) {
	var bb Bitboard = 0x0

	SetBit(&bb, A1)

	if bb != 0x1 {
		t.Errorf("SetBit failed: expected 0x1, got 0x%x", bb)
	}

	SetBit(&bb, H8)

	if bb != 0x8000000000000001 {
		t.Errorf("SetBit failed: expected 0x8000000000000001, got 0x%x", bb)
	}
}

// TestClearBit tests the ClearBit function.
func TestClearBit(t *testing.T) {
	var bb Bitboard = 0x8000000000000001

	ClearBit(&bb, A1)

	if bb != 0x8000000000000000 {
		t.Errorf("ClearBit failed: expected 0x8000000000000000, got 0x%x", bb)
	}

	ClearBit(&bb, H8)

	if bb != 0x0 {
		t.Errorf("ClearBit failed: expected 0x0, got 0x%x", bb)
	}
}

// TestGetBit tests the GetBit function.
func TestGetBit(t *testing.T) {
	var bb Bitboard = 0x8000000000000001

	testCases := []struct {
		sq       int
		expected bool
	}{
		{A1, true},
		{H8, true},
		{A2, false},
		{H7, false},
	}

	for _, tc := range testCases {
		if result := GetBit(bb, tc.sq); result != tc.expected {
			t.Errorf("GetBit failed for bb = 0x%x, sq = %d: expected %t, got %t",
				bb, tc.sq, tc.expected, result)
		}
	}
}
