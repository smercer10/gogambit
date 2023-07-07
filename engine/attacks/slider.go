// Package attacks provides attack generation utilities.
package attacks

import (
	"fmt"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"math/rand"
)

// SetOccupancy sets an occupancy combination for the mask of relevant occupancy bits.
// This can be used to generate all possible occupancy combinations.
func SetOccupancy(mask Bitboard, maskBitCount int, idx int) Bitboard {
	occupancy := Bitboard(0x0)

	for count := 0; count < maskBitCount; count++ {
		sq := mask.GetLeastSignificantBit()

		mask = mask.ClearBit(sq)

		if idx&(1<<count) != 0 {
			occupancy = occupancy.SetBit(sq)
		}
	}

	return occupancy
}

// GenMagicNumCandidate generates a random Bitboard with a small number of set bits.
func GenMagicNumCandidate() Bitboard {
	return Bitboard(rand.Uint64() & rand.Uint64() & rand.Uint64())
}

// FindMagicNumber finds a magic number for the given piece type at the given square.
func FindMagicNumber(sq int, bishopOrRook int) Bitboard {
	var occupancies [4096]Bitboard

	var attacks [4096]Bitboard

	var usedAttacks [4096]Bitboard

	var mask Bitboard

	var relevantBitCount int

	if bishopOrRook == Bishop {
		mask = MaskRelevantBishopOccupancy(sq)
		relevantBitCount = BishopRelevantOccupancyBitCounts[sq]
	} else { // Rook
		mask = MaskRelevantRookOccupancy(sq)
		relevantBitCount = RookRelevantOccupancyBitCounts[sq]
	}

	// Range of unique occupancies (2^relevantBitCount)
	occupancyIndices := 1 << relevantBitCount

	for idx := 0; idx < occupancyIndices; idx++ {
		occupancies[idx] = SetOccupancy(mask, relevantBitCount, idx)

		if bishopOrRook == Bishop {
			attacks[idx] = GenBishopAttacksOnTheFly(sq, occupancies[idx])
		} else { // Rook
			attacks[idx] = GenRookAttacksOnTheFly(sq, occupancies[idx])
		}
	}

	// Test magic numbers
	for randomCount := 0; randomCount < 1000000000; randomCount++ {
		magicNumber := GenMagicNumCandidate()

		// Skip inappropriate magic numbers
		if ((magicNumber * mask) & 0xFF00000000000000).CountBits() < 6 {
			continue
		}

		for idx := 0; idx < occupancyIndices; idx++ {
			usedAttacks[idx] = 0x0
		}

		fail := false

		// Test magic index
		for idx := 0; idx < occupancyIndices && !fail; idx++ {
			magicIndex := int((occupancies[idx] * magicNumber) >> (64 - relevantBitCount))

			if usedAttacks[magicIndex] == 0x0 {
				usedAttacks[magicIndex] = attacks[idx]
			} else if usedAttacks[magicIndex] != attacks[idx] {
				fail = true
			}
		}

		if !fail {
			return magicNumber
		}
	}
	fmt.Println("Failed to find magic number")

	return 0x0
}
