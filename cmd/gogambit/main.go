package main

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	m "gogambit/engine/moves"
)

func main() {
	initAll()

	ParseFEN("r3k1nr/1bqP1ppp/p3p3/1pP11Q2/1P2Pp2/1B4P1/1p3PBP/Rp4K1 w kq b6 0 27")

	PrintCurrentBoard()

	fmt.Println()

	m.GenMoves()
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAtt()
	a.InitSliderAtt(Bishop)
	a.InitSliderAtt(Rook)
}
