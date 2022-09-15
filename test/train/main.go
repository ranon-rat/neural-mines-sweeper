package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

// the training works
func main() {
	dataset, expected := core.LoadData("../../data/minessweeper.csv", 1, game.Bomb-1, 1)
	model := "../../neuralNetwork/mines-sweeper-self-trained.json"

	b := brain.OpenModel(model)
	b.Train(dataset, expected, 0.01, 100, true)
	for j := 0; j < 5; j++ {
		i := rand.Intn(len(dataset))
		v := dataset[i]
		fmt.Printf("expected: %v | output :%v\n", core.GetBiggerIndex(expected[i]), core.GetBiggerIndex(b.Predict(v)))
		for j, k := range v {
			if j%5 == 0 {
				fmt.Println()
			}
			s := strconv.Itoa(int(k*(game.Bomb-1))) + " "
			if _, e := game.Characters[int(k*(game.Bomb-1))]; e {
				s = game.Characters[int(k*(game.Bomb-1))]
			}
			fmt.Printf("%s", s)
		}
		fmt.Println()
	}
	b.SaveModel(model)
}
