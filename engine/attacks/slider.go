// Package attacks provides attack generation utilities.
package attacks

import (
	"fmt"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
	"math/rand"
)

// SetOccupancy sets an occupancy combination for the mask of relevant occupancy bits.
func SetOccupancy(mask Bitboard, bits int, idx int) Bitboard {
	occ := Bitboard(0x0)

	for b := 0; b < bits; b++ {
		sq := mask.GetLeastSignificantBit()

		mask = mask.ClearBit(sq)

		if idx&(1<<b) != 0 {
			occ = occ.SetBit(sq)
		}
	}

	return occ
}

// GenMagicNumCand generates a random Bitboard with a small number of set bits.
func GenMagicNumCand() Bitboard {
	return Bitboard(rand.Uint64() & rand.Uint64() & rand.Uint64())
}

// FindMagicNumber finds a magic number for a bishop or rook at a given square.
func FindMagicNumber(sq int, piece int) Bitboard {
	var occ [4096]Bitboard

	var att [4096]Bitboard

	var usedAtt [4096]Bitboard

	var mask Bitboard

	var relBits int

	if piece == Bishop {
		mask = MaskRelBishopOcc(sq)
		relBits = BishopRelOccBits[sq]
	} else { // Rook
		mask = MaskRelRookOcc(sq)
		relBits = RookRelOccBits[sq]
	}

	// Range of unique occupancies (2^relBits)
	occIndices := 1 << relBits

	for idx := 0; idx < occIndices; idx++ {
		occ[idx] = SetOccupancy(mask, relBits, idx)

		if piece == Bishop {
			att[idx] = GenBishopAttOTF(sq, occ[idx])
		} else { // Rook
			att[idx] = GenRookAttOTF(sq, occ[idx])
		}
	}

	// Test magic numbers
	for tries := 0; tries < 1000000000; tries++ {
		magicNum := GenMagicNumCand()

		// Skip inappropriate magic numbers
		if ((magicNum * mask) & 0xFF00000000000000).CountBits() < 6 {
			continue
		}

		for idx := 0; idx < occIndices; idx++ {
			usedAtt[idx] = 0x0
		}

		fail := false

		// Test magic index
		for idx := 0; idx < occIndices && !fail; idx++ {
			magicIdx := int((occ[idx] * magicNum) >> (64 - relBits))

			if usedAtt[magicIdx] == 0x0 {
				usedAtt[magicIdx] = att[idx]
			} else if usedAtt[magicIdx] != att[idx] {
				fail = true
			}
		}

		if !fail {
			return magicNum
		}
	}
	fmt.Println("Failed to find magic number")

	return 0x0
}

// InitSliderAttacks initializes the necessary LUTs to get bishop or rook attacks.
func InitSliderAttacks(piece int) {
	for sq := A1; sq <= H8; sq++ {
		if piece == Bishop {
			// Init bishop relevant occ masks LUT
			BishopOccMasks[sq] = MaskRelBishopOcc(sq)

			occIndices := 1 << BishopRelOccBits[sq]

			for idx := 0; idx < occIndices; idx++ {
				occ := SetOccupancy(BishopOccMasks[sq], BishopRelOccBits[sq], idx)
				magicIdx := int((occ * BishopMagicNums[sq]) >> (64 - BishopRelOccBits[sq]))

				// Init bishop attacks LUT
				BishopAttacks[sq][magicIdx] = GenBishopAttOTF(sq, occ)
			}
		} else { // Rook
			// Init rook relevant occ masks LUT
			RookOccMasks[sq] = MaskRelRookOcc(sq)

			occIndices := 1 << RookRelOccBits[sq]

			for idx := 0; idx < occIndices; idx++ {
				occ := SetOccupancy(RookOccMasks[sq], RookRelOccBits[sq], idx)
				magicIdx := int((occ * RookMagicNums[sq]) >> (64 - RookRelOccBits[sq]))

				// Init rook attacks LUT
				RookAttacks[sq][magicIdx] = GenRookAttOTF(sq, occ)
			}
		}
	}
}

// GetQueenAttacks returns possible queen attacks for a given square and board occ.
func GetQueenAttacks(sq int, occ Bitboard) Bitboard {
	return GetBishopAttacks(sq, occ) | GetRookAttacks(sq, occ)
}
