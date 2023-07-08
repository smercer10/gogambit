// Package bitboard provides the bitboard type and relevant utilities.
package bitboard

import (
	"fmt"
	"math/bits"
)

// Bitboard is a 64-bit unsigned integer used to represent a chess board.
type Bitboard uint64

// Print prints a bitboard in a human-readable format along with its hexadecimal representation.
func (bb Bitboard) Print() {
	fmt.Println("  +-----------------+")

	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1)

		for f := 0; f < 8; f++ {
			sq := r*8 + f
			if bb.GetBit(sq) {
				fmt.Print("1 ")
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

// SetBit returns a copy of a bitboard with the bit at the given square set to 1.
func (bb Bitboard) SetBit(sq int) Bitboard {
	return bb | (1 << sq)
}

// ClearBit returns a copy of a bitboard with the bit at the given square set to 0.
func (bb Bitboard) ClearBit(sq int) Bitboard {
	return bb &^ (1 << sq)
}

// GetBit returns the bit at the given square.
func (bb Bitboard) GetBit(sq int) bool {
	return bb&(1<<sq) != 0
}

// CountBits returns the number of bits set to 1 in a bitboard.
func (bb Bitboard) CountBits() int {
	return bits.OnesCount64(uint64(bb))
}

// GetLeastSignificantBit returns the index of the least significant bit set to 1 in a bitboard.
func (bb Bitboard) GetLeastSignificantBit() int {
	if bb == 0 {
		return -1
	}

	return bits.TrailingZeros64(uint64(bb))
}

// NotFileA is a bitboard with only the A file cleared.
const NotFileA = Bitboard(0xfefefefefefefefe)

// NotFileH is a bitboard with only the H file cleared.
const NotFileH = Bitboard(0x7f7f7f7f7f7f7f7f)

// NotFileAB is a bitboard with only the A and B files cleared.
const NotFileAB = Bitboard(0xfcfcfcfcfcfcfcfc)

// NotFileGH is a bitboard with only the G and H files cleared.
const NotFileGH = Bitboard(0x3f3f3f3f3f3f3f3f)

// PieceBitboards is an array of bitboards for each piece type.
var PieceBitboards = [12]Bitboard{
	// White
	Bitboard(0xff00), // Pawns
	Bitboard(0x42),   // Knights
	Bitboard(0x24),   // Bishops
	Bitboard(0x81),   // Rooks
	Bitboard(0x8),    // Queen
	Bitboard(0x10),   // King

	// Black
	Bitboard(0x00ff000000000000), // Pawns
	Bitboard(0x4200000000000000), // Knights
	Bitboard(0x2400000000000000), // Bishops
	Bitboard(0x8100000000000000), // Rooks
	Bitboard(0x800000000000000),  // Queen
	Bitboard(0x800000000000000),  // King
}

// OccupancyBitboards is an array of occupancy bitboards for white and/or black pieces.
var OccupancyBitboards = [3]Bitboard{
	// White pieces
	Bitboard(0x000000000000ffff), // Occupied squares

	// Black pieces
	Bitboard(0xffff000000000000), // Occupied squares

	// Both pieces
	Bitboard(0xffff00000000ffff), // Occupied squares
}
