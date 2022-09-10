package main

import (
	"fmt"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
)

// the training works
func main() {
	dataset, expected := core.LoadData("data/xor.csv", 1, 1, 1)
	b := brain.NewNeuralNetwork([]int{2, 3, 1}, []string{"tanh", "tanh"}, "just a xor model ")
	b.Train(dataset, expected, 1, 400, true)
	for i, v := range dataset {
		fmt.Printf("input: %v | expected: %v | output :%v\n", v, expected[i], b.Predict(v))
	}
}
