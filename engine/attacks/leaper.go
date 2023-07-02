// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
)

// PawnAttacks is a lookup table for pawn attacks.
var PawnAttacks [2][64]Bitboard

// KnightAttacks is a lookup table for knight attacks.
var KnightAttacks [64]Bitboard

// KingAttacks is a lookup table for king attacks.
var KingAttacks [64]Bitboard

// InitLeaperAttacks initializes the lookup tables for leaper piece attacks.
func InitLeaperAttacks() {
	for sq := A1; sq <= H8; sq++ {
		PawnAttacks[White][sq] = MaskPawnAttacks(sq, White)
		PawnAttacks[Black][sq] = MaskPawnAttacks(sq, Black)
		KnightAttacks[sq] = MaskKnightAttacks(sq)
		KingAttacks[sq] = MaskKingAttacks(sq)
	}
}
