package brain

import "math"

func relu(x float64) float64 {
	return (math.Max(0, (x)))

}
func sigmoid(x float64) float64 {

	return 1 / (1 + math.Exp(-x))

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
			"derivative": sigmoid,
			"activate":   devSigmoid,
		},
		"relu": {
			"derivative": relu,
			"activate":   devRelu,
		},
		"tanh": {
			"derivative": tanh,
			"activate":   devTanh,
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

// i already know that the mod function only works for integers
// but its something that i need for making this
// and golang dont permit me to do this kind of stuff so I need to do this
// sorry math bros

func mod(y, x float64) float64 {
	val := x
	for val > y {
		val -= y
	}
	return val
}