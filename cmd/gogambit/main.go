package main

import (
	a "gogambit/engine/attacks"
	b "gogambit/engine/bitboard"
)

func main() {
	a.InitLeaperAttacks()

	for sq := b.A1; sq <= b.H8; sq++ {
		b.PrintBitboard(a.KnightAttacks[sq])
	}
}
