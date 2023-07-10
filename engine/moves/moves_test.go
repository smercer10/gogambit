// Package moves provides move generation utilities.
package moves

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"testing"
)

// TestIsAttacked tests the IsAttacked function.
func TestIsAttacked(t *testing.T) {
	testCases := []struct {
		sq     int
		by     int
		expect bool
	}{
		{H4, White, true},
		{A7, Black, true},
		{D1, White, true},
		{C1, Black, true},
		{E7, Black, true},
		{H6, White, false},
		{A1, White, false},
		{A3, Black, false},
		{H8, Black, false},
		{F4, Black, false},
	}

	a.InitLeaperAtt()
	a.InitSliderAtt(Bishop)
	a.InitSliderAtt(Rook)
	ParseFEN("r3k1nr/1bq2ppp/p2p4/1p1P1Q2/1P6/1B4P1/4PPBP/R5K1 b kq - 0 27")

	for _, tc := range testCases {
		if result := IsAttacked(tc.sq, tc.by); result != tc.expect {
			t.Errorf("IsAttacked failed for sq = %s, by = %s: expect %t, got %t",
				Squares[tc.sq], Sides[tc.by], tc.expect, result)
		}
	}
}
