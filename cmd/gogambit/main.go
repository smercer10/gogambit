package main

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
)

func main() {
	fmt.Println(a.GenMagicNumCandidate())
	fmt.Printf("%064b\n", a.GenMagicNumCandidate())
	Bitboard(a.GenMagicNumCandidate()).Print()
}
