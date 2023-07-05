// Package attacks provides attack generation utilities.
package attacks

import . "gogambit/engine/bitboard"

// MaskRelevantRookOccupancy masks the relevant rook occupancy bits for a given square.
// This forms a key for magic bitboards.
func MaskRelevantRookOccupancy(sq int) Bitboard {
	occupancy := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// N
	for r := tr + 1; r <= 6; r++ {
		occupancy = occupancy.SetBit(r*8 + tf)
	}

	// S
	for r := tr - 1; r >= 1; r-- {
		occupancy = occupancy.SetBit(r*8 + tf)
	}

	// E
	for f := tf + 1; f <= 6; f++ {
		occupancy = occupancy.SetBit(tr*8 + f)
	}

	// W
	for f := tf - 1; f >= 1; f-- {
		occupancy = occupancy.SetBit(tr*8 + f)
	}

	return occupancy
}

// RookRelevantOccupancyBitCounts is a LUT with the bit count of the relevant rook occupancies for each square.
var RookRelevantOccupancyBitCounts = [64]int{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

// GenRookAttacksOnTheFly generates possible rook attacks for a given square and mask of blockers.
func GenRookAttacksOnTheFly(sq int, blockers Bitboard) Bitboard {
	attacks := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// N
	for r := tr + 1; r <= 7; r++ {
		attacks = attacks.SetBit(r*8 + tf)

		if blockers.GetBit(r*8 + tf) {
			break
		}
	}

	// S
	for r := tr - 1; r >= 0; r-- {
		attacks = attacks.SetBit(r*8 + tf)

		if blockers.GetBit(r*8 + tf) {
			break
		}
	}

	// E
	for f := tf + 1; f <= 7; f++ {
		attacks = attacks.SetBit(tr*8 + f)

		if blockers.GetBit(tr*8 + f) {
			break
		}
	}

	// W
	for f := tf - 1; f >= 0; f-- {
		attacks = attacks.SetBit(tr*8 + f)

		if blockers.GetBit(tr*8 + f) {
			break
		}
	}

	return attacks
}
