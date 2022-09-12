package game

import "github.com/ranon-rat/neural-mines-sweeper/src/core"

func MakeAMove(y, x int, visibleBoard, board [][]int) (finalBoard [][]int, lose, wins bool) {
	cell := board[y][x]
	lose = cell == Bomb

	visibleBoard[y][x] = cell
	if !lose {
		_, visibleBoard = discover(y, x, core.UniquePosition{}, visibleBoard, board)
		wins = ThePlayerWins(visibleBoard, board)
	}
	finalBoard = visibleBoard
	return
}

// fuck , i cant compare a 2 dimensional array in go
// i only can do it with nil
//hm
func ThePlayerWins(visibleBoard, board [][]int) bool {
	//I fucking hate my existence
	cells := 0
	for y := 0; y < len(visibleBoard); y++ {
		for x := 0; x < len(visibleBoard[y]); x++ {
			if board[y][x] == Bomb && visibleBoard[y][x] != board[y][x] {
				cells++
				continue

			}
			if visibleBoard[y][x] == board[y][x] {
				cells++
				// this shit i so fucking shitty
			}
		}
	}
	return cells == (len(visibleBoard))*(len(visibleBoard[0]))
}
