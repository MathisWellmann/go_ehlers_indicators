package EhlersTAIndicators

import (
	"math"
	"math/rand"
)

// GenerateTimeseries will return a randomly generated positive path
func GenerateTimeseries(length int) []float64 {
	out := make([]float64, length)
	out[0] = rand.Float64()
	for i := 1; i < len(out); i++ {
		change := RandChange()
		if out[i-1]+change <= 0 {
			change = math.Abs(change)
		}
		out[i] = out[i-1] + change
	}
	return out
}

// randChange returns a random number in the range (-1, 1)
func RandChange() float64 {
	r := rand.Float64()
	// with 50/50 probability
	if rand.Float64() < 0.5 {
		r *= -1 // make random number negative
	}
	return r
}
