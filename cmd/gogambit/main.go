package main

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	PieceBitboards[BR].Print()

	fmt.Printf("Piece: %c\n", AsciiPieces[BR])
	fmt.Printf("Piece: %c\n", UnicodePieces[BR])
	fmt.Printf("Piece: %c\n", UnicodePieces[CharToPiece['r']])
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAttacks()
	a.InitSliderAttacks(Bishop)
	a.InitSliderAttacks(Rook)
}
