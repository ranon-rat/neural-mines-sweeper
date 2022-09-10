package main

import (
	"fmt"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
)

// the training works
func main() {
	dataset, expected := core.LoadData("../../data/minessweeper.csv", 1, 9, 1)
	b := brain.NewNeuralNetwork([]int{9, 16, 16, 1}, []string{"sigmoid", "sigmoid", "sigmoid"}, "just a simple model  ")
	b.Train(dataset, expected, 0.5, 500, true)
	for i, v := range dataset {
		fmt.Printf("input: %v | expected: %v | output :%v\n", v, expected[i], b.Predict(v))
	}
	b.SaveModel("../../neuralNetwork/mines-sweeper.json")
}
