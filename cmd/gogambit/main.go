package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	occupancy := Bitboard(0x0)
	occupancy = occupancy.SetBit(E5)
	occupancy = occupancy.SetBit(A4)

	occupancy.Print()
	a.GetBishopAttacks(D4, occupancy).Print()
}

// initAll initializes all necessary LUTs.
func initAll() {
	a.InitLeaperAttacks()
	a.InitSliderAttacks(Bishop)
	a.InitSliderAttacks(Rook)
}
