package go_ehlers_indicators

import "math"

func RoofingFilter(vals []float64) []float64 {
	filt := make([]float64, len(vals))
	hps := make([]float64, len(vals))

	for i := 2; i < len(vals); i++ {
		// Highpass filter cyclic components whose periods are shorter tahn 48 bars
		// converted cos and sin arguments from degrees to radians as mentioned in the paper
		alpha1 := (math.Cos(0.0925) + math.Sin(0.0925) - 1) / math.Cos(0.0925)
		hps[i] = math.Pow(1.0-alpha1/2.0, 2)*(vals[i]-2*vals[i-1]+vals[i-2]) + 2*(1-alpha1)*hps[i-1] - math.Pow(1-alpha1, 2)*hps[i-2]

		// smooth with a super smoother
		a1 := math.Exp(-1.414 * 3.14159 / 10)
		b1 := 2 * a1 * math.Cos(0.436)

		c3 := -a1 * a1
		c1 := 1 - b1 - c3
		filt[i] = c1*(hps[i]+hps[i-1])/2 + b1*filt[i-1] + c3*filt[i-2]
	}
	return filt
}
