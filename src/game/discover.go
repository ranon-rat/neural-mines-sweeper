package game

import "github.com/ranon-rat/neural-mines-sweeper/src/core"

func discover(y, x int, discoveredPosition core.UniquePosition, visibleBoard, board [][]int) (discoveredCells core.UniquePosition, finalBoard [][]int) {
	finalBoard = visibleBoard
	discoveredCells = discoveredPosition
	discoveredCells.Add(core.XY{X: x, Y: y})
	finalBoard[y][x] = board[y][x]
	if board[y][x] > 0 {
		return
	}
	/*
		up
		down
		left
		right

	*/
	discoveredCells, finalBoard = discoverLeftAndRight(y, x, discoveredCells, finalBoard, board)

	if y != 0 {
		if !discoveredPosition[core.XY{X: x, Y: y - 1}] {

			discoveredCells, finalBoard = discover(y-1, x, discoveredCells, finalBoard, board)
		}

		discoveredCells, finalBoard = discoverLeftAndRight(y-1, x, discoveredCells, finalBoard, board)
	}
	if y != len(board)-1 {
		if !discoveredPosition[core.XY{X: x, Y: y + 1}] {

			discoveredCells, finalBoard = discover(y+1, x, discoveredCells, finalBoard, board)
		}
		discoveredCells, finalBoard = discoverLeftAndRight(y+1, x, discoveredCells, finalBoard, board)
	}
	return
}
func discoverLeftAndRight(y, x int, discoveredPosition core.UniquePosition, visibleBoard, board [][]int) (discoveredCells core.UniquePosition, finalBoard [][]int) {
	finalBoard = visibleBoard
	finalBoard[y][x] = board[y][x]
	discoveredCells = discoveredPosition
	discoveredCells.Add(core.XY{X: x, Y: y})
	if board[y][x] > 0 {
		return
	}
	if x != 0 {
		if !discoveredPosition[core.XY{X: x - 1, Y: y}] {
			discoveredCells, finalBoard = discover(y, x-1, discoveredCells, finalBoard, board)
		}
	}
	if x != len(board[y])-1 {
		if !discoveredPosition[core.XY{X: x + 1, Y: y}] {
			discoveredCells, finalBoard = discover(y, x+1, discoveredCells, finalBoard, board)
		}

	}
	return
}
