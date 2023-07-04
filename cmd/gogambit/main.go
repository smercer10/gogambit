package main

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/globals"
)

func main() {
	mask := a.MaskBishopOccupancy(C3)

	for i := 0; i < 10; i++ {
		occupancy := a.SetOccupancy(mask, i)
		occupancy.Print()
		fmt.Println(i)
	}
}
