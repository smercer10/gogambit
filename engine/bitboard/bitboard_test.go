// Package bitboard provides the bitboard type and utilities.
package bitboard

import "testing"

// TestSetBit tests the SetBit function.
func TestSetBit(t *testing.T) {
	bb := Bitboard(0x0)

	bb.SetBit(A1)

	if bb != 0x1 {
		t.Errorf("SetBit failed: expect 0x1, got 0x%x", bb)
	}

	bb.SetBit(H8)

	if bb != 0x8000000000000001 {
		t.Errorf("SetBit failed: expect 0x8000000000000001, got 0x%x", bb)
	}
}

// TestClearBit tests the ClearBit function.
func TestClearBit(t *testing.T) {
	bb := Bitboard(0x8000000000000001)

	bb.ClearBit(A1)

	if bb != 0x8000000000000000 {
		t.Errorf("ClearBit failed: expect 0x8000000000000000, got 0x%x", bb)
	}

	bb.ClearBit(H8)

	if bb != 0x0 {
		t.Errorf("ClearBit failed: expect 0x0, got 0x%x", bb)
	}
}

// TestGetBit tests the GetBit function.
func TestGetBit(t *testing.T) {
	bb := Bitboard(0x8000000000000001)

	testCases := []struct {
		sq     int
		expect bool
	}{
		{A1, true},
		{H8, true},
		{A2, false},
		{H7, false},
	}

	for _, tc := range testCases {
		if result := bb.GetBit(tc.sq); result != tc.expect {
			t.Errorf("GetBit failed for bb = 0x%x, sq = %d: expect %t, got %t",
				bb, tc.sq, tc.expect, result)
		}
	}
}
