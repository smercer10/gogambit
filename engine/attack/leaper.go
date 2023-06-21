// Package attack provides attack generation utilities.
package attack

import b "gogambit/engine/bitboard"

// PawnAttacks is a lookup table for pawn attacks.
var PawnAttacks [2][64]b.Bitboard

// KnightAttacks is a lookup table for knight attacks.
var KnightAttacks [64]b.Bitboard

// KingAttacks is a lookup table for king attacks.
var KingAttacks [64]b.Bitboard

// InitLeaperAttacks initializes the lookup tables for leaper piece attacks.
func InitLeaperAttacks() {
	for sq := b.A1; sq <= b.H8; sq++ {
		PawnAttacks[White][sq] = MaskPawnAttacks(sq, White)
		PawnAttacks[Black][sq] = MaskPawnAttacks(sq, Black)
		KnightAttacks[sq] = MaskKnightAttacks(sq)
		KingAttacks[sq] = MaskKingAttacks(sq)
	}
}
