// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

// GenPawnAttacks generates the pawn attacks for a given square and side.
func GenPawnAttacks(sq int, side int) Bitboard {
	attacks := Bitboard(0x0)

	bb := Bitboard(0x0)

	bb = bb.SetBit(sq)

	if side == White {
		attacks |= (bb << 7) & NotFileH // Left
		attacks |= (bb << 9) & NotFileA // Right
	} else { // Black
		attacks |= (bb >> 7) & NotFileA // Left
		attacks |= (bb >> 9) & NotFileH // Right
	}

	return attacks
}

// GenKnightAttacks generates the knight attacks for a given square.
func GenKnightAttacks(sq int) Bitboard {
	attacks := Bitboard(0x0)

	bb := Bitboard(0x0)

	bb = bb.SetBit(sq)

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

// GenKingAttacks generates the king attacks for a given square.
func GenKingAttacks(sq int) Bitboard {
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

// PawnAttacks is a lookup table for pawn attacks.
var PawnAttacks [2][64]Bitboard

// KnightAttacks is a lookup table for knight attacks.
var KnightAttacks [64]Bitboard

// KingAttacks is a lookup table for king attacks.
var KingAttacks [64]Bitboard

// InitLeaperAttacks initializes the lookup tables for leaper piece attacks.
func InitLeaperAttacks() {
	for sq := A1; sq <= H8; sq++ {
		PawnAttacks[White][sq] = GenPawnAttacks(sq, White)
		PawnAttacks[Black][sq] = GenPawnAttacks(sq, Black)
		KnightAttacks[sq] = GenKnightAttacks(sq)
		KingAttacks[sq] = GenKingAttacks(sq)
	}
}
