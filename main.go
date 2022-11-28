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
	width, height := 9, 9

	x, y := rand.Intn(width-1), rand.Intn(height-1)
	b, v := game.CreateABoard(9, x, y, height, width)
	p := player.NewPlayer(v, []string{"soft-plus", "soft-plus", "soft-plus", "soft-plus", "soft-plus"}, []int{20, 20, 20, 20}, model, false, false, true)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		p.Brain.SaveModel(model)
		os.Exit(1)
	}()
	av := 0.0
	for i := 0; i < 10000 && !p.Won; i++ {
		m := 0.0
		for j := 0; !p.Lose && !p.Won; j++ {

			p.Evaluate(b)
			if i%100 == 0 {
				fmt.Printf("| moves: %d | average: %.4f | epoch : %d \n ", j, av/100, i)
				game.ShowBoard(p.VisibleBoard)
				time.Sleep(time.Second / 10)
			}
			m++
		}
		p.Train()
		if i%100 == 0 {
			av = 0
		}
		av += m

		x, y = rand.Intn(width-1), rand.Intn(height-1)
		b, v = game.CreateABoard(9, x, y, height, width)

		p.Clear(v)

	}
	p.Brain.SaveModel(model)

}
