package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	ParseFEN("r3k1nr/11q11ppp/p3p3/1pP11Q2/1P2Pp2/1B4P1/1p3PBP/Rp4K1 w kq b6 0 27")
	PrintCurrentBoard()

	CopyBoard()

	ParseFEN("r5k1/1p1b1ppp/1q1p4/3Pp3/1P2P3/2P2N2/2Q2PPP/R4RK1 b - - 0 1")
	PrintCurrentBoard()

	RestoreBoard()

	PrintCurrentBoard()
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAtt()
	a.InitSliderAtt(Bishop)
	a.InitSliderAtt(Rook)
}
