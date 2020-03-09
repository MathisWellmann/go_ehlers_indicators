package EhlersTAIndicators

import "math"

// CyberCycle Indicator with alpha set to standard ema alpha: 2.0 / (windowLen + 1)
//
func CyberCycle(vals []float64, windowLen int) []float64 {
	out := make([]float64, len(vals))

	alpha := 2.0 / float64(windowLen+1)
	smooth := make([]float64, len(vals))
	for i := windowLen; i < len(out); i++ {
		smooth[i] = (vals[i] + 2*vals[i-1] + 2*vals[i-2] + vals[i-3]) / 6

		cc := math.Pow(1-0.5*alpha, 2)*(smooth[i]-2*smooth[i-1]+smooth[i-2]) + 2*(1-alpha)*out[i-1] - math.Pow(1-alpha, 2)*out[i-2]
		out[i] = cc
	}

	return out
}

// CyberCycle Indicator with custom alpha
func CyberCycleAlpha(vals []float64, windowLen int, alpha float64) []float64 {
	out := make([]float64, len(vals))

	for i := windowLen; i < len(out); i++ {
		smooth := make([]float64, windowLen)
		for j := 3; j < len(smooth); j++ {
			smooth[i] = (vals[i-j] + 2*vals[i-j-1] + 2*vals[i-j-2] + vals[i-j-3]) / 6
		}

		cc := math.Pow(1-0.5*alpha, 2)*(smooth[windowLen-1]-2*smooth[windowLen-2]) + 2*(1-alpha)*out[i-1] - math.Pow(1-alpha, 2)*out[i-2]
		out[i] = cc
	}

	return out
}
