package main

import (
	a "gogambit/engine/attacks"
	. "gogambit/engine/enums"
)

func main() {
	bb := a.MaskRookOccupancy(H8)

	bb.Print()
}
