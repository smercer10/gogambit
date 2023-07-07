// Package attacks provides attack generation utilities.
package attacks

import . "gogambit/engine/bitboard"

// MaskRelevantRookOccupancy masks the relevant rook occupancy bits for a given square.
func MaskRelevantRookOccupancy(sq int) Bitboard {
	mask := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// N
	for r := tr + 1; r <= 6; r++ {
		mask = mask.SetBit(r*8 + tf)
	}

	// S
	for r := tr - 1; r >= 1; r-- {
		mask = mask.SetBit(r*8 + tf)
	}

	// E
	for f := tf + 1; f <= 6; f++ {
		mask = mask.SetBit(tr*8 + f)
	}

	// W
	for f := tf - 1; f >= 1; f-- {
		mask = mask.SetBit(tr*8 + f)
	}

	return mask
}

// RookRelevantOccupancyBitCounts is a LUT with the bit count of the relevant rook occupancies for each square.
var RookRelevantOccupancyBitCounts = [64]int{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

// GenRookAttacksOnTheFly generates possible rook attacks for a given square and mask of blockers.
func GenRookAttacksOnTheFly(sq int, blockers Bitboard) Bitboard {
	attacks := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// N
	for r := tr + 1; r <= 7; r++ {
		attacks = attacks.SetBit(r*8 + tf)

		if blockers.GetBit(r*8 + tf) {
			break
		}
	}

	// S
	for r := tr - 1; r >= 0; r-- {
		attacks = attacks.SetBit(r*8 + tf)

		if blockers.GetBit(r*8 + tf) {
			break
		}
	}

	// E
	for f := tf + 1; f <= 7; f++ {
		attacks = attacks.SetBit(tr*8 + f)

		if blockers.GetBit(tr*8 + f) {
			break
		}
	}

	// W
	for f := tf - 1; f >= 0; f-- {
		attacks = attacks.SetBit(tr*8 + f)

		if blockers.GetBit(tr*8 + f) {
			break
		}
	}

	return attacks
}

// RookMagicNumbers is a LUT with a rook magic number for each square.
var RookMagicNumbers = [64]Bitboard{
	0x80008040002010,
	0xc240041000200040,
	0x2080088020001001,
	0xe80049001800800,
	0x8200100200200804,
	0x200041008018200,
	0x480008002000100,
	0x100124100088122,
	0x200800080400020,
	0x2421400220005000,
	0x4180801000a000,
	0x20801004080080,
	0x2001000800110004,
	0x21000208040100,
	0xb0020008c1020004,
	0x2000844288201,
	0x2000818000400031,
	0x414000201002,
	0x2040808020001000,
	0x4000a002200c090,
	0x1050008010090,
	0xa808002000400,
	0x202008080020100,
	0x4040020000804401,
	0xa000800080204000,
	0x100050024000e002,
	0x210700b100402002,
	0x410001100090020,
	0x2c008080040800,
	0x8a000280800400,
	0x4028400012810,
	0x80210200008844,
	0x20204000800092,
	0x40844004802002,
	0x2008042001020,
	0x1039801004800800,
	0x840800400800800,
	0xc008004800200,
	0x8000100844000102,
	0x100040082000041,
	0x440842040028002,
	0x4040200250024000,
	0x100020008080,
	0x8440400a00220010,
	0x41005e28010010,
	0x100e001144420008,
	0x8000010002008080,
	0x811840e420019,
	0x8880802110420200,
	0x80400080200080,
	0x800100020008080,
	0x1010002012090100,
	0xa03011084080100,
	0x1510020080040080,
	0x20148100400,
	0x4010040041008200,
	0x8000210010408001,
	0x2008400010850421,
	0x250040910a2001,
	0x600a078104006,
	0x4003000800100205,
	0x40050008c2040001,
	0xa020009008020104,
	0x210000408c002102,
}
