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
	Weights         [][][]float32 `json:"weights"          `
	Bias            [][]float32   `json:"bias"             `
	ActivationFuncs []string      `json:"activation-funcs" `
	Comment         string        `json:"comment" `
}

func NewNeuralNetwork(neuronsPerLayer []int, activationFuncs []string, comment string) NN {
	rand.Seed(time.Now().Unix())
	if len(activationFuncs)+1 < len(neuronsPerLayer) {
		panic("the activation funcs are different")
	}
	weights := make([][][]float32, len(neuronsPerLayer)-1)
	bias := make([][]float32, len(neuronsPerLayer)-1)
	for i := 0; i < len(neuronsPerLayer)-1; i++ {

		for n := 0; n < neuronsPerLayer[i]; n++ {
			weights[i] = append(weights[i], []float32{})
			for w := 0; w < neuronsPerLayer[i+1]; w++ {
				weights[i][n] = append(weights[i][n], rand.Float32()-0.5)
			}

		}
		for n := 0; n < neuronsPerLayer[i+1]; n++ {
			bias[i] = append(bias[i], rand.Float32()-0.5)
		}
	}

	return NN{Weights: weights, Bias: bias, ActivationFuncs: activationFuncs, Comment: comment}
}

func (net NN) FeedFoward(input []float32) (layers [][]float32) {
	layers = make([][]float32, len(net.Bias)+1)

	layers[0] = make([]float32, len(input))
	copy(layers[0], input)

	for l := 0; l < len(layers)-1; l++ {
		layers[l+1] = make([]float32, len(net.Bias[l]))
		copy(layers[l+1], net.Bias[l])
		// layer*weight
		for n := 0; n < len(layers[l]); n++ {

			for i, w := range net.Weights[l][n] {
				layers[l+1][i] += w * layers[l][n]

			}

		}
		//layer(l+1)=f(bias)
		for i := range layers[l+1] {
			layers[l+1][i] = MathFuncs[net.ActivationFuncs[l]]["activate"](layers[l+1][i])

		}

	}
	return
}
func (net NN) Predict(input []float32) []float32 {
	lays := net.FeedFoward(input)
	return lays[len(lays)-1]

}

func (net *NN) BackPropagation(layers [][]float32, expected []float32) ([][][]float32, [][]float32) {

	bd := make([][]float32, len(net.Bias))
	wd := make([][][]float32, len(net.Weights))
	errors := make([]float32, len(expected))
	layer := layers[len(layers)-1]
	// I dont need to explain this one
	for i, n := range layer {
		errors[i] = n - expected[i]
	}

	for l := len(net.Bias) - 1; l >= 0; l-- {
		bd[l] = make([]float32, len(net.Bias[l]))
		wd[l] = make([][]float32, len(net.Weights[l]))

		for i := range net.Bias[l] {
			//gradient=errors*dy/dx(fx)(layer[l+1])
			bd[l][i] += (errors[i] * MathFuncs[net.ActivationFuncs[l]]["derivative"](layer[i]))
		}
		//layer_t *gradient
		for n := 0; n < len(wd[l]); n++ {
			wd[l][n] = make([]float32, len(net.Weights[l][n]))

			for i := range net.Weights[l][n] {

				wd[l][n][i] += layers[l][n] * (bd[l][i])
			}
		}

		if l == 0 {
			continue
		}

		layer = layers[l]
		errorcp := make([]float32, len(layer))
		// errors=weights_t*errors
		for i := range layer {

			var err float32 = 0.0
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
func (net *NN) UpdateWeightAndBias(size, learningRate float32, weightsGrad [][][]float32, biasGrad [][]float32) {

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
func (net *NN) Train(dataset, expected [][]float32, learningRate float32, epochs int, logs bool) {
	var wg sync.WaitGroup
	for i := 0; i < epochs; i++ {
		var err float32 = 0.0
		dbList := make([][][]float32, len(dataset))
		wdList := make([][][][]float32, len(dataset))

		for j, v := range dataset {
			wg.Add(1)
			go func(j int, v []float32) {
				layers := net.FeedFoward(v)

				wd, bd := net.BackPropagation(layers, expected[j])
				if i%10 == 0 && logs {
					err += Cost(expected[j], layers[len(layers)-1])

				}
				wdList[j], dbList[j] = wd, bd
				dbList[j] = bd
				wg.Done()
			}(j, v)
		}
		wg.Wait()
		for j := range dbList {
			net.UpdateWeightAndBias(float32(len(dataset)), learningRate, wdList[j], dbList[j])
		}
		if i%10 == 0 && logs {

			fmt.Printf("| cost: %9.5f |epochs %d\n", err, i)
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
