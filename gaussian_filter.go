package go_ehlers_indicators

import (
	"math"
)

// GaussianFilter from this paper: https://www.mesasoftware.com/papers/GaussianFilters.pdf
func GaussianFilter(vals []float64, period, numPoles int) []float64 {
	filt := make([]float64, len(vals))

	for i := 4; i < len(filt); i++ {
		beta := (1 - math.Cos(2*math.Pi/float64(period))) / (math.Pow(2, 1.0/float64(numPoles)) - 1.0)
		alpha := -beta + math.Sqrt(math.Pow(beta, 2)+2*beta)

		if numPoles == 1 {
			filt[i] = alpha*vals[i] + (1-alpha)*filt[i-1]
		} else if numPoles == 2 {
			filt[i] = math.Pow(alpha, 2)*vals[i] + 2*(1-alpha)*filt[i-1] - math.Pow(1-alpha, 2)*filt[i-2]
		} else if numPoles == 3 {
			filt[i] = math.Pow(alpha, 3)*vals[i] + 3*(1-alpha)*filt[i-1] - 3*math.Pow(1-alpha, 2)*filt[i-2] + math.Pow(1-alpha, 3)*filt[i-3]
		} else if numPoles == 4 {
			filt[i] = math.Pow(alpha, 4)*vals[i] + 4*(1-alpha)*filt[i-1] - 6*math.Pow(1-alpha, 2)*filt[i-2] + 4*math.Pow(1-alpha, 3)*filt[i-3] - math.Pow(1-alpha, 4)*filt[i-4]
		}
	}
	return filt
}
