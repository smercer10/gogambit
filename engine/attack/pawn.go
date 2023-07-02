// Package attack provides attack generation utilities.
package attack

import b "gogambit/engine/bitboard"

// MaskPawnAttacks generates the pawn attacks for a given square and side.
func MaskPawnAttacks(sq int, side int) b.Bitboard {
	attacks := b.Bitboard(0x0)

	bb := b.Bitboard(0x0)

	bb.SetBit(sq)

	if side == White {
		attacks |= (bb << 7) & b.NotFileH // Left
		attacks |= (bb << 9) & b.NotFileA // Right
	} else { // side == Black
		attacks |= (bb >> 7) & b.NotFileA // Left
		attacks |= (bb >> 9) & b.NotFileH // Right
	}

	return attacks
}
