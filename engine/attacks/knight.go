// Package attacks provides attack generation utilities.
package attacks

import b "gogambit/engine/bitboard"

// MaskKnightAttacks generates the knight attacks for a given square.
func MaskKnightAttacks(sq int) b.Bitboard {
	attacks := b.Bitboard(0x0)

	var bb b.Bitboard = 0x0

	b.SetBit(&bb, sq)

	attacks |= (bb << 17) & b.NotFileA  // Up 2, Left 1
	attacks |= (bb << 10) & b.NotFileAB // Up 1, Left 2
	attacks |= (bb << 15) & b.NotFileH  // Up 2, Right 1
	attacks |= (bb << 6) & b.NotFileGH  // Up 1, Right 2
	attacks |= (bb >> 6) & b.NotFileAB  // Down 1, Left 2
	attacks |= (bb >> 15) & b.NotFileA  // Down 2, Left 1
	attacks |= (bb >> 10) & b.NotFileGH // Down 1, Right 2
	attacks |= (bb >> 17) & b.NotFileH  // Down 2, Right 1

	return attacks
}
