package EhlersTAIndicators

import "math"

// InverseFisherTransform applies the Tanh function to all values in input series and returns an array of same length
func InverseFisherTransform(vals []float64) []float64 {
	out := make([]float64, len(vals))

	for i := 0; i < len(out); i++ {
		out[i] = math.Tanh(vals[i])
	}
	return out
}
