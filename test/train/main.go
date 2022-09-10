package main

import (
	"fmt"
	"strconv"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

// the training works
func main() {
	dataset, expected := core.LoadData("../../data/minessweeper.csv", 1, 9, 1)
	b := brain.OpenModel("../../neuralNetwork/mines-sweeper.json") //brain.NewNeuralNetwork([]int{9, 36, 32, 1}, []string{"tanh", "tanh", "tanh", "tanh"}, "mines sweeper model , isnt that perfect but it kinda works")
	b.Train(dataset, expected, 0.2, 500, true)
	for i, v := range dataset[:10] {
		fmt.Printf("expected: %v | output :%v\n", expected[i], b.Predict(v))
		for j, k := range v {
			if j%3 == 0 {
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
