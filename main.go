package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
	"github.com/ranon-rat/neural-mines-sweeper/src/player"
)

// this model is trained
func main() {

	width, height := 10, 10

	x, y := rand.Intn(width-1), rand.Intn(height-1)
	b, v := game.CreateABoard(x, y, height, width, 0.2)
	p := player.NewPlayer(v, nil, nil, "neuralNetwork/mines-sweeper.json", false)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		p.Brain.SaveModel("neuralNetwork/mines-sweeper.json")
		os.Exit(1)
	}()
	for i := 0; !p.Won && i < 10000; i++ {
		for j := 0; !p.Lose && !p.Won; j++ {
			p.Evaluate(b)
			if i%40 == 0 {
				fmt.Printf("| Generation: %d | move: %d |\n", i, j)
				game.ShowBoard(p.VisibleBoard)
				time.Sleep(time.Second / 3)

			}

		}
		x, y = rand.Intn(width), rand.Intn(height)
		ex, logs := p.Train(b), p.LogsInput

		core.CreateData("data/minessweeper.csv", logs, ex, true, 9)

		b, v = game.CreateABoard(x, y, height, width, 0.2)
		p.Clear(v)

	}
	p.Brain.SaveModel("neuralNetwork/mines-sweeper.json")
}
