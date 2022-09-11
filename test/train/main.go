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
	dataset, expected := core.LoadData("../../data/minessweeper.csv", 1, 9, 1)
	b := brain.OpenModel("../../neuralNetwork/mines-sweeper.json")
	b.Train(dataset, expected, 0.272, 300, true)
	for j := 0; j < 10; j++ {
		i := rand.Intn(len(dataset))
		v := dataset[i]
		fmt.Printf("expected: %v | output :%v\n", expected[i], b.Predict(v))
		for j, k := range v {
			if j%5 == 0 {
				fmt.Println()
			}
			s := strconv.Itoa(int(k*9)) + " "
			if _, e := game.Characters[int(k*10)]; e {
				s = game.Characters[int(k*9)]
			}
			fmt.Printf("%s", s)
		}
		fmt.Println()
	}
	b.SaveModel("../../neuralNetwork/mines-sweeper.json")
}
