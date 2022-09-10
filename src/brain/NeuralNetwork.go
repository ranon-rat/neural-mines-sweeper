package brain

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type NN struct {
	Weights         [][][]float64 `json:"weights"`
	Bias            [][]float64   `json:"bias"`
	ActivationFuncs []string      `json:"activation-funcs"`
}

func NewNeuralNetwork(neuronsPerLayer []int, activationFuncs []string) NN {
	rand.Seed(time.Now().Unix())
	weights := make([][][]float64, len(neuronsPerLayer)-1)
	bias := make([][]float64, len(neuronsPerLayer)-1)
	for i := 0; i < len(neuronsPerLayer)-1; i++ {

		for n := 0; n < neuronsPerLayer[i]; n++ {
			weights[i] = append(weights[i], []float64{})
			for w := 0; w < neuronsPerLayer[i+1]; w++ {
				weights[i][n] = append(weights[i][n], rand.Float64()-0.5)
			}

		}
		for n := 0; n < neuronsPerLayer[i+1]; n++ {
			bias[i] = append(bias[i], rand.Float64()-0.5)
		}
	}

	return NN{Weights: weights, Bias: bias, ActivationFuncs: activationFuncs}
}

func (net NN) FeedFoward(input []float64) (layers [][]float64) {
	fmt.Println(len(net.Bias))
	layers = make([][]float64, len(net.Bias)+1)

	layers[0] = make([]float64, len(input))
	copy(layers[0], input)

	for l := 0; l < len(layers)-1; l++ {
		fmt.Println(l)
		layers[l+1] = make([]float64, len(net.Bias[l]))
		copy(layers[l+1], net.Bias[l])

		for n := 0; n < len(layers[l]); n++ {

			for i, w := range net.Weights[l][n] {
				layers[l+1][i] += w * layers[l][n]

			}

		}

		for i := range layers[l+1] {
			layers[l+1][i] = MathFuncs[net.ActivationFuncs[l]]["activate"](layers[l+1][i])

		}

	}
	return
}

func (net *NN) BackPropagation(layers [][]float64, expected []float64) ([][][]float64, [][]float64) {

	bd := make([][]float64, len(net.Bias))
	wd := make([][][]float64, len(net.Weights))
	errors := make([]float64, len(expected))
	layer := layers[len(layers)-1]
	// I dont need to explain this one
	for i, n := range layer {
		errors[i] = n - expected[i]
	}

	for l := len(net.Bias) - 1; l >= 0; l-- {
		bd[l] = make([]float64, len(net.Bias[l]))
		wd[l] = make([][]float64, len(net.Weights[l]))

		for i := range net.Bias[l] {
			//gradient=errors*dy/dx(fx)(layer[l+1])
			bd[l][i] += (errors[i] * MathFuncs[net.ActivationFuncs[l]]["derivative"](layer[i]))
		}
		//layer_t *gradient
		for n := 0; n < len(wd[l]); n++ {
			wd[l][n] = make([]float64, len(net.Weights[l][n]))

			for i := range net.Weights[l][n] {

				wd[l][n][i] += layers[l][n] * (bd[l][i])
			}
		}

		if l == 0 {
			continue
		}

		layer = layers[l]
		errorcp := make([]float64, len(layer))
		// errors=weights_t*errors
		for i := range layer {

			err := 0.0
			for j := range errors {
				err += net.Weights[l][i][j] * errors[j]
			}
			errorcp[i] = err
		}
		errors = errorcp

	}

	return wd, bd
}

// so , this update the weights and bias
// yeah its really simple
func (net *NN) UpdateWeightAndBias(size, learningRate float64, weightsGrad [][][]float64, biasGrad [][]float64) {

	for l := 0; l < len(net.Weights); l++ {

		for n := 0; n < len(net.Weights[l]); n++ {

			for i := range net.Weights[l][n] {
				// this reduce the error

				net.Weights[l][n][i] -= ((weightsGrad[l][n][i]) * learningRate) / size
			}

		}
		// same for this
		for i := range net.Bias[l] {
			net.Bias[l][i] -= (biasGrad[l][i] * learningRate) / size
		}

	}

}
func (net *NN) Train(activationFuncs []string, weights [][][]float64, bias, dataset, expected [][]float64, learningRate float64, epochs int) {
	var wg sync.WaitGroup
	for i := 0; i < epochs; i++ {
		err := 0.0
		dbList := make([][][]float64, len(dataset))
		wdList := make([][][][]float64, len(dataset))

		for j, v := range dataset {
			wg.Add(1)
			go func(j int, v []float64) {
				layers := net.FeedFoward(v)

				wd, bd := net.BackPropagation(layers, expected[j])
				if i%10 == 0 {
					err += Cost(expected[j], layers[len(layers)-1])
				}
				wdList[j] = wd
				dbList[j] = bd
				wg.Done()
			}(j, v)
		}
		wg.Wait()
		for j := range dbList {
			net.UpdateWeightAndBias(float64(len(dataset)), learningRate, wdList[j], dbList[j])
		}
		if i%10 == 0 {
			fmt.Println("| epoch:", i, "| cost:", err/float64(len(dataset)), "|")
		}

	}

}

// dir+"/"+"name"+".json"
func (net *NN) SaveModel(name string) {

	f, _ := os.Create(name)
	json.NewEncoder(f).Encode(net)
}

func OpenModel(name string) NN {
	var net NN
	f, err := os.Open(name)
	if err != nil {
		panic(name + ".json doesnt exist ")
	}
	json.NewDecoder(f).Decode(&net)
	return net

}
