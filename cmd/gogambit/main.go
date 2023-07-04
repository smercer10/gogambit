package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
)

func main() {
	blockers := Bitboard(0x0)
	attack := a.GenBishopAttacksOnTheFly(H8, blockers)

	attack.Print()
	blockers.Print()
}
