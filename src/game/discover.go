package game

import "github.com/ranon-rat/neural-mines-sweeper/src/core"

func discover(y, x int, discoveredPos core.UniquePosition, visibleBoard, board [][]int) (discCells core.UniquePosition, finalBoard [][]int) {
	finalBoard = visibleBoard
	discCells = discoveredPos
	discCells.Add(core.XY{X: x, Y: y})
	finalBoard[y][x] = board[y][x]
	if board[y][x] > 0 {
		return
	}

	discCells, finalBoard = discoverLAndR(y, x, discCells, finalBoard, board)

	if y != 0 {
		if !discoveredPos[core.XY{X: x, Y: y - 1}] {

			discCells, finalBoard = discover(y-1, x, discCells, finalBoard, board)
		}

		discCells, finalBoard = discoverLAndR(y-1, x, discCells, finalBoard, board)
	}
	if y != len(board)-1 {
		if !discoveredPos[core.XY{X: x, Y: y + 1}] {

			discCells, finalBoard = discover(y+1, x, discCells, finalBoard, board)
		}
		discCells, finalBoard = discoverLAndR(y+1, x, discCells, finalBoard, board)
	}
	return
}
func discoverLAndR(y, x int, discoveredPos core.UniquePosition, visibleBoard, board [][]int) (discCells core.UniquePosition, finalBoard [][]int) {
	finalBoard = visibleBoard
	finalBoard[y][x] = board[y][x]
	discCells = discoveredPos
	discCells.Add(core.XY{X: x, Y: y})
	if board[y][x] > 0 {
		return
	}
	if x != 0 {
		if !discoveredPos[core.XY{X: x - 1, Y: y}] {
			discCells, finalBoard = discover(y, x-1, discCells, finalBoard, board)
		}
	}
	if x != len(board[y])-1 {
		if !discoveredPos[core.XY{X: x + 1, Y: y}] {
			discCells, finalBoard = discover(y, x+1, discCells, finalBoard, board)
		}

	}
	return
}
