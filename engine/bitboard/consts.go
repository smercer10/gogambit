// Package bitboard provides the bitboard type and relevant utilities.
package bitboard

// NotFileA is a bitboard with only the A file cleared.
const NotFileA = Bitboard(0xfefefefefefefefe)

// NotFileH is a bitboard with only the H file cleared.
const NotFileH = Bitboard(0x7f7f7f7f7f7f7f7f)

// NotFileAB is a bitboard with only the A and B files cleared.
const NotFileAB = Bitboard(0xfcfcfcfcfcfcfcfc)

// NotFileGH is a bitboard with only the G and H files cleared.
const NotFileGH = Bitboard(0x3f3f3f3f3f3f3f3f)
