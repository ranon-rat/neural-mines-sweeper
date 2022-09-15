package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/ranon-rat/neural-mines-sweeper/src/game"
	"github.com/ranon-rat/neural-mines-sweeper/src/player"
)

var (
	model = "neuralNetwork/mines-sweeper-self-trained.json"
)

// this model is trained
func main() {
	width, height := 15, 10

	x, y := rand.Intn(width-1), rand.Intn(height-1)
	b, v := game.CreateABoard(x, y, height, width, 0.1)
	p := player.NewPlayer(v, nil, nil, model, false)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		p.Brain.SaveModel(model)
		os.Exit(1)
	}()
	for i := 0; !p.Won && i < 10000; i++ {

		for j := 0; !p.Lose && !p.Won; j++ {
			p.Evaluate(b)

			if i%100 == 0 {
				fmt.Printf("| moves: %d | epoch : %d \n", j, i)
				game.ShowBoard(p.VisibleBoard)
				time.Sleep(time.Second / 5)
			}

		}
		x, y = rand.Intn(width), rand.Intn(height)
		b, v = game.CreateABoard(x, y, height, width, 0.1)
		//p.Train(b)

		//if len(p.LogsInput) > 3 || p.Won {
		//
		//	core.CreateData("data/minessweeper.csv", logs, ex, true, game.Bomb-1)
		//
		//}

		p.Clear(v)

	}
	p.Brain.SaveModel(model)

}
