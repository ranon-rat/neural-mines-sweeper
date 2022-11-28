package player

import (
	"math/rand"
	"time"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

type Player struct {
	Brain                                                          brain.NN
	VisibleBoard                                                   [][]int
	Lose, Won, saveData, learnEachIteration, reinforcementLearning bool

	LogsInput, LogsExpected, logsOutput [][]float32
}

// thinking in multiple ways of evaluating this
// hm maybe the way that i should do this is idk
func NewPlayer(visibleBoard [][]int, activationFuncs []string, hiddenLayersLength []int, modelFile string, save, learnEach, reinforcementLearning bool) (p Player) {
	p.saveData = save
	p.learnEachIteration = learnEach
	p.reinforcementLearning = reinforcementLearning
	if modelFile == "" {
		hiddenLayersLength = append(hiddenLayersLength, 1)
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
		input := GetInput(p.VisibleBoard, v.Y, v.X, 1)
		// then i pass it to the neural network
		out := p.Brain.Predict(input)
		// the index 0 is for opening the cell
		calfAndPos = append(calfAndPos, CalfAndPos{Calf: out, Pos: v})
		// 1 open obviously
		if core.GetBiggerIndex(out) == 1 {
			itsFine = append(itsFine, CalfAndPos{Calf: out, Pos: v})

		}

		if p.saveData {
			p.LogsInput = append(p.LogsInput, input)
			p.LogsExpected = append(p.LogsExpected, map[bool][]float32{true: {1}, false: {0}}[board[v.Y][v.X] == game.Bomb])
		}
		if p.learnEachIteration {

			p.Brain.Train([][]float32{input},
				[][]float32{map[bool][]float32{true: {1}, false: {0}}[board[v.Y][v.X] == game.Bomb]},
				0.125, 1, false)
		}

	}
	// i just see if the list "itsfine"  is empty , only for not having some weird errors
	if len(itsFine) > 0 {
		calfAndPos = itsFine
	}

	bigIndx := GetBestPos(calfAndPos)
	pos := calfAndPos[bigIndx].Pos

	p.VisibleBoard, p.Lose, p.Won = game.MakeAMove(pos.Y, pos.X, p.VisibleBoard, board)
	if p.reinforcementLearning {
		p.LogsInput = append(p.LogsInput, GetInput(p.VisibleBoard, calfAndPos[bigIndx].Pos.Y, calfAndPos[bigIndx].Pos.X, 1))
		p.logsOutput = append(p.logsOutput, calfAndPos[bigIndx].Calf)
	}

}

// the board its for knowing the predictions, IM NOT GOING TO FUCKING SAVE THE FUCKING LAYERS
func (p *Player) Train() {

	// for some reason that i detected i need to use a really low learning rate for the training process
	// maybe you can search for that
	wdSum, bdSum := [][][][]float32{}, [][][]float32{}
	for i, v := range p.LogsInput {
		layers := p.Brain.FeedFoward(v)
		wd, bd := p.Brain.BackPropagation(layers, p.logsOutput[i])
		wdSum = append(wdSum, wd)
		bdSum = append(bdSum, bd)

	}
	for i, v := range wdSum {
		p.Brain.UpdateWeightAndBias(float32(len(p.logsOutput)), 0.01, v, bdSum[i])
	}
}

func (p *Player) Clear(visibleBoard [][]int) {
	if len(p.LogsExpected) > 1000 {

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(p.LogsInput), func(i, j int) {
			p.LogsInput[i], p.LogsInput[j] = p.LogsInput[j], p.LogsInput[i]
			p.logsOutput[i], p.logsOutput[j] = p.logsOutput[j], p.logsOutput[i]
			p.LogsExpected[i], p.LogsExpected[j] = p.LogsExpected[j], p.LogsExpected[i]

		})
		p.logsOutput = p.logsOutput[:1000]
		p.LogsInput = p.logsOutput[:1000]

		p.LogsExpected = p.logsOutput[:1000]

	}
	p.VisibleBoard = visibleBoard
	p.Lose = false
}
