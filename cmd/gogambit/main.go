package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")

	PrintCurrentBoard()

	ParseFEN("r2q1rk1/ppp2ppp/2n1bn2/2b1p3/3pP3/3P1NPP/PPP1NPB1/R1BQ1RK1 w Qq e3 0 9")

	PrintCurrentBoard()
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAttacks()
	a.InitSliderAttacks(Bishop)
	a.InitSliderAttacks(Rook)
}
