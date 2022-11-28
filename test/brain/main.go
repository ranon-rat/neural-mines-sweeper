package main

import (
	"fmt"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
)

var dataset = [][]float32{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 1}}
var expected = [][]float32{
	{0},
	{1},
	{1},
	{0},
}

// the training works
func main() {
	b := brain.NewNeuralNetwork([]int{2, 3, 1}, []string{"soft-plus", "soft-plus"}, "just a xor model ")

	b.Train(dataset, expected, 2, 400, true)
	for i, v := range dataset {
		fmt.Printf("input: %v | expected: %v | output :%v\n", v, expected[i], b.Predict(v))
	}
}
