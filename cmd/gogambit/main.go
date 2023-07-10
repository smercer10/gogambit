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

	ParseFEN("r3k1nr/1bq2ppp/p2p4/1p1P1Q2/1P6/1B4P1/4PPBP/R5K1 b kq - 0 27")

	PrintCurrentBoard()

	fmt.Printf("%t", m.IsAttacked(C7, White))
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAtt()
	a.InitSliderAtt(Bishop)
	a.InitSliderAtt(Rook)
}
