// Package attacks provides attack generation utilities.
package attacks

import b "gogambit/engine/bitboard"

// MaskPawnAttacks generates the pawn attacks for a given square and side.
func MaskPawnAttacks(sq int, side int) b.Bitboard {
	var attacks b.Bitboard = 0x0

	var bb b.Bitboard = 0x0

	b.SetBit(&bb, sq)

	if side == White {
		attacks |= (bb << 7) & b.NotFileH // Left
		attacks |= (bb << 9) & b.NotFileA // Right
	} else { // side == Black
		attacks |= (bb >> 7) & b.NotFileA // Left
		attacks |= (bb >> 9) & b.NotFileH // Right
	}

	return attacks
}
