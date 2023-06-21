// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskKnightAttacks tests the MaskKnightAttacks function.
func TestMaskKnightAttacks(t *testing.T) {
	if result := MaskKnightAttacks(b.H4); result != 0x402000204000 {
		t.Errorf("MaskKnightAttacks failed: expected 0x402000204000, got 0x%x", result)
	}

	if result := MaskKnightAttacks(b.A1); result != 0x20400 {
		t.Errorf("MaskKnightAttacks failed: expected 0x20400, got 0x%x", result)
	}

	if result := MaskKnightAttacks(b.E5); result != 0x28440044280000 {
		t.Errorf("MaskKnightAttacks failed: expected 0x28440044280000, got 0x%x", result)
	}

	if result := MaskKnightAttacks(b.B7); result != 0x800080500000000 {
		t.Errorf("MaskKnightAttacks failed: expected 0x800080500000000, got 0x%x", result)
	}
}
