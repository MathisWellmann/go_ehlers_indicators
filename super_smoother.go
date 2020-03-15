package go_ehlers_indicators

import "math"

// SuperSmoother from this paper: http://www.mesasoftware.com/papers/PredictiveIndicatorsForEffectiveTrading%20Strategies.pdf
func SuperSmoother(vals []float64, length int) []float64 {
	filt := make([]float64, len(vals))

	for i := 2; i < len(vals); i++ {
		a1 := math.Exp(-1.414 * 3.14159 / float64(length))
		// Cos() argument has been converted from degrees to radians as mentioned in the paper
		b1 := 2 * a1 * math.Cos(4.44/float64(length))

		c3 := -a1 * a1
		c1 := 1 - b1 - c3
		filt[i] = c1*(vals[i]+vals[i-1])/2 + b1*filt[i-1] + c3*filt[i-2]
	}
	return filt
}
