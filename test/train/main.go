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
	dataset, expected := core.LoadData("../../data/minessweeper-dev.csv", 1, 1, 1)

	//b :=// brain.NewNeuralNetwork([]int{25, 64, 64, 30, 2}, []string{"soft-plus", "soft-plus", "soft-plus", "soft-plus"}, "a simple model lol")
	b := brain.OpenModel("./test.json")
	//b.Train(dataset, expected, 0.01, 100, true)
	for j := 0; j < 5; j++ {
		i := rand.Intn(len(dataset))
		v := dataset[i]
		fmt.Printf("expected: %v | output :%v\n", core.GetBiggerIndex(expected[i]), core.GetBiggerIndex(b.Predict(v)))
		for j, k := range v {
			if j%5 == 0 {
				fmt.Println()
			}
			s := strconv.Itoa(int(k*(game.Bomb-1))) + " "
			if _, e := game.Characters[int(k)]; e {
				s = game.Characters[int(k)]
			}
			fmt.Printf("%s", s)
		}
		fmt.Println()
	}
	b.SaveModel("test.json")
}
