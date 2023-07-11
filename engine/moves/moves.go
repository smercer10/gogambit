// Package moves provides move generation utilities.
package moves

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

// IsAttacked checks if a square is currently attacked by a given side.
func IsAttacked(sq int, by int) bool {
	if by == White && (a.PawnAttacks[Black][sq]&PieceOcc[WP]) != 0 {
		return true
	}

	if by == Black && (a.PawnAttacks[White][sq]&PieceOcc[BP]) != 0 {
		return true
	}

	if (by == White && a.KnightAttacks[sq]&PieceOcc[WN] != 0) ||
		(by == Black && a.KnightAttacks[sq]&PieceOcc[BN] != 0) {
		return true
	}

	if (by == White && a.KingAttacks[sq]&PieceOcc[WK] != 0) ||
		(by == Black && a.KingAttacks[sq]&PieceOcc[BK] != 0) {
		return true
	}

	if (by == White && a.GetBishopAttacks(sq, SideOcc[Both])&PieceOcc[WB] != 0) ||
		(by == Black && a.GetBishopAttacks(sq, SideOcc[Both])&PieceOcc[BB] != 0) {
		return true
	}

	if (by == White && a.GetRookAttacks(sq, SideOcc[Both])&PieceOcc[WR] != 0) ||
		(by == Black && a.GetRookAttacks(sq, SideOcc[Both])&PieceOcc[BR] != 0) {
		return true
	}

	if (by == White && a.GetQueenAttacks(sq, SideOcc[Both])&PieceOcc[WQ] != 0) ||
		(by == Black && a.GetQueenAttacks(sq, SideOcc[Both])&PieceOcc[BQ] != 0) {
		return true
	}

	return false
}

// PrintAttacked prints a bitboard with the squares currently attacked by a given side set to 1.
func PrintAttacked(by int) {
	fmt.Println("  +-----------------+")

	bb := Bitboard(0x0)

	for r := 7; r >= 0; r-- {
		fmt.Printf("%d | ", r+1)

		for f := 0; f < 8; f++ {
			sq := r*8 + f
			if IsAttacked(sq, by) {
				fmt.Print("1 ")

				bb = bb.SetBit(sq)
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")
	fmt.Printf("\nBitboard: 0x%x\n", bb)
}

// GenMoves generates all legal moves.
func GenMoves() {
	var srcSq int

	var trgtSq int

	var bb Bitboard

	var att Bitboard

	var opp int

	for p := WP; p <= BK; p++ {
		bb = PieceOcc[p]

		if SideToMove == White {
			// Pawn moves
			if p == WP {
				for bb != 0x0 {
					srcSq = bb.GetLSB()
					trgtSq = srcSq + 8

					if trgtSq <= H8 && !SideOcc[Both].IsSet(trgtSq) {
						// Promotion
						if srcSq >= A7 && srcSq <= H7 {
							fmt.Printf("Promotion: %s%s\n", Squares[srcSq], Squares[trgtSq])
						} else { // Single push
							fmt.Printf("Single push: %s%s\n", Squares[srcSq], Squares[trgtSq])
							// Double push
							if srcSq >= A2 && srcSq <= H2 && !SideOcc[Both].IsSet(trgtSq+8) {
								fmt.Printf("Double push: %s%s\n", Squares[srcSq], Squares[trgtSq+8])
							}
						}
					}

					att = a.PawnAttacks[White][srcSq] & SideOcc[Black]

					for att != 0x0 {
						trgtSq = att.GetLSB()
						// Capture with promotion
						if srcSq >= A7 && srcSq <= H7 {
							fmt.Printf("Capture with promotion: %s%s\n", Squares[srcSq], Squares[trgtSq])
						} else { // Capture
							fmt.Printf("Capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
						}

						att = att.ClearBit(trgtSq)
					}
					// En passant capture
					if EnPassantSq != NA {
						enPassantAtt := Bitboard(a.PawnAttacks[White][srcSq] & (1 << EnPassantSq))

						if enPassantAtt != 0x0 {
							enPassantTrgt := enPassantAtt.GetLSB()
							fmt.Printf("En passant capture: %s%s\n", Squares[srcSq], Squares[enPassantTrgt])
						}
					}

					bb = bb.ClearBit(srcSq)
				}
			}

			if p == WK {
				// Castle kingside
				if CastlingRights&WKS != 0 {
					if !SideOcc[Both].IsSet(F1) && !SideOcc[Both].IsSet(G1) {
						if !IsAttacked(E1, Black) && !IsAttacked(F1, Black) {
							fmt.Println("Castle kingside")
						}
					}
				}
				// Castle queenside
				if CastlingRights&WQS != 0 {
					if !SideOcc[Both].IsSet(D1) && !SideOcc[Both].IsSet(C1) && !SideOcc[Both].IsSet(B1) {
						if !IsAttacked(E1, Black) && !IsAttacked(D1, Black) {
							fmt.Println("Castle queenside")
						}
					}
				}
			}
		} else { // Black
			// Pawn moves
			if p == BP {
				for bb != 0x0 {
					srcSq = bb.GetLSB()
					trgtSq = srcSq - 8

					if trgtSq >= A1 && !SideOcc[Both].IsSet(trgtSq) {
						// Promotion
						if srcSq >= A2 && srcSq <= H2 {
							fmt.Printf("Promotion: %s%s\n", Squares[srcSq], Squares[trgtSq])
						} else { // Single push
							fmt.Printf("Single push: %s%s\n", Squares[srcSq], Squares[trgtSq])
							// Double push
							if srcSq >= A7 && srcSq <= H7 && !SideOcc[Both].IsSet(trgtSq-8) {
								fmt.Printf("Double push: %s%s\n", Squares[srcSq], Squares[trgtSq-8])
							}
						}
					}

					att = a.PawnAttacks[Black][srcSq] & SideOcc[White]

					for att != 0x0 {
						trgtSq = att.GetLSB()

						// Capture with promotion
						if srcSq >= A2 && srcSq <= H2 {
							fmt.Printf("Capture with promotion: %s%s\n", Squares[srcSq], Squares[trgtSq])
						} else { // Capture
							fmt.Printf("Capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
						}

						att = att.ClearBit(trgtSq)
					}
					// En passant capture
					if EnPassantSq != NA {
						enPassantAtt := Bitboard(a.PawnAttacks[Black][srcSq] & (1 << EnPassantSq))

						if enPassantAtt != 0x0 {
							enPassantTrgt := enPassantAtt.GetLSB()
							fmt.Printf("En passant capture: %s%s\n", Squares[srcSq], Squares[enPassantTrgt])
						}
					}

					bb = bb.ClearBit(srcSq)
				}
			}

			if p == BK {
				// Castle kingside
				if CastlingRights&BKS != 0 {
					if !SideOcc[Both].IsSet(F8) && !SideOcc[Both].IsSet(G8) {
						if !IsAttacked(E8, White) && !IsAttacked(F8, White) {
							fmt.Println("Castle kingside")
						}
					}
				}
				// Castle queenside
				if CastlingRights&BQS != 0 {
					if !SideOcc[Both].IsSet(D8) && !SideOcc[Both].IsSet(C8) && !SideOcc[Both].IsSet(B8) {
						if !IsAttacked(E8, White) && !IsAttacked(D8, White) {
							fmt.Println("Castle queenside")
						}
					}
				}
			}
		}

		if SideToMove == White {
			opp = Black
		} else { // Black
			opp = White
		}

		// Knight moves
		if SideToMove == White && p == WN || SideToMove == Black && p == BN {
			for bb != 0x0 {
				srcSq = bb.GetLSB()

				att = a.KnightAttacks[srcSq] & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						fmt.Printf("Quiet knight move: %s%s\n", Squares[srcSq], Squares[trgtSq])
					} else { // Capture
						fmt.Printf("Knight capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
					}

					att = att.ClearBit(trgtSq)
				}

				bb = bb.ClearBit(srcSq)
			}
		}

		// King moves
		if SideToMove == White && p == WK || SideToMove == Black && p == BK {
			for bb != 0x0 {
				srcSq = bb.GetLSB()

				att = a.KingAttacks[srcSq] & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						fmt.Printf("Quiet king move: %s%s\n", Squares[srcSq], Squares[trgtSq])
					} else { // Capture
						fmt.Printf("King capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
					}

					att = att.ClearBit(trgtSq)
				}

				bb = bb.ClearBit(srcSq)
			}
		}

		// Bishop moves
		if SideToMove == White && p == WB || SideToMove == Black && p == BB {
			for bb != 0x0 {
				srcSq = bb.GetLSB()

				att = a.GetBishopAttacks(srcSq, SideOcc[Both]) & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						fmt.Printf("Quiet bishop move: %s%s\n", Squares[srcSq], Squares[trgtSq])
					} else { // Capture
						fmt.Printf("Bishop capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
					}

					att = att.ClearBit(trgtSq)
				}

				bb = bb.ClearBit(srcSq)
			}
		}

		// Rook moves
		if SideToMove == White && p == WR || SideToMove == Black && p == BR {
			for bb != 0x0 {
				srcSq = bb.GetLSB()

				att = a.GetRookAttacks(srcSq, SideOcc[Both]) & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						fmt.Printf("Quiet rook move: %s%s\n", Squares[srcSq], Squares[trgtSq])
					} else { // Capture
						fmt.Printf("Rook capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
					}

					att = att.ClearBit(trgtSq)
				}

				bb = bb.ClearBit(srcSq)
			}
		}

		// Queen moves
		if SideToMove == White && p == WQ || SideToMove == Black && p == BQ {
			for bb != 0x0 {
				srcSq = bb.GetLSB()

				att = a.GetQueenAttacks(srcSq, SideOcc[Both]) & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						fmt.Printf("Quiet queen move: %s%s\n", Squares[srcSq], Squares[trgtSq])
					} else { // Capture
						fmt.Printf("Queen capture: %s%s\n", Squares[srcSq], Squares[trgtSq])
					}

					att = att.ClearBit(trgtSq)
				}

				bb = bb.ClearBit(srcSq)
			}
		}
	}
}
