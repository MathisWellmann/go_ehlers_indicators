package EhlersTAIndicators

import "math"

// TrendFlex indicator
// from https://financial-hacker.com/petra-on-programming-a-new-zero-lag-indicator/
func TrendFlex(vals []float64, windowLen int) []float64 {
	out := make([]float64, len(vals))

	flts := make([]float64, len(out))
	mss := make([]float64, len(out))
	for i := windowLen; i < len(out); i++ {
		a1 := math.Exp(-1.414 * 2.0 * math.Pi / float64(windowLen))
		b1 := 2.0 * a1 * math.Cos(1.414*math.Pi/float64(windowLen))
		c2 := b1
		c3 := -a1 * a1
		c1 := 1 - c2 - c3

		flt := c1*(vals[i]+vals[i-1])/2 + c2*flts[i-1] + c3*flts[i-2]
		flts[i] = flt

		// sum the differences
		var dSum float64
		for j := 1; j < windowLen; j++ {
			dSum += flt - flts[i-j]
		}
		dSum /= float64(windowLen)

		// normalize in terms of standard deviations
		ms0 := 0.04*dSum*dSum + 0.96*mss[i-1]
		mss[i] = ms0

		var trendflex float64
		if ms0 > 0.0 {
			trendflex = dSum / math.Sqrt(ms0)
		}
		out[i] = trendflex
	}
	return out
}
