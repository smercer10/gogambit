// Package attacks provides attack generation utilities.
package attacks

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskPawnAttacks tests the MaskPawnAttacks function.
func TestMaskPawnAttacks(t *testing.T) {
	if result := MaskPawnAttacks(b.H4, Black); result != 0x400000 {
		t.Errorf("MaskPawnAttacks failed: expected 0x400000, got 0x%x", result)
	}

	if result := MaskPawnAttacks(b.B3, White); result != 0x5000000 {
		t.Errorf("MaskPawnAttacks failed: expected 0x5000000, got 0x%x", result)
	}

	if result := MaskPawnAttacks(b.E8, White); result != 0x0 {
		t.Errorf("MaskPawnAttacks failed: expected 0x0, got 0x%x", result)
	}

	if result := MaskPawnAttacks(b.A1, Black); result != 0x0 {
		t.Errorf("MaskPawnAttacks failed: expected 0x0, got 0x%x", result)
	}
}
