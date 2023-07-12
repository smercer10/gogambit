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

	move := m.EncMove(D8, E6, WP, NA, 1, 0, 0, 0)
	move2 := m.EncMove(E2, B6, WP, WN, 1, 0, 0, 0)

	list := m.MoveList{}
	list.AddMove(move)
	list.AddMove(move2)

	m.PrintMoveList(&list)

	for sq := 0; sq < 64; sq++ {
		fmt.Printf("0x%x\n", a.FindMagicNumber(sq, Bishop))
	}
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAtt()
	a.InitSliderAtt(Bishop)
	a.InitSliderAtt(Rook)
}
