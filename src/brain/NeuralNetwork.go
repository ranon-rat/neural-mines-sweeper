package brain

import (
	"math/rand"
	"time"
)

func NeuralNetwork(neuronsPerLayer []int) (weights [][][]float64, bias [][]float64) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(neuronsPerLayer); i++ {
		bias = append(bias, []float64{})
		weights = append(weights, [][]float64{})
		for n := 0; n < neuronsPerLayer[i]; n++ {
			// i dont need to know the bias for the input but whatever
			if i != 0 {
				bias[i] = append(bias[i], rand.Float64())
			}
			if i < len(neuronsPerLayer)-1 {
				weights[i] = append(weights[i], []float64{})
				for w := 0; w < neuronsPerLayer[i+1]; w++ {
					weights[i][n] = append(weights[i][n], rand.Float64())
				}
			}
		}
	}

	return
}

func FeedFoward(mathFuncPerLayer []string, input []float64, weights [][][]float64, bias [][]float64) (output []float64, layers [][]float64) {
	output = input
	layers = make([][]float64, len(bias))

	layers[0] = input
	for l := 0; l < len(weights)-1; l++ {

		for n := 0; n < len(output); n++ {

			layers[l+1] = make([]float64, len(bias[l+1]))
			for i, w := range weights[l][n] {
				layers[l+1][i] += w * output[n]

			}

		}
		for i, n := range layers[l+1] {
			layers[l+1][i] = MathFuncs[mathFuncPerLayer[l]]["activate"](n + bias[l+1][i])

		}

		output = layers[l+1]

	}
	return
}

// what im doing with my life :weary:
// how the fuck i did this ?
func BackPropagation(learningRate float64, weights [][][]float64, bias, layers [][]float64, expected []float64, mathFuncPerLayer []string) [][]float64 {
	bm := make([][]float64, len(bias))
	errors := make([]float64, len(expected))
	layer := layers[len(layers)-1]
	for i, n := range layers[len(layers)-1] {
		errors[i] = expected[i] - n

	}
	for l := len(layers) - 2; l >= 0; l-- {
		bm[l+1] = make([]float64, len(bias[l+1]))

		for i := range bias[l+1] {

			bm[l+1][i] = (errors[i] * MathFuncs[mathFuncPerLayer[l]]["derivative"](layer[i]))
		}

		if l == 0 {
			continue
		}

		layer = layers[l]
		errors = make([]float64, len(layer))
		for i := range layer {
			var err float64 = 0.0
			for j, n := range bias[l+1] {
				err += weights[l][i][j] * n
			}
			errors[i] = err
		}
	}

	return bm
}
func Train(learningRate float64, mathFuncs []string, weights [][][]float64, bias, dataset [][]float64, expected [][]float64, epochs int) ([][][]float64, [][]float64) {
	//var wg sync.WaitGroup
	for i := 0; i < epochs; i++ {

		for j, v := range dataset {
			_, layers := FeedFoward(mathFuncs, v, weights, bias)
			bd := BackPropagation(learningRate, weights, bias, layers, expected[j], mathFuncs)

			for l := 0; l < len(weights)-1; l++ {

				for n := 0; n < len(weights[l]); n++ {

					for i := range weights[l][n] {

						weights[l][n][i] += ((bd[l+1][i]) * layers[l][n] * learningRate) // float64(len(dataset))
					}

				}
				for i := range bias[l+1] {
					bias[l+1][i] = bd[l+1][i] // float64(len(dataset)+epochs)
				}
				//	wg.Done()
				//}(j, v)
				//wg.Wait()
			}

		}

	}
	return weights, bias
}
