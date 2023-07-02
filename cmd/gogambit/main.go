package main

import (
	a "gogambit/engine/attack"
	b "gogambit/engine/bitboard"
)

func main() {
	bb := a.MaskRookOccupancy(b.H8)

	bb.Print()
}
