package EhlersTAIndicators

import "math"

// ZeroLag filter from this paper: http://mesasoftware.com/papers/ZeroLag.pdf
func ZeroLag(vals []float64, length, gainLimit int) []float64 {
	ema := make([]float64, len(vals))
	ecs := make([]float64, len(vals))

	for i := 0; i < len(ecs); i++ {
		alpha := 2.0 / (float64(length) + 1.0)
		ema[i] = alpha*vals[i] + (1.0-alpha)*ema[i-1]
		leastErr := 1e7
		var bestGain float64

		for v := -gainLimit; v < gainLimit; v++ {
			gain := float64(v) / 10.0
			ecs[i] = alpha*(ema[i]+gain*(vals[i]-ecs[i-1])) + (1.0-alpha)*ecs[i-1]
			err := vals[i] - ecs[i]
			absErr := math.Abs(err)
			if absErr < leastErr {
				leastErr = absErr
				bestGain = gain
			}
		}
		ecs[i] = alpha*(ema[i]+bestGain*(vals[i]-ecs[i-1])) + (1-alpha)*ecs[i-1]
	}
	return ecs
}

func ZeroLagDefault(vals []float64, length int) []float64 {
	return ZeroLag(vals, length, 50)
}
