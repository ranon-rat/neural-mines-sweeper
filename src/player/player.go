package player

import (
	"sync"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

type Player struct {
	Weights           [][][]float64 `json:"weights"`
	Biases            [][]float64   `json:"biases"`
	MathFuncsPerLayer []string      `json:"functions"`
	VisibleBoard      [][]int       `json:"-"`
	Lose, Won         bool
}

//thinking in multiple ways of evaluating this
//hm maybe the way that i should do this is idk
func NewPlayer(visibleBoard [][]int) (p Player) {
	p.Weights, p.Biases = brain.NeuralNetwork([]int{
		9, 32, 16, 8, 2,
	})
	p.MathFuncsPerLayer = []string{"sigmoid", "sigmoid", "sigmoid", "sigmoid"}
	p.VisibleBoard = make([][]int, len(visibleBoard))
	for i, v := range visibleBoard {
		p.VisibleBoard[i] = make([]int, len(visibleBoard[0]))
		for j, k := range v {

			p.VisibleBoard[i][j] = k
		}

	}
	//0 open
	//1 dontopen

	return
}

func (p *Player) Evaluate(yMv, xMv int, board [][]int) {
	var wg sync.WaitGroup
	xyAvaible := FindAvaibleCells(xMv, yMv, p.VisibleBoard)
	calfAndPos := []CalfAndPos{}
	for _, v := range xyAvaible {
		wg.Add(1)

		go func(v core.XY) {
			// first I get the input from the board
			input := getInput(p.VisibleBoard, v.Y, v.X)
			// then i pass it to the neural network
			output, _ := brain.FeedFoward(input, p.MathFuncsPerLayer, p.Weights, p.Biases)
			// the index 0 is for opening the cell
			calfAndPos = append(calfAndPos, CalfAndPos{Calf: output[0], Pos: v})

			wg.Done()
		}(v)
	}
	wg.Wait()
	bigIndx := GetBestPos(calfAndPos)
	pos := calfAndPos[bigIndx].Pos
	p.VisibleBoard, p.Lose, p.Won = game.MakeAMove(pos.Y, pos.X, p.VisibleBoard, board)

}
