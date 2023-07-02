// Package attack provides attack generation utilities.
package attack

import b "gogambit/engine/bitboard"

// MaskRookOccupancy masks the relevant rook occupancy bits for a given square.
// This forms a key for magic bitboards.
func MaskRookOccupancy(sq int) b.Bitboard {
	occupancy := b.Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// N
	for r := tr + 1; r <= 6; r++ {
		occupancy.SetBit(r*8 + tf)
	}

	// S
	for r := tr - 1; r >= 1; r-- {
		occupancy.SetBit(r*8 + tf)
	}

	// E
	for f := tf + 1; f <= 6; f++ {
		occupancy.SetBit(tr*8 + f)
	}

	// W
	for f := tf - 1; f >= 1; f-- {
		occupancy.SetBit(tr*8 + f)
	}

	return occupancy
}
