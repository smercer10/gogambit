package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/enums"
)

func main() {
	block := Bitboard(0x0)

	attack := a.GenBishopAttacksOnTheFly(H8, block)

	attack.Print()
	block.Print()
}
