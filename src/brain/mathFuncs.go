package brain

import "math"

func relu(x float64) float64 {
	return (math.Max(0, (x)))

}
func sigmoid(x float64) float64 {

	return 1 / (1 + math.Exp(x*(-1)))

}

func tanh(x float64) float64 {
	return (math.Tanh((x)))
}

func devSigmoid(x float64) float64 {
	return x * (1 - x)
}
func devRelu(x float64) float64 {
	out := 0.0
	if x > 0 {
		out = 1
	}
	return (out)
}
func devTanh(x float64) float64 {
	return 1 - (x * x)

}

var (
	MathFuncs = map[string]map[string]func(float64) float64{
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

func subtract(x, y []float64) (out []float64) {
	out = make([]float64, len(x))
	for i := range x {
		out[i] = x[i] - y[i]
	}
	return
}

func Cost(target []float64, output []float64) float64 {
	err := 0.0
	for i := range target {
		err += math.Abs(output[i] - target[i])
	}
	return err
}
