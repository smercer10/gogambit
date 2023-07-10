package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	occ := Bitboard(0x0)

	occ = occ.SetBit(C1)
	occ = occ.SetBit(C2)
	occ = occ.SetBit(G4)
	occ = occ.SetBit(G6)

	occ.Print()
	a.GetQueenAttacks(C4, occ).Print()
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAttacks()
	a.InitSliderAttacks(Bishop)
	a.InitSliderAttacks(Rook)
}
