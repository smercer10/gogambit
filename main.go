package main

import (
	b "gogambit/engine/bitboard"
	c "gogambit/engine/constants"
)

func main() {
	var bb b.Bitboard = 0x8000000000000001

	b.ClearBit(&bb, c.A1)
	b.ClearBit(&bb, c.H8)

	b.PrintBitboard(bb)
}
