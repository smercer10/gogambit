// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

// GenPawnAttacks generates the pawn attacks for a given square and side.
func GenPawnAttacks(sq, side int) Bitboard {
	att, bb := Bitboard(0x0), Bitboard(0x0)

	bb = bb.SetBit(sq)

	if side == White {
		att |= (bb << 7) & NotFileH // Left
		att |= (bb << 9) & NotFileA // Right
	} else { // Black
		att |= (bb >> 7) & NotFileA // Left
		att |= (bb >> 9) & NotFileH // Right
	}

	return att
}

// GenKnightAttacks generates the knight attacks for a given square.
func GenKnightAttacks(sq int) Bitboard {
	att, bb := Bitboard(0x0), Bitboard(0x0)

	bb = bb.SetBit(sq)

	att |= (bb << 17) & NotFileA  // N2W1
	att |= (bb << 10) & NotFileAB // N1W2
	att |= (bb << 15) & NotFileH  // N2E1
	att |= (bb << 6) & NotFileGH  // N1E2
	att |= (bb >> 6) & NotFileAB  // S1W2
	att |= (bb >> 15) & NotFileA  // S2W1
	att |= (bb >> 10) & NotFileGH // S1E2
	att |= (bb >> 17) & NotFileH  // S2E1

	return att
}

// GenKingAttacks generates the king attacks for a given square.
func GenKingAttacks(sq int) Bitboard {
	att, bb := Bitboard(0x0), Bitboard(0x0)

	bb = bb.SetBit(sq)

	att |= bb << 8              // N
	att |= (bb << 7) & NotFileH // NE
	att |= (bb >> 1) & NotFileH // E
	att |= (bb >> 9) & NotFileH // SE
	att |= bb >> 8              // S
	att |= (bb >> 7) & NotFileA // SW
	att |= (bb << 1) & NotFileA // W
	att |= (bb << 9) & NotFileA // NW

	return att
}

// PawnAttacks is a LUT for pawn attacks.
var PawnAttacks [2][64]Bitboard

// KnightAttacks is a LUT for knight attacks.
var KnightAttacks [64]Bitboard

// KingAttacks is a LUT for king attacks.
var KingAttacks [64]Bitboard

// InitLeaperAtt initializes the lookup tables for leaper piece attacks.
func InitLeaperAtt() {
	for sq := A1; sq <= H8; sq++ {
		PawnAttacks[White][sq] = GenPawnAttacks(sq, White)
		PawnAttacks[Black][sq] = GenPawnAttacks(sq, Black)
		KnightAttacks[sq] = GenKnightAttacks(sq)
		KingAttacks[sq] = GenKingAttacks(sq)
	}
}
