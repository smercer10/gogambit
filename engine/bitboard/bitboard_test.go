// Package b provides bitboard utilities.
package b

import (
	c "gogambit/engine/const"
	"testing"
)

// TestSetBit tests the SetBit function.
func TestSetBit(t *testing.T) {
	var bb Bitboard = 0x0

	if bb = SetBit(bb, c.A1); bb != 0x1 {
		t.Errorf("SetBit failed: expected 0x1, got 0x%x", bb)
	}

	if bb = SetBit(bb, c.H8); bb != 0x8000000000000001 {
		t.Errorf("SetBit failed: expected 0x8000000000000001, got 0x%x", bb)
	}
}

// TestClearBit tests the ClearBit function.
func TestClearBit(t *testing.T) {
	var bb Bitboard = 0x8000000000000001

	if bb = ClearBit(bb, c.A1); bb != 0x8000000000000000 {
		t.Errorf("ClearBit failed: expected 0x8000000000000000, got 0x%x", bb)
	}

	if bb = ClearBit(bb, c.H8); bb != 0x0 {
		t.Errorf("ClearBit failed: expected 0x0, got 0x%x", bb)
	}
}

// TestGetBit tests the GetBit function.
func TestGetBit(t *testing.T) {
	var bb Bitboard = 0x8000000000000001

	if !GetBit(bb, c.A1) {
		t.Errorf("GetBit failed: expected true, got false")
	}

	if !GetBit(bb, c.H8) {
		t.Errorf("GetBit failed: expected true, got false")
	}

	if GetBit(bb, c.A2) {
		t.Errorf("GetBit failed: expected false, got true")
	}

	if GetBit(bb, c.H7) {
		t.Errorf("GetBit failed: expected false, got true")
	}
}