package EhlersTAIndicators

import "math"

// ReFlex indicator
// from: https://financial-hacker.com/petra-on-programming-a-new-zero-lag-indicator/
func ReFlex(vals []float64, windowLen int) []float64 {
	out := make([]float64, len(vals))

	flts := make([]float64, len(vals))
	mss := make([]float64, len(vals))
	for i := windowLen; i < len(out); i++ {
		a1 := math.Exp(-1.414 * 2.0 * math.Pi / float64(windowLen))
		b1 := 2.0 * a1 * math.Cos(1.414*math.Pi/float64(windowLen))
		c2 := b1 // keeping this like the original paper
		c3 := -a1 * a1
		c1 := 1 - c2 - c3

		flt2 := flts[i-2]
		flt1 := flts[i-1]
		flt := c1*(vals[i]+vals[i-1])/2 + c2*flt1 + c3*flt2

		flts[i] = flt

		slope := (flts[i-windowLen] - flt) / float64(windowLen)

		// sum the differences
		var dSum float64
		for j := 0; j < windowLen; j++ {
			dSum += (flt + float64(j)*slope) - flts[i-j]
		}
		dSum /= float64(windowLen)

		// normalize in terms of standard deviations
		ms0 := 0.04*dSum*dSum + 0.96*mss[i-1]
		mss[i] = ms0

		var reflex float64
		if ms0 > 0.0 {
			reflex = dSum / math.Sqrt(ms0)
		}

		out[i] = reflex
	}

	return out
}
