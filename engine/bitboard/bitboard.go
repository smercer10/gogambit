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
