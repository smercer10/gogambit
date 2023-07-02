// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
)

// MaskPawnAttacks generates the pawn attacks for a given square and side.
func MaskPawnAttacks(sq int, side int) Bitboard {
	attacks := Bitboard(0x0)

	bb := Bitboard(0x0)

	bb.SetBit(sq)

	if side == White {
		attacks |= (bb << 7) & NotFileH // Left
		attacks |= (bb << 9) & NotFileA // Right
	} else { // side == Black
		attacks |= (bb >> 7) & NotFileA // Left
		attacks |= (bb >> 9) & NotFileH // Right
	}

	return attacks
}
