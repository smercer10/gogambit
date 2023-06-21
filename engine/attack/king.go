// Package attack provides attack generation utilities.
package attack

import b "gogambit/engine/bitboard"

// MaskKingAttacks generates the king attacks for a given square.
func MaskKingAttacks(sq int) b.Bitboard {
	attacks := b.Bitboard(0x0)

	var bb b.Bitboard = 0x0

	b.SetBit(&bb, sq)

	attacks |= bb << 8              // N
	attacks |= bb << 7 & b.NotFileH // NE
	attacks |= bb >> 1 & b.NotFileH // E
	attacks |= bb >> 9 & b.NotFileH // SE
	attacks |= bb >> 8              // S
	attacks |= bb >> 7 & b.NotFileA // SW
	attacks |= bb << 1 & b.NotFileA // W
	attacks |= bb << 9 & b.NotFileA // NW

	return attacks
}
