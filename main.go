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
		{1, 0},
		{0, 1},
		{0, 1},
		{1, 0},
	}
)

func main() {

	w, bias := brain.NeuralNetwork([]int{2, 3, 3, 2})
	mathFuncs := []string{"sigmoid", "sigmoid", "sigmoid"}
	fmt.Println(bias)
	w, bias = brain.Train(0.1, mathFuncs, w, bias, dataset, target, 1000)

	for i, v := range dataset {
		output, _ := brain.FeedFoward(v, mathFuncs, w, bias)
		fmt.Println("is x bigger than x", output, target[i])

	}

}
