package main

import (
	"fmt"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

func main() {
	bb := Bitboard(0x0)
	bb = bb.SetBit(H7)
	bb = bb.SetBit(D8)

	bb.Print()
	fmt.Printf("%b\n", bb)
	fmt.Println(Squares[bb.GetLeastSignificantBit()])
}
