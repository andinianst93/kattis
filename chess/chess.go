package chess

import "fmt"

const (
	correctKings = 1
	correctQueens = 1
	correctRooks = 2
	correctBishops = 2
	correctKnights = 2
	correctPawns = 8
)

func Chess() {
	var kings, queens, rooks, bishops, knights, pawns int
	fmt.Scan(&kings, &queens, &rooks, &bishops, &knights, &pawns)

	neededKings := correctKings - kings
	neededQueens := correctQueens - queens
	neededRooks := correctRooks - rooks
	neededBishops := correctBishops - bishops
	neededKnights := correctKnights - knights
	neededPawns := correctPawns - pawns

	fmt.Printf("%d %d %d %d %d %d\n", neededKings, neededQueens, neededRooks, neededBishops, neededKnights, neededPawns)
}
