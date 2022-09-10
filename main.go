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

	w, bias := brain.NeuralNetwork([]int{2, 3, 9, 1})
	mathFuncs := []string{"sigmoid", "sigmoid", "sigmoid"}
	fmt.Println(bias)

	w, bias = brain.Train(0.1, mathFuncs, w, bias, dataset, target, 800)

}
