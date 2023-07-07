// Package attacks provides attack generation utilities.
package attacks

import . "gogambit/engine/bitboard"

// MaskRelevantBishopOccupancy masks the relevant bishop occupancy bits for a given square.
func MaskRelevantBishopOccupancy(sq int) Bitboard {
	mask := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// NE
	for r, f := tr+1, tf+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		mask = mask.SetBit(r*8 + f)
	}

	// NW
	for r, f := tr+1, tf-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		mask = mask.SetBit(r*8 + f)
	}

	// SE
	for r, f := tr-1, tf+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		mask = mask.SetBit(r*8 + f)
	}

	// SW
	for r, f := tr-1, tf-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		mask = mask.SetBit(r*8 + f)
	}

	return mask
}

// BishopRelevantOccupancyBitCounts is a LUT with the bit count of the relevant bishop occupancies for each square.
var BishopRelevantOccupancyBitCounts = [64]int{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}

// GenBishopAttacksOnTheFly generates possible bishop attacks for a given square and mask of blockers.
func GenBishopAttacksOnTheFly(sq int, blockers Bitboard) Bitboard {
	attacks := Bitboard(0x0)

	tr := sq / 8
	tf := sq % 8

	// NE
	for r, f := tr+1, tf+1; r <= 7 && f <= 7; r, f = r+1, f+1 {
		attacks = attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	// NW
	for r, f := tr+1, tf-1; r <= 7 && f >= 0; r, f = r+1, f-1 {
		attacks = attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	// SE
	for r, f := tr-1, tf+1; r >= 0 && f <= 7; r, f = r-1, f+1 {
		attacks = attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	// SW
	for r, f := tr-1, tf-1; r >= 0 && f >= 0; r, f = r-1, f-1 {
		attacks = attacks.SetBit(r*8 + f)

		if blockers.GetBit(r*8 + f) {
			break
		}
	}

	return attacks
}

// BishopMagicNumbers is a LUT with a bishop magic number for each square.
var BishopMagicNumbers = [64]Bitboard{
	0xb008011000811102,
	0x888210800810400,
	0x10208081041870,
	0x24440080404801,
	0x200404200000000c,
	0x4404242008020800,
	0x404208a0582002,
	0x2001288800901008,
	0x4121102081010202,
	0x1c0a00101290500,
	0x990010a0e2500,
	0x40824081004042,
	0x82a11041011200,
	0x800c82410080024,
	0x8020822029000,
	0x1011406011440,
	0xc0082148015100,
	0x840800101048d480,
	0x20200100c821102,
	0x884201202020062,
	0x2042000402310840,
	0x4200010100824a,
	0x10440e21042000,
	0x1402121043001,
	0x10080442480120,
	0x103000050c4480,
	0x242080040408022,
	0x8c004024010083,
	0x100840088802000,
	0x80880c8001100098,
	0x2048002580804,
	0x48820201008080a0,
	0x850d01004284204,
	0xc005642020020801,
	0x4040200040020,
	0x8094020080080080,
	0x2004004010040101,
	0x881100100808040,
	0x10020884820080,
	0x1002005201844200,
	0x880090380830a040,
	0x200440208512000,
	0x82020202020101,
	0x1004200808809,
	0x301400882000120,
	0x1a0008180a900,
	0x110022e04001050,
	0x1202623409200100,
	0x8000a41008040000,
	0x2004a08440001,
	0x6328460a01044022,
	0x288026042021000,
	0x80024202820006,
	0x9000403104012000,
	0x120088200840800,
	0x220010401004808,
	0x220130813042000,
	0xc0010109100231,
	0x103008100495000,
	0x602004309020a800,
	0xc000006005050404,
	0x9210002004618201,
	0x20000424880a0400,
	0x48089002820194,
}
