package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
	"github.com/ranon-rat/neural-mines-sweeper/src/player"
)

// this model is trained
func main() {

	width, height := 10, 10
	x, y := rand.Intn(width), rand.Intn(height)

	b, v := game.CreateABoard(x, y, height, width, 0.2)
	p := player.NewPlayer(v, nil, nil, "neuralNetwork/mines-sweeper.json", false)

	for i := 0; !p.Won && i < 10000; i++ {
		for j := 0; !p.Lose && !p.Won; j++ {
			p.Evaluate(b)
			if i%40 == 0 {
				fmt.Printf("| Generation: %d | move: %d |\n", i, j)
				game.ShowBoard(p.VisibleBoard)
				time.Sleep(time.Second / 2)

			}

		}
		x, y = 1+rand.Intn(width-1), 1+rand.Intn(height-1)

		core.CreateData("data/minessweeper.csv", p.LogsInput, p.Train(b), true, 9)

		b, v = game.CreateABoard(x, y, height, width, 0.2)
		p.Clear(v)

	}
	p.Brain.SaveModel("neuralNetwork/mines-sweeper.json")
}
