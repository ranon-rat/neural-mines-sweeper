package game

import (
	"math/rand"
	"time"

	"github.com/ranon-rat/neural-mines-sweeper/src/core"
)

var bin = map[bool]int{false: 1, true: 0}

func CreateABoard(xMv, yMv, height, width int) (board, visibleBoard [][]int) {
	rand.Seed(time.Now().Unix())
	for y := 0; y < height; y++ {
		visibleBoard = append(visibleBoard, []int{})
		board = append(board, []int{})
		for x := 0; x < width; x++ {

			board[y] = append(board[y], 0)
			visibleBoard[y] = append(visibleBoard[y], -1)
		}
	}
	for y := 0; y < len(board); y++ {

		for x := 0; x < len(board[y]); x++ {
			/*
				###
				#8#
				###
			*/

			if rand.Float64() < 0.2 && !checkLR(y, x, xMv, yMv) && !checkLR(y+1, x, xMv, yMv) && !checkLR(y-1, x, xMv, yMv) {
				board[y][x] = bomb

				//					winnerBoard[y][x] = undiscoveredCell

				board = addInLeftAndRight(y, x, width, board, false)

				if y > 0 {
					//board[y-1][x] += bin[board[y-1][x] == bomb]
					//winnerBoard[y-1][x] = board[y-1][x]
					//
					board = addInLeftAndRight(y-1, x, width, board, false)

				}
				if y != height-1 {
					//board[y+1][x] += bin[board[y+1][x] == bomb]
					//winnerBoard[y+1][x] = board[y+1][x]
					board = addInLeftAndRight(y+1, x, width, board, false)
				}

			}

		}
	}
	visibleBoard[yMv][xMv] = board[yMv][xMv]
	_, visibleBoard = discover(yMv, xMv, core.UniquePosition{}, visibleBoard, board)
	return
}
func checkLR(y, x, xMv, yMv int) bool {
	return ((x == xMv) || (x == xMv+1) || (xMv-1 == x)) && (y == yMv)
}
func addInLeftAndRight(y, x, width int, board [][]int, thereIsABomb bool) [][]int {

	board[y][x] += bin[board[y][x] == bomb]

	if x > 0 {
		board[y][x-1] += bin[board[y][x-1] == bomb]
	}
	if x != width-1 {
		board[y][x+1] += bin[board[y][x+1] == bomb]
	}
	return board
}