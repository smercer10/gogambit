// Package attack provides attack generation utilities.
package attack

import (
	b "gogambit/engine/bitboard"
	"testing"
)

// TestMaskKingAttacks tests the MaskKingAttacks function.
func TestMaskKingAttacks(t *testing.T) {
	if result := MaskKingAttacks(b.H8); result != 0x40c0000000000000 {
		t.Errorf("Expected 0x40c0000000000000, got %x", result)
	}

	if result := MaskKingAttacks(b.A1); result != 0x302 {
		t.Errorf("Expected 0x302, got %x", result)
	}

	if result := MaskKingAttacks(b.E5); result != 0x382838000000 {
		t.Errorf("Expected 0x382838000000, got %x", result)
	}

	if result := MaskKingAttacks(b.A4); result != 0x302030000 {
		t.Errorf("Expected 0x302030000, got %x", result)
	}
}
