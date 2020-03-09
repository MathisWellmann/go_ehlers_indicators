package EhlersTAIndicators

import (
	"errors"
	"math"
)

// FisherTransform applies the fisher transform function to an array and outputs one with the same length
// Only accepts inputs in range (-1, 1)
// original pdf: https://www.mesasoftware.com/papers/UsingTheFisherTransform.pdf
func FisherTransform(vals []float64) ([]float64, error) {
	out := make([]float64, len(vals))

	for i := 0; i < len(out); i++ {
		if vals[i] > 1 || vals[i] < -1 {
			return out, errors.New("vals[i] > 1 || vals[i] < -1")
		}
		out[i] = 0.5 * math.Log((1+vals[i])/(1-vals[i]))
	}
	return out, nil
}
