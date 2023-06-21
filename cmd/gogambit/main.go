package main

import (
	a "gogambit/engine/attack"
	b "gogambit/engine/bitboard"
)

func main() {
	a.InitLeaperAttacks()

	for sq := b.A1; sq <= b.H8; sq++ {
		b.PrintBitboard(a.KingAttacks[sq])
	}
}
