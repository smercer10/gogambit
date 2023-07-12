// Package bitboard provides the bitboard type and relevant utilities.
package bitboard

import (
	"fmt"
	. "gogambit/engine/globals"
	"math/bits"
	"strings"
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
			if bb.IsSet(sq) {
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

// IsSet checks if the bit at the given square is set to 1.
func (bb Bitboard) IsSet(sq int) bool {
	return bb&(1<<sq) != 0
}

// CountBits returns the number of bits set to 1 in a bitboard.
func (bb Bitboard) CountBits() int {
	return bits.OnesCount64(uint64(bb))
}

// GetLSB returns the index of the least significant bit set to 1 in a bitboard.
func (bb Bitboard) GetLSB() int {
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

// PieceOcc is an array of occupancy bitboards for each piece type.
var PieceOcc = [12]Bitboard{
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
	Bitboard(0x1000000000000000), // King
}

// SideOcc is an array of occupancy bitboards for white and/or black pieces.
var SideOcc = [3]Bitboard{
	// White
	Bitboard(0x000000000000ffff),

	// Black
	Bitboard(0xffff000000000000),

	// Both
	Bitboard(0xffff00000000ffff),
}

// PieceOccC is the copy of PieceOcc.
var PieceOccC [12]Bitboard

// SideOccC is the copy of SideOcc.
var SideOccC [3]Bitboard

// PrintCurrentBoard prints the current board position and game states.
func PrintCurrentBoard() {
	fmt.Println("  +-----------------+")

	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1)

		for f := 0; f < 8; f++ {
			sq := r*8 + f
			piece := -1

			for p := WP; p <= BK; p++ {
				if PieceOcc[p].IsSet(sq) {
					piece = p
					break
				}
			}

			if piece == -1 {
				fmt.Print(". ")
			} else {
				// fmt.Printf("%c ", AsciiPieces[piece])
				fmt.Printf("%c ", UnicodePieces[piece])
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")

	fmt.Printf("\nSide to move: %s\n", Sides[SideToMove])
	fmt.Printf("En passant square: %s\n", Squares[EnPassantSq])
	fmt.Printf("Castling rights: %s\n", CastMap[CastRights])
}

// ParseFEN parses a FEN string and sets the board position and game states accordingly.
func ParseFEN(fen string) {
	// Reset all bitboards and game states
	for p := WP; p <= BK; p++ {
		PieceOcc[p] = 0x0
	}

	for s := White; s <= Both; s++ {
		SideOcc[s] = 0x0
	}

	SideToMove = White
	EnPassantSq = NA
	CastRights = 0b0000

	fenSplit := strings.Split(fen, " ")

	// Set piece/side bitboards
	r := 7
	f := 0

	for _, char := range fenSplit[0] {
		if char == '/' {
			r -= 1
			f = 0
		} else if char >= '1' && char <= '8' {
			f += int(char - '0')
		} else {
			piece := CharToPiece[byte(char)]
			sq := r*8 + f

			PieceOcc[piece] = PieceOcc[piece].SetBit(sq)
			SideOcc[Both] = SideOcc[Both].SetBit(sq)

			if piece <= WK {
				SideOcc[White] = SideOcc[White].SetBit(sq)
			} else {
				SideOcc[Black] = SideOcc[Black].SetBit(sq)
			}
			f++
		}
	}

	// Set side to move
	if len(fenSplit) > 1 && fenSplit[1] == "w" {
		SideToMove = White
	} else {
		SideToMove = Black
	}

	// Set castling rights
	if len(fenSplit) > 2 && fenSplit[2] != "-" {
		for _, char := range fenSplit[2] {
			switch char {
			case 'K':
				CastRights |= WKS
			case 'Q':
				CastRights |= WQS
			case 'k':
				CastRights |= BKS
			case 'q':
				CastRights |= BQS
			}
		}
	}

	// Set en passant square
	if len(fenSplit) > 3 && fenSplit[3] != "-" {
		EnPassantSq = CharToSquare[fenSplit[3]]
	}
}

// CopyBoard copies the current board position and game states.
func CopyBoard() {
	copy(PieceOccC[:], PieceOcc[:])
	copy(SideOccC[:], SideOcc[:])

	SideToMoveC = SideToMove
	EnPassantSqC = EnPassantSq
	CastRightsC = CastRights
}

// RestoreBoard restores the board position and game states from the copies.
func RestoreBoard() {
	copy(PieceOcc[:], PieceOccC[:])
	copy(SideOcc[:], SideOccC[:])

	SideToMove = SideToMoveC
	EnPassantSq = EnPassantSqC
	CastRights = CastRightsC
}
