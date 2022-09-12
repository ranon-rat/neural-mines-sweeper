package player

import (
	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

type Player struct {
	Brain                  brain.NN
	VisibleBoard           [][]int
	Lose, Won, SupLearning bool

	LogsInput [][]float32
	LogsMoves []core.XY
}

//thinking in multiple ways of evaluating this
//hm maybe the way that i should do this is idk
func NewPlayer(visibleBoard [][]int, activationFuncs []string, hiddenLayersLength []int, modelFile string, supLearning bool) (p Player) {
	p.SupLearning = supLearning
	if modelFile == "" {
		hiddenLayersLength = append(hiddenLayersLength, 2)
		layers := append([]int{25}, hiddenLayersLength...)
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
	itsFine := []CalfAndPos{}
	for _, v := range xyAvaible {
		// first I get the input from the board
		input := GetInput(p.VisibleBoard, v.Y, v.X, game.Bomb-1)
		// then i pass it to the neural network
		out := p.Brain.Predict(input)
		// the index 0 is for opening the cell
		calfAndPos = append(calfAndPos, CalfAndPos{Calf: out[1], Pos: v})
		// 1 open obviously
		if core.GetBiggerIndex(out) == 1 {
			itsFine = append(calfAndPos, CalfAndPos{Calf: out[1], Pos: v})
		}

	}
	// i just see if the list "itsfine"  is empty , only for not having some weird errors
	if len(itsFine) > 0 {
		calfAndPos = itsFine
	}

	bigIndx := GetBestPos(calfAndPos)
	pos := calfAndPos[bigIndx].Pos
	if !p.SupLearning {
		p.LogsInput = append(p.LogsInput, GetInput(p.VisibleBoard, pos.Y, pos.X, game.Bomb-1))
		p.LogsMoves = append(p.LogsMoves, pos)
	}
	p.VisibleBoard, p.Lose, p.Won = game.MakeAMove(pos.Y, pos.X, p.VisibleBoard, board)

}

// the board its for knowing the predictions, IM NOT GOING TO FUCKING SAVE THE FUCKING LAYERS
func (p *Player) Train(board [][]int) [][]float32 {
	expected := [][]float32{}
	for _, v := range p.LogsMoves {
		// because I need to have a double output i need to do this

		expected = append(expected, map[bool][]float32{true: {1, 0}, false: {0, 1}}[board[v.Y][v.X] == game.Bomb])

	}
	// for some reason that i detected i need to use a really low learning rate for the training process
	// maybe you can search for that
	p.Brain.Train(p.LogsInput, expected, 0.000002, 40, false)
	return expected

}

func (p *Player) Clear(visibleBoard [][]int) {
	p.LogsMoves = []core.XY{}
	p.LogsInput = [][]float32{}
	p.VisibleBoard = visibleBoard
	p.Lose = false
}
