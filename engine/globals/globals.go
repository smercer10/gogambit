// Package globals provides common variables and constants used throughout the engine.
package globals

// Enum for squares on the board.
const (
	A1 = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NA
)

// Enum for player sides.
const (
	White = iota
	Black
	Both
)

// Enum for bishop or rook.
const (
	Rook = iota
	Bishop
)

// Enum for piece types.
const (
	WP = iota
	WN
	WB
	WR
	WQ
	WK
	BP
	BN
	BB
	BR
	BQ
	BK
)

// Enum for castling rights (bit flags).
const (
	WhiteKingside = 1 << iota
	WhiteQueenside
	BlackKingside
	BlackQueenside
)

// Squares is an array of board coordinates as strings.
var Squares = [65]string{
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"N/A",
}

// Sides is an array for player sides.
var Sides = [2]string{
	"White",
	"Black",
}

// AsciiPieces is an array of ASCII characters representing pieces.
var AsciiPieces = [12]byte{
	'P', 'N', 'B', 'R', 'Q', 'K',
	'p', 'n', 'b', 'r', 'q', 'k',
}

// UnicodePieces is an array of Unicode characters representing pieces.
var UnicodePieces = [12]rune{
	'♙', '♘', '♗', '♖', '♕', '♔',
	'♟', '♞', '♝', '♜', '♛', '♚',
}

// CharToPiece maps ASCII characters to piece types.
var CharToPiece = map[byte]int{
	'P': WP,
	'N': WN,
	'B': WB,
	'R': WR,
	'Q': WQ,
	'K': WK,
	'p': BP,
	'n': BN,
	'b': BB,
	'r': BR,
	'q': BQ,
	'k': BK,
}

// CharToSquare maps string coordinates to the Squares enum.
var CharToSquare = map[string]int{
	"a1": A1, "b1": B1, "c1": C1, "d1": D1, "e1": E1, "f1": F1, "g1": G1, "h1": H1,
	"a2": A2, "b2": B2, "c2": C2, "d2": D2, "e2": E2, "f2": F2, "g2": G2, "h2": H2,
	"a3": A3, "b3": B3, "c3": C3, "d3": D3, "e3": E3, "f3": F3, "g3": G3, "h3": H3,
	"a4": A4, "b4": B4, "c4": C4, "d4": D4, "e4": E4, "f4": F4, "g4": G4, "h4": H4,
	"a5": A5, "b5": B5, "c5": C5, "d5": D5, "e5": E5, "f5": F5, "g5": G5, "h5": H5,
	"a6": A6, "b6": B6, "c6": C6, "d6": D6, "e6": E6, "f6": F6, "g6": G6, "h6": H6,
	"a7": A7, "b7": B7, "c7": C7, "d7": D7, "e7": E7, "f7": F7, "g7": G7, "h7": H7,
	"a8": A8, "b8": B8, "c8": C8, "d8": D8, "e8": E8, "f8": F8, "g8": G8, "h8": H8,
}

// SideToMove is the current side to move.
var SideToMove int = White

// EnPassantSquare is the current en passant square.
var EnPassantSquare int = NA

// CastlingRights is the current castling rights (bit flags).
var CastlingRights int = WhiteKingside | WhiteQueenside | BlackKingside | BlackQueenside

// CastlingRightsMap maps castling rights to an ASCII representation.
var CastlingRightsMap = map[int]string{
	WhiteKingside | WhiteQueenside | BlackKingside | BlackQueenside: "KQkq",
	WhiteKingside | WhiteQueenside | BlackKingside:                  "KQk-",
	WhiteKingside | WhiteQueenside | BlackQueenside:                 "KQ-q",
	WhiteKingside | WhiteQueenside:                                  "KQ--",
	WhiteKingside | BlackKingside | BlackQueenside:                  "K-kq",
	WhiteKingside | BlackKingside:                                   "K-k-",
	WhiteKingside | BlackQueenside:                                  "K--q",
	WhiteKingside:                                                   "K---",
	WhiteQueenside | BlackKingside | BlackQueenside:                 "-Qkq",
	WhiteQueenside | BlackKingside:                                  "-Qk-",
	WhiteQueenside | BlackQueenside:                                 "-Q-q",
	WhiteQueenside:                                                  "-Q--",
	BlackKingside | BlackQueenside:                                  "--kq",
	BlackKingside:                                                   "--k-",
	BlackQueenside:                                                  "---q",
	0b000:                                                           "----",
}
