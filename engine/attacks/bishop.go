// Package attacks provides attack generation utilities.
package attacks

import . "gogambit/engine/bitboard"

// MaskBishopOccupancy masks the relevant bishop occupancy bits for a given square.
// This forms a key for magic bitboards.
func MaskBishopOccupancy(sq int) Bitboard {
	occupancy := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// NE
	for r, f := tr+1, tf+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		occupancy.SetBit(r*8 + f)
	}

	// NW
	for r, f := tr+1, tf-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		occupancy.SetBit(r*8 + f)
	}

	// SE
	for r, f := tr-1, tf+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		occupancy.SetBit(r*8 + f)
	}

	// SW
	for r, f := tr-1, tf-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		occupancy.SetBit(r*8 + f)
	}

	return occupancy
}

// GenBishopAttacksOnTheFly generates possible bishop attacks for a given square and mask of blockers.
func GenBishopAttacksOnTheFly(sq int, blockers Bitboard) Bitboard {
	attacks := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// NE
	for r, f := tr+1, tf+1; r <= 7 && f <= 7; r, f = r+1, f+1 {
		attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	// NW
	for r, f := tr+1, tf-1; r <= 7 && f >= 0; r, f = r+1, f-1 {
		attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	// SE
	for r, f := tr-1, tf+1; r >= 0 && f <= 7; r, f = r-1, f+1 {
		attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	// SW
	for r, f := tr-1, tf-1; r >= 0 && f >= 0; r, f = r-1, f-1 {
		attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	return attacks
}
