// Package moves provides move generation utilities.
package moves

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

// IsAttacked checks if a square is currently attacked by a given side.
func IsAttacked(sq int, by int) bool {
	if by == White && (a.PawnAttacks[Black][sq]&PieceOcc[WP]) != 0 {
		return true
	}

	if by == Black && (a.PawnAttacks[White][sq]&PieceOcc[BP]) != 0 {
		return true
	}

	if (by == White && a.KnightAttacks[sq]&PieceOcc[WN] != 0) ||
		(by == Black && a.KnightAttacks[sq]&PieceOcc[BN] != 0) {
		return true
	}

	if (by == White && a.KingAttacks[sq]&PieceOcc[WK] != 0) ||
		(by == Black && a.KingAttacks[sq]&PieceOcc[BK] != 0) {
		return true
	}

	if (by == White && a.GetBishopAttacks(sq, SideOcc[Both])&PieceOcc[WB] != 0) ||
		(by == Black && a.GetBishopAttacks(sq, SideOcc[Both])&PieceOcc[BB] != 0) {
		return true
	}

	if (by == White && a.GetRookAttacks(sq, SideOcc[Both])&PieceOcc[WR] != 0) ||
		(by == Black && a.GetRookAttacks(sq, SideOcc[Both])&PieceOcc[BR] != 0) {
		return true
	}

	if (by == White && a.GetQueenAttacks(sq, SideOcc[Both])&PieceOcc[WQ] != 0) ||
		(by == Black && a.GetQueenAttacks(sq, SideOcc[Both])&PieceOcc[BQ] != 0) {
		return true
	}

	return false
}

// PrintAttacked prints a bitboard with the squares currently attacked by a given side set to 1.
func PrintAttacked(by int) {
	fmt.Println("  +-----------------+")

	bb := Bitboard(0x0)

	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1)

		for f := 0; f < 8; f++ {
			sq := r*8 + f
			if IsAttacked(sq, by) {
				fmt.Print("1 ")

				bb = bb.SetBit(sq)
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")
	fmt.Printf("\nBitboard: 0x%x\n", bb)
}
