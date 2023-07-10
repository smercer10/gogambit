// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestMaskRelBishopOcc tests the MaskRelBishopOcc function.
func TestMaskRelBishopOcc(t *testing.T) {
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
		if result := MaskRelBishopOcc(tc.sq); result != tc.expect {
			t.Errorf("MaskRelBishopOcc failed for sq = %d: expect 0x%x, got 0x%x",
				tc.sq, tc.expect, result)
		}
	}
}

// TestGenBishopAttOTF tests the GenBishopAttOTF function.
func TestGenBishopAttOTF(t *testing.T) {
	testCases := []struct {
		sq     int
		occ    Bitboard
		expect Bitboard
	}{
		{D4, 0x40020000140000, 0x40221400140000},
		{E3, 0x20020000000, 0x20428002844},
		{A1, 0x40000, 0x40200},
		{H8, 0x0, 0x40201008040201},
	}

	for _, tc := range testCases {
		if result := GenBishopAttOTF(tc.sq, tc.occ); result != tc.expect {
			t.Errorf("GenBishopAttOTF failed for sq = %d, occ = 0x%x: expect 0x%x, got 0x%x",
				tc.sq, tc.occ, tc.expect, result)
		}
	}
}

// TestGetBishopAttacks tests the GetBishopAttacks function.
func TestGetBishopAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		occ    Bitboard
		expect Bitboard
	}{
		{D4, 0x200102000000, 0x1221400142241},
		{E7, 0x280000100000, 0x2800280000000000},
		{A1, 0x0, 0x8040201008040200},
		{A1, 0x40000, 0x40200},
		{G5, 0x10800080000800, 0x10a000a0100800},
		{C2, 0x1001000000, 0x804020110a000a},
	}

	InitSliderAttacks(Bishop)

	for _, tc := range testCases {
		if result := GetBishopAttacks(tc.sq, tc.occ); result != tc.expect {
			t.Errorf("GetBishopAttacks failed for sq = %d, occ = 0x%x: expect 0x%x, got 0x%x",
				tc.sq, tc.occ, tc.expect, result)
		}
	}
}
