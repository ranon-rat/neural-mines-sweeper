package brain

import (
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
				weights[i][n] = append(weights[i][n], rand.Float64()-0.5)
			}

		}
		for n := 0; n < neuronsPerLayer[i+1]; n++ {
			// i dont need to know the bias for the input but whatever
			bias[i] = append(bias[i], rand.Float64()-0.5)
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
		for i := range layers[l+1] {
			layers[l+1][i] = MathFuncs[mathFuncPerLayer[l]]["activate"](layers[l+1][i] + bias[l][i])

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
	// i get the errors  doing a really simple thing
	for i, n := range layer {
		errors[i] = n - expected[i]

	}

	for l := len(bias) - 1; l >= 0; l-- {
		bd[l] = make([]float64, len(bias[l]))
		wd[l] = make([][]float64, len(weights[l]))

		// okay This supose to get the gradient
		// I dont really know why its not working properly

		for i := range bias[l] {
			// so , the gradient its added to the bias
			// so , by doing this its correct
			bd[l][i] = (errors[i] * MathFuncs[mathFuncPerLayer[l]]["derivative"](layer[i]))
		}
		// okay , so I multiply layers_lt*gradient
		//
		for n := 0; n < len(weights[l]); n++ {
			wd[l][n] = make([]float64, len(weights[l][n]))

			for i := range weights[l][n] {

				wd[l][n][i] = layers[l][n] * (bd[l][i])
			}
		}

		if l == 0 {
			continue
		}

		layer = layers[l]
		errorcp := make([]float64, len(layer))
		// i just multiply the errors by the weight of this layer
		// it supose to get the value that i Want but I dont really know why its not working or maybe it is but im too stupid for find that

		for i := range layer {
			// this supose to get me the errors of the l-1
			// but maybe something its not working properly
			// but I dont really know
			err := 0.0
			for j := range errors {
				// I just
				err += weights[l][i][j] * errors[j]
			}
			errorcp[i] = err
		}
		errors = errorcp

	}

	return wd, bd
}
func UpdateWeightAndBias(size, learningRate float64, weights [][][]float64, bias [][]float64, weightsGrad [][][]float64, biasGrad [][]float64) ([][][]float64, [][]float64) {

	for l := 0; l < len(weights); l++ {

		for n := 0; n < len(weights[l]); n++ {

			for i := range weights[l][n] {

				weights[l][n][i] -= ((weightsGrad[l][n][i]) * learningRate) / size
			}

		}
		for i := range bias[l] {
			bias[l][i] -= ((biasGrad[l][i]) * learningRate) / size
		}

	}

	return weights, bias
}
func Train(learningRate float64, mathFuncs []string, weights [][][]float64, bias, dataset [][]float64, expected [][]float64, epochs int) ([][][]float64, [][]float64) {
	for i := 0; i < epochs; i++ {

		//bdSum := make([][]float64, len(bias))
		//wdSum := make([][][]float64, len(weights))
		for j, v := range dataset {
			_, layers := FeedFoward(v, mathFuncs, weights, bias)

			wd, bd := BackPropagation(weights, bias, layers, expected[j], mathFuncs)
			weights, bias = UpdateWeightAndBias(1, learningRate, weights, bias, wd, bd)

			//	for l := len(layers) - 2; l >= 0; l-- {
			//		if len(bdSum[l]) == 0 {
			//			bdSum[l] = make([]float64, len(bias[l]))
			//			wdSum[l] = make([][]float64, len(weights[l]))
			//		}
			//		for i := range bias[l] {
			//
			//			bdSum[l][i] += bd[l][i]
			//		}
			//		for n := 0; n < len(weights[l]); n++ {
			//			if len(wdSum[l][n]) == 0 {
			//				wdSum[l][n] = make([]float64, len(weights[l][n]))
			//			}
			//			for i := range weights[l][n] {
			//
			//				wdSum[l][n][i] += wd[l][n][i]
			//			}
			//		}
			//
			//	}

		}
		//		weights, bias = UpdateWeightAndBias(float64(len(dataset)), learningRate, weights, bias, wdSum, bdSum)

	}
	return weights, bias
}
