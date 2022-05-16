package move

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"

)

type Move struct {
	Name uint8
	Piece uint64
	Start uint64
	End uint64
}

func NewMove(name uint8, piece uint64, start uint64, end uint64) *Move {
	return &Move{name, piece, start, end}
}

func GetAllMoves(st *state.State) []*Move {
	if st.White == 1 {
		return GetAllMovesWhite(st)
	} else {
		return GetAllMovesBlack(st)
	}
}

func GetAllMovesWhite(st *state.State) []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white)
		for _, m := range moves {
			res = append(res, NewMove(0, st.WK, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, white, black)
		for _, m := range moves {
			res = append(res, NewMove(1, st.WQ, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.WR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, white, black)
		for _, m := range moves {
			res = append(res, NewMove(2, st.WR, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.WB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, white, black)
		for _, m := range moves {
			res = append(res, NewMove(3, st.WB, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.WN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, white)
		for _, m := range moves {
			res = append(res, NewMove(4, st.WN, n, m))
		}
	}
	//Pawns
	pawns := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, white, black)
		for _, m := range moves {
			res = append(res, NewMove(5, st.WP, p, m))
		}
	}

	return res
}

func GetAllMovesBlack(st *state.State) []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.BK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, black)
		for _, m := range moves {
			res = append(res, NewMove(0, st.BK, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.BQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, black, white)
		for _, m := range moves {
			res = append(res, NewMove(1, st.BQ, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.BR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, black, white)
		for _, m := range moves {
			res = append(res, NewMove(2, st.BR, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.BB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, black, white)
		for _, m := range moves {
			res = append(res, NewMove(3, st.BB, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.BN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, black)
		for _, m := range moves {
			res = append(res, NewMove(4, st.BN, n, m))
		}
	}
	//Pawns
	pawns := pieces.GetPositionsFromBoard(st.BP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, black, white)
		for _, m := range moves {
			res = append(res, NewMove(5, st.BP, p, m))
		}
	}

	return res
}