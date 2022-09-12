package brain

import (
	"math"
)

func relu(x float32) float32 {
	return float32(math.Max(0, float64(x)))

}
func sigmoid(x float32) float32 {

	return float32(1 / (1 + math.Exp(float64(x)*(-1))))

}

func tanh(x float32) float32 {
	return float32(math.Tanh(float64(x)))
}

func devSigmoid(x float32) float32 {
	return x * (1 - x)
}
func devRelu(x float32) float32 {
	out := 0.0
	if x > 0 {
		out = 1
	}
	return float32(out)
}
func devTanh(x float32) float32 {
	return 1 - (x * x)

}

var (
	MathFuncs = map[string]map[string]func(float32) float32{
		"sigmoid": {
			"derivative": devSigmoid,
			"activate":   sigmoid,
		},
		"relu": {
			"derivative": devRelu,
			"activate":   relu,
		},
		"tanh": {
			"derivative": devTanh,
			"activate":   tanh,
		},
	}
)

func Accuracy(target []float32, output []float32) float32 {
	acc := 0
	for i := range output {
		if (math.Round(float64(target[i]))) == (math.Round(float64(output[i]))) {
			acc++
		}
	}

	return float32(acc) / float32(len(output))
}

func Cost(target []float32, output []float32) float32 {
	err := 0.0
	for i := range target {
		err += math.Pow(float64(output[i]-target[i]), 2)
	}
	return float32(err)
}
