package main

import (
	"fmt"
	"math/rand"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
)

func main() {
	dataset := [][]float64{}
	target := [][]float64{}
	for len(dataset) < 30 {
		x, y := rand.Float64(), rand.Float64()
		dataset = append(dataset, []float64{x, y})
		target = append(target, []float64{map[bool]float64{true: 1, false: 0}[(x > y)]})

	}
	w, bias := brain.NeuralNetwork([]int{2, 3, 1})
	mathFuncs := []string{"tanh", "tanh"}
	fmt.Println(w, "\n", bias, "")

	w, bias = brain.Train(0.5, mathFuncs, w, bias, dataset, target, 80)
	for i := 0; i < 30; i++ {
		x, y := rand.Float64(), rand.Float64()
		output, _ := brain.FeedFoward(mathFuncs, []float64{x, y}, w, bias)
		fmt.Println("is x bigger than x", output[0], x > y)

	}

}
