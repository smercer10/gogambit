package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	m "gogambit/engine/moves"
)

func main() {
	initAll()

	ParseFEN("r3k2r/11q11ppp/p23p3/1pPQ1Q2/1P2Pp2/1B4P1/1p3PBP/Rp4K1 b kq b6 0 27")

	moves := m.MoveList{}

	m.GenMoves(&moves)

	for ct := 0; ct < moves.Count; ct++ {
		CopyBoard()

		m.MakeMove(moves.Moves[ct], AllMoves)

		PrintCurrentBoard()
		m.PrintMove(moves.Moves[ct])
		SideOcc[Black].Print()
		SideOcc[White].Print()
		SideOcc[Both].Print()

		TakeBack()
	}
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAtt()
	a.InitSliderAtt(Bishop)
	a.InitSliderAtt(Rook)
}
