package player

import (
	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

type Player struct {
	Brain                  brain.NN
	VisibleBoard           [][]int `json:"-"`
	Lose, Won, SupLearning bool

	LogsInput [][]float32
	LogsMoves []core.XY
}

//thinking in multiple ways of evaluating this
//hm maybe the way that i should do this is idk
func NewPlayer(visibleBoard [][]int, activationFuncs []string, hiddenLayersLength []int, modelFile string, supLearning bool) (p Player) {
	p.SupLearning = supLearning
	if modelFile == "" {
		hiddenLayersLength = append(hiddenLayersLength, 1)
		layers := append([]int{9}, hiddenLayersLength...)
		p.Brain = brain.NewNeuralNetwork(layers, activationFuncs, "a simple model lol")
	} else {
		p.Brain = brain.OpenModel(modelFile)
	}
	p.VisibleBoard = visibleBoard

	//0 open
	//1 dontopen

	return
}

func (p *Player) Evaluate(board [][]int) {
	xyAvaible := FindAvaibleCells(p.VisibleBoard)
	calfAndPos := []CalfAndPos{}
	for _, v := range xyAvaible {
		// first I get the input from the board
		input := GetInput(p.VisibleBoard, v.Y, v.X, 9)
		// then i pass it to the neural network
		out := p.Brain.Predict(input)
		// the index 0 is for opening the cell
		calfAndPos = append(calfAndPos, CalfAndPos{Calf: out[0], Pos: v})
	}

	bigIndx := GetBestPos(calfAndPos)
	pos := calfAndPos[bigIndx].Pos
	p.LogsInput = append(p.LogsInput, GetInput(p.VisibleBoard, pos.Y, pos.X, 9))
	p.LogsMoves = append(p.LogsMoves, pos)

	p.VisibleBoard, p.Lose, p.Won = game.MakeAMove(pos.Y, pos.X, p.VisibleBoard, board)

}
func (p *Player) Train(board [][]int) [][]float32 {
	expected := [][]float32{}
	for _, v := range p.LogsMoves {
		expected = append(expected, []float32{map[bool]float32{true: 0, false: 1}[board[v.Y][v.X] == 9]})

	}
	p.Brain.Train(p.LogsInput, expected, 0.25, 50, false)
	return expected

}

func (p *Player) Clear(visibleBoard [][]int) {
	p.LogsMoves = []core.XY{}
	p.LogsInput = [][]float32{}
	p.VisibleBoard = visibleBoard
	p.Lose = false
}
