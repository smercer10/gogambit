package main

import (
	b "gogambit/engine/bitboard"
	c "gogambit/engine/const"
)

func main() {
	var bb b.Bitboard = 0x0

	bb = b.SetBit(bb, c.D4)

	b.Print(bb)
}
