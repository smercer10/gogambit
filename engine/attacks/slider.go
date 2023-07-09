// Package attacks provides attack generation utilities.
package attacks

import (
	"fmt"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"math/rand"
)

// SetOccupancy sets an occupancy combination for the mask of relevant occupancy bits.
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

// FindMagicNumber finds a magic number for a bishop or rook at a given square.
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
	for tries := 0; tries < 1000000000; tries++ {
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

// InitSliderAttacks initializes the necessary LUTs to get bishop or rook attacks.
func InitSliderAttacks(bishopOrRook int) {
	for sq := A1; sq <= H8; sq++ {
		if bishopOrRook == Bishop {
			// Init bishop relevant occupancy masks LUT
			BishopOccupancyMasks[sq] = MaskRelevantBishopOccupancy(sq)

			occupancyIndices := 1 << BishopRelevantOccupancyBitCounts[sq]

			for idx := 0; idx < occupancyIndices; idx++ {
				occupancy := SetOccupancy(BishopOccupancyMasks[sq], BishopRelevantOccupancyBitCounts[sq], idx)
				magicIndex := int((occupancy * BishopMagicNumbers[sq]) >> (64 - BishopRelevantOccupancyBitCounts[sq]))

				// Init bishop attacks LUT
				BishopAttacks[sq][magicIndex] = GenBishopAttacksOnTheFly(sq, occupancy)
			}
		} else { // Rook
			// Init rook relevant occupancy masks LUT
			RookOccupancyMasks[sq] = MaskRelevantRookOccupancy(sq)

			occupancyIndices := 1 << RookRelevantOccupancyBitCounts[sq]

			for idx := 0; idx < occupancyIndices; idx++ {
				occupancy := SetOccupancy(RookOccupancyMasks[sq], RookRelevantOccupancyBitCounts[sq], idx)
				magicIndex := int((occupancy * RookMagicNumbers[sq]) >> (64 - RookRelevantOccupancyBitCounts[sq]))

				// Init rook attacks LUT
				RookAttacks[sq][magicIndex] = GenRookAttacksOnTheFly(sq, occupancy)
			}
		}
	}
}

// GetQueenAttacks returns possible queen attacks for a given square and board occupancy.
func GetQueenAttacks(sq int, occupancy Bitboard) Bitboard {
	return GetBishopAttacks(sq, occupancy) | GetRookAttacks(sq, occupancy)
}
