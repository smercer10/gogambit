package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	occupancy := Bitboard(0x0)

	occupancy = occupancy.SetBit(C1)
	occupancy = occupancy.SetBit(C2)
	occupancy = occupancy.SetBit(G4)
	occupancy = occupancy.SetBit(G6)

	occupancy.Print()
	a.GetQueenAttacks(C4, occupancy).Print()
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAttacks()
	a.InitSliderAttacks(Bishop)
	a.InitSliderAttacks(Rook)
}
