package main

import (
	"math/rand"

	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
	"github.com/ranon-rat/neural-mines-sweeper/src/player"
)

// this model is trained
func main() {
	input := [][]float32{}
	output := [][]float32{}
	for i := 0; i < 10000; i++ {
		x, y := rand.Intn(9), rand.Intn(9)

		board, visibleBoard := game.CreateABoard(x, y, 10, 10, 0.2)

		pos := player.FindAvaibleCells(visibleBoard)
		for _, v := range pos {
			part := player.GetInput(visibleBoard, v.Y, v.X, 1)
			input = append(input, part)

			output = append(output, map[bool][]float32{true: {0}, false: {1}}[board[v.Y][v.X] == 9])

		}
	}
	core.CreateData("../../data/minessweeper.csv", input, output, true)
}
