// Package moves provides move generation utilities.
package moves

import (
	"fmt"
	a "gogambit/engine/attacks"
	. "gogambit/engine/bitboard"
	. "gogambit/engine/globals"
)

// IsAttacked checks if a square is currently attacked by a given side.
func IsAttacked(sq, by int) bool {
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

	if (by == White && a.GetBishopAtt(sq, SideOcc[Both])&PieceOcc[WB] != 0) ||
		(by == Black && a.GetBishopAtt(sq, SideOcc[Both])&PieceOcc[BB] != 0) {
		return true
	}

	if (by == White && a.GetRookAtt(sq, SideOcc[Both])&PieceOcc[WR] != 0) ||
		(by == Black && a.GetRookAtt(sq, SideOcc[Both])&PieceOcc[BR] != 0) {
		return true
	}

	if (by == White && a.GetQueenAtt(sq, SideOcc[Both])&PieceOcc[WQ] != 0) ||
		(by == Black && a.GetQueenAtt(sq, SideOcc[Both])&PieceOcc[BQ] != 0) {
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

// GenMoves generates all legal moves for the current board position and side to move.
func GenMoves(moves *MoveList) {
	moves.Count = 0

	var srcSq, trgtSq int

	var bb, att Bitboard

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
							moves.AddMove(EncMove(srcSq, trgtSq, p, WN, 0, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, WB, 0, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, WR, 0, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, WQ, 0, 0, 0, 0))
						} else { // Single push
							moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
							// Double push
							if srcSq >= A2 && srcSq <= H2 && !SideOcc[Both].IsSet(trgtSq+8) {
								moves.AddMove(EncMove(srcSq, trgtSq+8, p, 0, 0, 1, 0, 0))
							}
						}
					}

					att = a.PawnAttacks[White][srcSq] & SideOcc[Black]

					for att != 0x0 {
						trgtSq = att.GetLSB()
						// Capture with promotion
						if srcSq >= A7 && srcSq <= H7 {
							moves.AddMove(EncMove(srcSq, trgtSq, p, WN, 1, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, WB, 1, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, WR, 1, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, WQ, 1, 0, 0, 0))
						} else { // Capture
							moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
						}

						att = att.ClearBit(trgtSq)
					}
					// En passant capture
					if EnPassantSq != NA {
						enPassantAtt := Bitboard(a.PawnAttacks[White][srcSq] & (1 << EnPassantSq))

						if enPassantAtt != 0x0 {
							enPassantTrgt := enPassantAtt.GetLSB()
							moves.AddMove(EncMove(srcSq, enPassantTrgt, p, 0, 1, 0, 1, 0))
						}
					}

					bb = bb.ClearBit(srcSq)
				}
			}

			if p == WK {
				// Castle kingside
				if CastRights&WKS != 0 {
					if !SideOcc[Both].IsSet(F1) && !SideOcc[Both].IsSet(G1) {
						if !IsAttacked(E1, Black) && !IsAttacked(F1, Black) {
							moves.AddMove(EncMove(E1, G1, p, 0, 0, 0, 0, 1))
						}
					}
				}
				// Castle queenside
				if CastRights&WQS != 0 {
					if !SideOcc[Both].IsSet(D1) && !SideOcc[Both].IsSet(C1) && !SideOcc[Both].IsSet(B1) {
						if !IsAttacked(E1, Black) && !IsAttacked(D1, Black) {
							moves.AddMove(EncMove(E1, C1, p, 0, 0, 0, 0, 1))
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
							moves.AddMove(EncMove(srcSq, trgtSq, p, BN, 0, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, BB, 0, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, BR, 0, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, BQ, 0, 0, 0, 0))
						} else { // Single push
							moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
							// Double push
							if srcSq >= A7 && srcSq <= H7 && !SideOcc[Both].IsSet(trgtSq-8) {
								moves.AddMove(EncMove(srcSq, trgtSq-8, p, 0, 0, 1, 0, 0))
							}
						}
					}

					att = a.PawnAttacks[Black][srcSq] & SideOcc[White]

					for att != 0x0 {
						trgtSq = att.GetLSB()

						// Capture with promotion
						if srcSq >= A2 && srcSq <= H2 {
							moves.AddMove(EncMove(srcSq, trgtSq, p, BN, 1, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, BB, 1, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, BR, 1, 0, 0, 0))
							moves.AddMove(EncMove(srcSq, trgtSq, p, BQ, 1, 0, 0, 0))
						} else { // Capture
							moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
						}

						att = att.ClearBit(trgtSq)
					}
					// En passant capture
					if EnPassantSq != NA {
						enPassantAtt := Bitboard(a.PawnAttacks[Black][srcSq] & (1 << EnPassantSq))

						if enPassantAtt != 0x0 {
							enPassantTrgt := enPassantAtt.GetLSB()
							moves.AddMove(EncMove(srcSq, enPassantTrgt, p, 0, 1, 0, 1, 0))
						}
					}

					bb = bb.ClearBit(srcSq)
				}
			}

			if p == BK {
				// Castle kingside
				if CastRights&BKS != 0 {
					if !SideOcc[Both].IsSet(F8) && !SideOcc[Both].IsSet(G8) {
						if !IsAttacked(E8, White) && !IsAttacked(F8, White) {
							moves.AddMove(EncMove(E8, G8, p, 0, 0, 0, 0, 1))
						}
					}
				}
				// Castle queenside
				if CastRights&BQS != 0 {
					if !SideOcc[Both].IsSet(D8) && !SideOcc[Both].IsSet(C8) && !SideOcc[Both].IsSet(B8) {
						if !IsAttacked(E8, White) && !IsAttacked(D8, White) {
							moves.AddMove(EncMove(E8, C8, p, 0, 0, 0, 0, 1))
						}
					}
				}
			}
		}

		opp := SideToMove ^ 1

		// Knight moves
		if SideToMove == White && p == WN || SideToMove == Black && p == BN {
			for bb != 0x0 {
				srcSq = bb.GetLSB()

				att = a.KnightAttacks[srcSq] & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
					} else { // Capture
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
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
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
					} else { // Capture
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
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

				att = a.GetBishopAtt(srcSq, SideOcc[Both]) & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
					} else { // Capture
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
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

				att = a.GetRookAtt(srcSq, SideOcc[Both]) & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
					} else { // Capture
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
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

				att = a.GetQueenAtt(srcSq, SideOcc[Both]) & ^SideOcc[SideToMove]

				for att != 0x0 {
					trgtSq = att.GetLSB()

					// Quiet move
					if !SideOcc[opp].IsSet(trgtSq) {
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 0, 0, 0, 0))
					} else { // Capture
						moves.AddMove(EncMove(srcSq, trgtSq, p, 0, 1, 0, 0, 0))
					}

					att = att.ClearBit(trgtSq)
				}

				bb = bb.ClearBit(srcSq)
			}
		}
	}
}

// EncMove encodes move details into a single integer.
func EncMove(src, trgt, pc, proPc, capt, dp, ep, cast int) int {
	return src | trgt<<6 | pc<<12 | proPc<<16 |
		capt<<20 | dp<<21 | ep<<22 | cast<<23
}

// DecSrc returns the source square from an encoded move.
func DecSrc(m int) int {
	return m & 0x3F
}

// DecTrgt returns the target square from an encoded move.
func DecTrgt(m int) int {
	return m & 0xFC0 >> 6
}

// DecPc returns the moved piece type from an encoded move.
func DecPc(m int) int {
	return m & 0xF000 >> 12
}

// DecProPc returns the promotion piece type from an encoded move.
func DecProPc(m int) int {
	return m & 0xF0000 >> 16
}

// DecCapt returns the piece captured flag from an encoded move.
func DecCapt(m int) int {
	return m & 0x100000 >> 20
}

// DecDp returns the double pawn push flag from an encoded move.
func DecDp(m int) int {
	return m & 0x200000 >> 21
}

// DecEp returns the en passant flag from an encoded move.
func DecEp(m int) int {
	return m & 0x400000 >> 22
}

// DecCast returns the castling flag of an encoded move.
func DecCast(m int) int {
	return m & 0x800000 >> 23
}

// MoveList contains a list of encoded moves and the count of those moves.
type MoveList struct {
	Moves []int
	Count int
}

// AddMove adds an encoded move to a move list.
func (l *MoveList) AddMove(m int) {
	l.Moves = append(l.Moves, m)
	l.Count++
}

// PrintMove prints an encoded move in the UCI format.
func PrintMove(m int) {
	fmt.Printf("%s%s%c\n", Squares[DecSrc(m)], Squares[DecTrgt(m)], PromPiece[DecProPc(m)])
}

// PrintMoveList prints all the moves in a move list.
func PrintMoveList(l *MoveList) {
	for ct := 0; ct < l.Count; ct++ {
		PrintMove(l.Moves[ct])
	}
}

// MakeMove executes an encoded move unless it leaves the king in check.
func MakeMove(m, flag int) int {
	if flag == AllMoves {
		CopyBoard()

		src := DecSrc(m)
		trgt := DecTrgt(m)
		pc := DecPc(m)
		proPc := DecProPc(m)
		capt := DecCapt(m)
		dp := DecDp(m)
		ep := DecEp(m)
		cast := DecCast(m)

		// Move piece
		PieceOcc[pc] = PieceOcc[pc].ClearBit(src)
		SideOcc[SideToMove] = SideOcc[SideToMove].ClearBit(src)
		SideOcc[Both] = SideOcc[Both].ClearBit(src)

		PieceOcc[pc] = PieceOcc[pc].SetBit(trgt)
		SideOcc[SideToMove] = SideOcc[SideToMove].SetBit(trgt)
		SideOcc[Both] = SideOcc[Both].SetBit(trgt)

		// Promotion
		if proPc != 0 {
			PieceOcc[pc] = PieceOcc[pc].ClearBit(trgt)
			PieceOcc[proPc] = PieceOcc[proPc].SetBit(trgt)
		}

		opp := SideToMove ^ 1

		// Capture
		if capt == 1 {
			var startPc, endPc int

			if SideToMove == White {
				startPc = BP
				endPc = BK
			} else { // Black
				startPc = WP
				endPc = WK
			}
			// Remove captured piece
			for p := startPc; p <= endPc; p++ {
				if PieceOcc[p].IsSet(trgt) {
					PieceOcc[p] = PieceOcc[p].ClearBit(trgt)
					SideOcc[opp] = SideOcc[opp].ClearBit(trgt)

					break
				}
			}
		}

		EnPassantSq = NA

		// Double pawn push (set en passant square)
		if dp == 1 {
			if SideToMove == White {
				EnPassantSq = trgt - 8
			} else { // Black
				EnPassantSq = trgt + 8
			}
		}
		// En passant capture
		if ep == 1 {
			if SideToMove == White {
				PieceOcc[BP] = PieceOcc[BP].ClearBit(trgt - 8)
				SideOcc[Black] = SideOcc[Black].ClearBit(trgt - 8)
				SideOcc[Both] = SideOcc[Both].ClearBit(trgt - 8)
			} else { // Black
				PieceOcc[WP] = PieceOcc[WP].ClearBit(trgt + 8)
				SideOcc[White] = SideOcc[White].ClearBit(trgt + 8)
				SideOcc[Both] = SideOcc[Both].ClearBit(trgt + 8)
			}
		}
		// Castling (move rook)
		if cast == 1 {
			switch trgt {
			case G1:
				PieceOcc[WR] = PieceOcc[WR].ClearBit(H1)
				SideOcc[White] = SideOcc[White].ClearBit(H1)
				SideOcc[Both] = SideOcc[Both].ClearBit(H1)

				PieceOcc[WR] = PieceOcc[WR].SetBit(F1)
				SideOcc[White] = SideOcc[White].SetBit(F1)
				SideOcc[Both] = SideOcc[Both].SetBit(F1)
			case C1:
				PieceOcc[WR] = PieceOcc[WR].ClearBit(A1)
				SideOcc[White] = SideOcc[White].ClearBit(A1)
				SideOcc[Both] = SideOcc[Both].ClearBit(A1)

				PieceOcc[WR] = PieceOcc[WR].SetBit(D1)
				SideOcc[White] = SideOcc[White].SetBit(D1)
				SideOcc[Both] = SideOcc[Both].SetBit(D1)
			case G8:
				PieceOcc[BR] = PieceOcc[BR].ClearBit(H8)
				SideOcc[Black] = SideOcc[Black].ClearBit(H8)
				SideOcc[Both] = SideOcc[Both].ClearBit(H8)

				PieceOcc[BR] = PieceOcc[BR].SetBit(F8)
				SideOcc[Black] = SideOcc[Black].SetBit(F8)
				SideOcc[Both] = SideOcc[Both].SetBit(F8)
			case C8:
				PieceOcc[BR] = PieceOcc[BR].ClearBit(A8)
				SideOcc[Black] = SideOcc[Black].ClearBit(A8)
				SideOcc[Both] = SideOcc[Both].ClearBit(A8)

				PieceOcc[BR] = PieceOcc[BR].SetBit(D8)
				SideOcc[Black] = SideOcc[Black].SetBit(D8)
				SideOcc[Both] = SideOcc[Both].SetBit(D8)
			}
		}

		// Update castling rights
		CastRights &= PostMoveCast[src]
		CastRights &= PostMoveCast[trgt]

		var k int

		if SideToMove == White {
			k = WK
		} else { // Black
			k = BK
		}

		// Check if move leaves the king in check
		if IsAttacked(PieceOcc[k].GetLSB(), opp) {
			TakeBack()
			return 0
		} else {
			return 1
		}
	} else if DecCapt(m) == 1 {
		MakeMove(m, AllMoves)
	}

	return 0
}
