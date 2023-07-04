// Package attacks provides attack generation utilities.
package attacks

import . "gogambit/engine/bitboard"

// MaskKingAttacks generates the king attacks for a given square.
func MaskKingAttacks(sq int) Bitboard {
	attacks := Bitboard(0x0)

	bb := Bitboard(0x0)

	bb = bb.SetBit(sq)

	attacks |= bb << 8              // N
	attacks |= (bb << 7) & NotFileH // NE
	attacks |= (bb >> 1) & NotFileH // E
	attacks |= (bb >> 9) & NotFileH // SE
	attacks |= bb >> 8              // S
	attacks |= (bb >> 7) & NotFileA // SW
	attacks |= (bb << 1) & NotFileA // W
	attacks |= (bb << 9) & NotFileA // NW

	return attacks
}
