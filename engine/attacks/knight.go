// Package attacks provides attack generation utilities.
package attacks

import . "gogambit/engine/bitboard"

// MaskKnightAttacks generates the knight attacks for a given square.
func MaskKnightAttacks(sq int) Bitboard {
	attacks := Bitboard(0x0)

	bb := Bitboard(0x0)

	bb.SetBit(sq)

	attacks |= (bb << 17) & NotFileA  // N2W1
	attacks |= (bb << 10) & NotFileAB // N1W2
	attacks |= (bb << 15) & NotFileH  // N2E1
	attacks |= (bb << 6) & NotFileGH  // N1E2
	attacks |= (bb >> 6) & NotFileAB  // S1W2
	attacks |= (bb >> 15) & NotFileA  // S2W1
	attacks |= (bb >> 10) & NotFileGH // S1E2
	attacks |= (bb >> 17) & NotFileH  // S2E1

	return attacks
}
