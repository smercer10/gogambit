package main

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
)

func main() {
	blockers := Bitboard(0x0)
	blockers = blockers.SetBit(D2)
	blockers = blockers.SetBit(D5)
	attacks := a.GenRookAttacksOnTheFly(D4, blockers)

	fmt.Println(attacks.CountBits())

	attacks.Print()
}
