package main

import (
	"fmt"

	"github.com/ranon-rat/neural-mines-sweeper/src/brain"
)

func main() {
	b := brain.OpenModel("neuralNetwork/XoR-model.json")
	lays := b.FeedFoward([]float64{1, 1})

	fmt.Println(lays[len(lays)-1])
}
