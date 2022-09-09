package brain

import (
	"math/rand"
	"sync"
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
				weights[i][n] = append(weights[i][n], rand.Float64())
			}

		}
		for n := 0; n < neuronsPerLayer[i+1]; n++ {
			// i dont need to know the bias for the input but whatever
			bias[i] = append(bias[i], rand.Float64())
		}
	}

	return
}

func FeedFoward(mathFuncPerLayer []string, input []float64, weights [][][]float64, bias [][]float64) (output []float64, layers [][]float64) {
	layers = make([][]float64, len(bias)+1)

	layers[0] = input
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
func BackPropagation(weights [][][]float64, bias, layers [][]float64, expected []float64, mathFuncPerLayer []string) ([][][]float64, [][]float64) {
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

				wd[l][n][i] = ((bd[l][i]) * layers[l][n])
			}
		}

		if l == 0 {
			continue
		}

		layer = layers[l]
		errorcp := make([]float64, len(layer))
		for i := range layer {
			var err float64 = 0.0
			for j := range errors {
				err += weights[l][i][j] * errors[j]
			}
			errorcp[i] = err
		}
		errors = errorcp

	}

	return wd, bd
}
func Train(learningRate float64, mathFuncs []string, weights [][][]float64, bias, dataset [][]float64, expected [][]float64, epochs int) ([][][]float64, [][]float64) {
	var wg sync.WaitGroup
	for i := 0; i < epochs; i++ {
		weightsGrad := make([][][][]float64, len(dataset))
		biasGrad := make([][][]float64, len(dataset))
		for j, v := range dataset {
			wg.Add(1)
			go func(j int, v []float64) {
				_, layers := FeedFoward(mathFuncs, v, weights, bias)

				wd, bd := BackPropagation(weights, bias, layers, expected[j], mathFuncs)
				weightsGrad[j] = wd
				biasGrad[j] = bd
				wg.Done()

			}(j, v)
		}
		wg.Wait()
		for j := range dataset {
			for l := 0; l < len(weights)-1; l++ {

				for n := 0; n < len(weights[l]); n++ {

					for i := range weights[l][n] {

						weights[l][n][i] -= (weightsGrad[j][l][n][i] / float64(len(dataset))) * learningRate
					}

				}
				for i := range bias[l] {
					bias[l][i] -= (biasGrad[j][l][i] / float64(len(dataset))) * learningRate
				}
				//	wg.Done()
				//}(j, v)
				//wg.Wait()
			}
		}

	}
	return weights, bias
}
