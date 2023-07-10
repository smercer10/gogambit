// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestGenPawnAttacks tests the GenPawnAttacks function.
func TestGenPawnAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		side   int
		expect Bitboard
	}{
		{H4, Black, 0x400000},
		{B3, White, 0x5000000},
		{E8, White, 0x0},
		{A1, Black, 0x0},
	}

	for _, tc := range testCases {
		if result := GenPawnAttacks(tc.sq, tc.side); result != tc.expect {
			t.Errorf("GenPawnAttacks failed for sq = %s, side = %s: expect 0x%x, got 0x%x",
				Squares[tc.sq], Sides[tc.side], tc.expect, result)
		}
	}
}

// TestGenKnightAttacks tests the GenKnightAttacks function.
func TestGenKnightAttacks(t *testing.T) {
	testCases := []struct {
		sq     int
		expect Bitboard
	}{
		{H4, 0x402000204000},
		{A1, 0x20400},
		{E5, 0x28440044280000},
		{B7, 0x800080500000000},
	}

	for _, tc := range testCases {
		if result := GenKnightAttacks(tc.sq); result != tc.expect {
			t.Errorf("GenKnightAttacks failed for sq = %s: expect 0x%x, got 0x%x",
				Squares[tc.sq], tc.expect, result)
		}
	}
}

// TestGenKingAttacks tests the GenKingAttacks function.
func TestGenKingAttacks(t *testing.T) {
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
		if result := GenKingAttacks(tc.sq); result != tc.expect {
			t.Errorf("GenKingAttacks failed for sq = %s: expect 0x%x, got 0x%x",
				Squares[tc.sq], tc.expect, result)
		}
	}
}
