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

// TestEncAndDec tests the EncMove and associated Dec functions.
func TestEncAndDec(t *testing.T) {
	m := EncMove(E2, E4, WP, 0b1111, 0, 1, 0, 0)

	if m != 3081996 {
		t.Errorf("EncMove failed: expect 3081996, got %d", m)
	}

	src := DecSrc(m)
	trgt := DecTrgt(m)
	pc := DecPc(m)
	prom := DecProm(m)
	capt := DecCap(m)
	dp := DecDp(m)
	ep := DecEp(m)
	cast := DecCast(m)

	if src != E2 {
		t.Errorf("DecSrc failed: expect e2, got %s", Squares[src])
	}

	if trgt != E4 {
		t.Errorf("DecTrgt failed: expect e4, got %s", Squares[trgt])
	}

	if pc != WP {
		t.Errorf("DecPc failed: expect WP, got %s", Pieces[pc])
	}

	if prom != 0b1111 {
		t.Errorf("DecProm failed: expect 0b1111, got 0b%b", prom)
	}

	if capt != 0 {
		t.Errorf("DecCap failed: expect 0, got %d", capt)
	}

	if dp != 1 {
		t.Errorf("DecDp failed: expect 1, got %d", dp)
	}

	if ep != 0 {
		t.Errorf("DecEp failed: expect 0, got %d", ep)
	}

	if cast != 0 {
		t.Errorf("DecCast failed: expect 0, got %d", cast)
	}
}
