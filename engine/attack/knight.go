// Package attack provides attack generation utilities.
package attack

import b "gogambit/engine/bitboard"

// MaskKnightAttacks generates the knight attacks for a given square.
func MaskKnightAttacks(sq int) b.Bitboard {
	attacks := b.Bitboard(0x0)

	var bb b.Bitboard = 0x0

	b.SetBit(&bb, sq)

	attacks |= (bb << 17) & b.NotFileA  // Up 2, left 1
	attacks |= (bb << 10) & b.NotFileAB // Up 1, left 2
	attacks |= (bb << 15) & b.NotFileH  // Up 2, right 1
	attacks |= (bb << 6) & b.NotFileGH  // Up 1, right 2
	attacks |= (bb >> 6) & b.NotFileAB  // Down 1, left 2
	attacks |= (bb >> 15) & b.NotFileA  // Down 2, left 1
	attacks |= (bb >> 10) & b.NotFileGH // Down 1, right 2
	attacks |= (bb >> 17) & b.NotFileH  // Down 2, right 1

	return attacks
}
