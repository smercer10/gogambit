package main

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/globals"
)

func main() {
	initAll()

	for sq := A1; sq <= H8; sq++ {
		fmt.Printf("0x%x,\n", a.FindMagicNumber(sq, Bishop))
	}
}

// initAll initializes all necessary variables.
func initAll() {
	a.InitLeaperAttacks()
}
