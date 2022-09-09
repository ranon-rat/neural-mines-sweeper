package brain

import (
	"fmt"
	"math/rand"
	"time"
)

func NeuralNetwork(neuronsPerLayer []int) (weights [][][]float64, bias [][]float64) {
	rand.Seed(time.Now().Unix())
	weights = make([][][]float64, len(neuronsPerLayer)-1)
	bias = make([][]float64, len(neuronsPerLayer)-1)
	for i := 0; i < len(neuronsPerLayer)-1; i++ {

		for n := 0; n < neuronsPerLayer[i]; n++ {
			// i dont need to know the bias for the input but whatever

			weights[i] = append(weights[i], []float64{})
			for w := 0; w < neuronsPerLayer[i+1]; w++ {
				weights[i][n] = append(weights[i][n], rand.Float64()/2)
			}

		}
		for n := 0; n < neuronsPerLayer[i+1]; n++ {
			// i dont need to know the bias for the input but whatever
			bias[i] = append(bias[i], rand.Float64()/2)
		}
	}

	return
}

func FeedFoward(input []float64, mathFuncPerLayer []string, weights [][][]float64, bias [][]float64) (output []float64, layers [][]float64) {
	layers = make([][]float64, len(bias)+1)

	layers[0] = make([]float64, len(input))
	copy(layers[0], input)
	for l := 0; l < len(layers)-1; l++ {
		layers[l+1] = make([]float64, len(bias[l]))

		for n := 0; n < len(layers[l]); n++ {

			for i, w := range weights[l][n] {
				layers[l+1][i] += w * layers[l][n]

			}

		}
		for i, n := range layers[l+1] {
			layers[l+1][i] = MathFuncs[mathFuncPerLayer[l]]["activate"](n + bias[l][i])

		}

	}
	output = layers[len(layers)-1]
	return
}

// what im doing with my life :weary:
// how the fuck i did this ?
func BackPropagation(size float64, weights [][][]float64, bias, layers [][]float64, expected []float64, mathFuncPerLayer []string) ([][][]float64, [][]float64) {
	bd := make([][]float64, len(bias))
	wd := make([][][]float64, len(weights))
	errors := make([]float64, len(expected))
	layer := layers[len(layers)-1]
	for i, n := range layer {
		errors[i] = n - expected[i]

	}
	for l := len(layers) - 2; l >= 0; l-- {
		bd[l] = make([]float64, len(bias[l]))
		wd[l] = make([][]float64, len(weights[l]))

		for i := range bias[l] {

			bd[l][i] = (errors[i] * MathFuncs[mathFuncPerLayer[l]]["derivative"](layer[i]))
		}
		for n := 0; n < len(weights[l]); n++ {
			wd[l][n] = make([]float64, len(weights[l][n]))

			for i := range weights[l][n] {

				wd[l][n][i] = ((bd[l][i]) * layers[l][n]) / size
				bd[l][i] /= size
			}
		}

		if l == 0 {
			continue
		}

		layer = layers[l]
		errorcp := make([]float64, len(layer))
		for i := range layer {
			err := 0.0
			for j, v := range errors {
				err += weights[l][i][j] * v
			}
			errorcp[i] = err
		}
		errors = errorcp

	}

	return wd, bd
}
func UpdateWeightAndBias(learningRate float64, weights [][][]float64, bias [][]float64, weightsGrad [][][]float64, biasGrad [][]float64) ([][][]float64, [][]float64) {

	for l := 0; l < len(weights)-1; l++ {

		for n := 0; n < len(weights[l]); n++ {

			for i := range weights[l][n] {

				weights[l][n][i] -= (weightsGrad[l][n][i]) * learningRate
			}

		}
		for i := range bias[l] {
			bias[l][i] -= (biasGrad[l][i]) * learningRate
		}

	}

	return weights, bias
}
func Train(learningRate float64, mathFuncs []string, weights [][][]float64, bias, dataset [][]float64, expected [][]float64, epochs int) ([][][]float64, [][]float64) {
	for i := 0; i < epochs; i++ {

		err := 0.0
		for j, v := range dataset {
			out, layers := FeedFoward(v, mathFuncs, weights, bias)

			wd, bd := BackPropagation(float64(len(dataset)), weights, bias, layers, expected[j], mathFuncs)
			weights, bias = UpdateWeightAndBias(learningRate, weights, bias, wd, bd)

			err += Cost(expected[j], out)

		}
		if i%10 == 0 {
			fmt.Println(err)
		}
	}
	return weights, bias
}
