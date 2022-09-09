package main

import (
	"fmt"
	"math/rand"

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

	for i := 0; i < 30; i++ {
		pos := rand.Intn(len(dataset) - 1)

		v := dataset[pos]

		w, bias = brain.Train(0.5, mathFuncs, w, bias, [][]float64{dataset[pos]}, [][]float64{target[pos]}, 1)

		output, _ := brain.FeedFoward(v, mathFuncs, w, bias)
		fmt.Println(v, output, target[pos])

	}

}
