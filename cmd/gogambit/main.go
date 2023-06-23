package main

import (
	a "gogambit/engine/attack"
	b "gogambit/engine/bitboard"
)

func main() {
	b.PrintBitboard(a.MaskRookOccupancy(b.H8))
}
