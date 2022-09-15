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
	// the output doesnt have a weight
	// the same for the bias
	weights := make([][][]float32, len(neuronsPerLayer)-1)
	bias := make([][]float32, len(neuronsPerLayer)-1)

	for i := 0; i < len(neuronsPerLayer)-1; i++ {
		// this add new weights for each neuron
		for n := 0; n < neuronsPerLayer[i]; n++ {
			weights[i] = append(weights[i], []float32{})
			// as you should know the weights of the neuron connects to the next layer
			// so this is really basic
			for w := 0; w < neuronsPerLayer[i+1]; w++ {
				weights[i][n] = append(weights[i][n], rand.Float32()-0.5)
			}

		}
		// because the input doesnt have a bias , this is why it looks like this
		for n := 0; n < neuronsPerLayer[i+1]; n++ {
			bias[i] = append(bias[i], rand.Float32()-0.5)
		}
	}

	return NN{Weights: weights, Bias: bias, ActivationFuncs: activationFuncs, Comment: comment}
}

// this is like the predict function
// but it returns you the layers

func (net NN) FeedFoward(input []float32) (layers [][]float32) {
	layers = make([][]float32, len(net.Bias)+1)

	layers[0] = make([]float32, len(input))
	copy(layers[0], input)

	for l := 0; l < len(layers)-1; l++ {
		layers[l+1] = make([]float32, len(net.Bias[l]))
		// layer*weight
		for n := 0; n < len(layers[l]); n++ {
			for i, w := range net.Weights[l][n] {
				layers[l+1][i] += w * layers[l][n]

			}

		}
		//layer(l+1)=f(bias)
		for i := range layers[l+1] {
			layers[l+1][i] = MathFuncs[net.ActivationFuncs[l]]["activate"](layers[l+1][i] + net.Bias[l][i])

		}

	}
	return
}

//you know why this is for
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
// the size is for the importance of each gradient
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
	// so , this works for making stuff much faster
	// concurrence baby
	var wg sync.WaitGroup

	for i := 0; i < epochs; i++ {
		var acc, err float32 = 0.0, 0.0
		// so , this updating the weight and bias
		dbList := make([][][]float32, len(dataset))
		wdList := make([][][][]float32, len(dataset))

		for j, v := range dataset {
			// just add a new function for execution
			wg.Add(1)

			go func(j int, v []float32) {
				layers := net.FeedFoward(v)

				wd, bd := net.BackPropagation(layers, expected[j])
				// I dont need to calculate the cost and accuracy in each epoch
				// so this is for avoiding any kind of shit
				if i%10 == 0 && logs {
					acc += Accuracy(expected[j], layers[len(layers)-1]) / float32(len(dataset))
					err += Cost(expected[j], layers[len(layers)-1])
				}
				// i just add new stuff and that for training
				wdList[j], dbList[j] = wd, bd
				dbList[j] = bd
				// then it finish
				wg.Done()
			}(j, v)
		}
		// after all functions finish to execute it continues
		wg.Wait()
		for j := range dbList {
			// with this i get the gradient of each input, so each gradient have the same gradient
			net.UpdateWeightAndBias(float32(len(dataset)), learningRate, wdList[j], dbList[j])
		}
		if i%10 == 0 && logs {

			fmt.Printf("| cost: %9.5f | Accuracy: %9.5f  |epochs %d\n", err, acc, i)
		}

	}

}

// the model is saved in a json format
func (net *NN) SaveModel(name string) {

	f, _ := os.Create(name)

	json.NewEncoder(f).Encode(net)

}

// json format
func OpenModel(name string) NN {
	var net NN
	f, err := os.Open(name)
	if err != nil {
		panic(name + " doesnt exist ")
	}
	json.NewDecoder(f).Decode(&net)
	return net

}
