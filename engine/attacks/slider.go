// Package attacks provides attack generation utilities.
package attacks

import (
	. "gogambit/engine/bitboard"
	"math/rand"
)

// SetOccupancy sets the occupancy combination for an attack mask at a given index.
// This can be used to generate all possible occupancy combinations for the mask.
func SetOccupancy(mask Bitboard, idx int) Bitboard {
	occupancy := Bitboard(0x0)

	for i := 0; i < mask.CountBits(); i++ {
		square := mask.GetLeastSignificantBit()

		mask = mask.ClearBit(square)

		if idx&(1<<i) != 0 {
			occupancy = occupancy.SetBit(square)
		}
	}

	return occupancy
}

// GenMagicNumCandidate generates a random uint64 with a low number of set bits.
func GenMagicNumCandidate() uint64 {
	return rand.Uint64() & rand.Uint64() & rand.Uint64()
}
