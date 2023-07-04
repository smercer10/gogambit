package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
)

func main() {
	blockers := Bitboard(0x0)
	blockers = blockers.SetBit(D5)
	attacks := a.GenRookAttacksOnTheFly(H8, blockers)

	attacks.Print()
	blockers.Print()
}
