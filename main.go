package main

import (
	"fmt"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
)

var (
	dataset = [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	target = [][]float64{
		{0},
		{1},
		{1},
		{0},
	}
)

func main() {

	w, bias := brain.NeuralNetwork([]int{2, 3, 3, 1})
	mathFuncs := []string{"tanh", "tanh", "tanh"}
	fmt.Println(bias)

	w, bias = brain.Train(0.1, mathFuncs, w, bias, dataset, target, 80)

	for i, v := range dataset {
		output, _ := brain.FeedFoward(v, mathFuncs, w, bias)
		fmt.Println(v, output, target[i])

	}
}
