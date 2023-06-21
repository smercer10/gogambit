// Package bitboard provides bitboard utilities.
package bitboard

import "fmt"

// Bitboard is a 64-bit unsigned integer used to represent a chess board.
type Bitboard uint64

// PrintBitboard prints a bitboard in a human-readable format along with its hexadecimal representation.
func PrintBitboard(bb Bitboard) {
	fmt.Println("  +-----------------+")

	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1)

		for f := 0; f < 8; f++ {
			sq := r*8 + f
			if GetBit(bb, sq) {
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

// SetBit sets the bit at the given square to 1.
func SetBit(bb *Bitboard, sq int) {
	*bb |= 1 << uint(sq)
}

// ClearBit sets the bit at the given square to 0.
func ClearBit(bb *Bitboard, sq int) {
	*bb &= ^(1 << uint(sq))
}

// GetBit returns the bit at the given square.
func GetBit(bb Bitboard, sq int) bool {
	return (bb & (1 << uint(sq))) != 0
}
