// Package attack provides attack generation utilities.
package attack

import b "gogambit/engine/bitboard"

// MaskKnightAttacks generates the knight attacks for a given square.
func MaskKnightAttacks(sq int) b.Bitboard {
	attacks := b.Bitboard(0x0)

	bb := b.Bitboard(0x0)

	bb.SetBit(sq)

	attacks |= (bb << 17) & b.NotFileA  // N2W1
	attacks |= (bb << 10) & b.NotFileAB // N1W2
	attacks |= (bb << 15) & b.NotFileH  // N2E1
	attacks |= (bb << 6) & b.NotFileGH  // N1E2
	attacks |= (bb >> 6) & b.NotFileAB  // S1W2
	attacks |= (bb >> 15) & b.NotFileA  // S2W1
	attacks |= (bb >> 10) & b.NotFileGH // S1E2
	attacks |= (bb >> 17) & b.NotFileH  // S2E1

	return attacks
}
