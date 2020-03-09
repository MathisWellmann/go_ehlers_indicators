package EhlersTAIndicators

import "math"

// FractalAdaptiveMA is a moving average inspired by the fractal nature of markets
// link to paper: http://www.stockspotter.com/Files/frama.pdf
func FractalAdaptiveMA(vals []float64, windowLen int) []float64 {
	out := make([]float64, len(vals))
	for i := 0; i < len(out); i++ {
		if i <= windowLen {
			out[i] = vals[i]
			continue
		}
		high := vals[i]
		low := vals[i]
		for j := 0; j < windowLen; j++ {
			if vals[i-j] > high {
				high = vals[i-j]
			} else if vals[i-j] < low {
				low = vals[i-j]
			}
		}
		n3 := (high - low) / float64(windowLen)

		// calculate high and low of first half onf windowLen
		highFirst := vals[i]
		lowFirst := vals[i]
		for j := 0; j < windowLen/2-1; j++ {
			if vals[i-j] > highFirst {
				highFirst = vals[i-j]
			} else if vals[i-j] < lowFirst {
				lowFirst = vals[i-j]
			}
		}
		n1 := (highFirst - lowFirst) / float64(windowLen) / 2

		// calculate high and low of last half of windowLen
		highLast := vals[i]
		lowLast := vals[i]
		for j := windowLen / 2; j < windowLen; j++ {
			if vals[i-j] > highFirst {
				highFirst = vals[i-j]
			} else if vals[i-j] < lowFirst {
				lowFirst = vals[i-j]
			}
		}
		n2 := (highLast - lowLast) / float64(windowLen) / 2

		var dimen float64
		if n1 > 0 && n2 > 0 && n3 > 0 {
			dimen = (math.Log(n1+n2) - math.Log(n3)) / math.Log(2)
		}
		alpha := math.Exp(-4.6 * (dimen - 1))
		if alpha < 0.01 {
			alpha = 0.01
		} else if alpha > 1 {
			alpha = 1
		}
		frama := alpha*vals[i] + (1-alpha)*out[i-1]
		if i <= windowLen {
			frama = vals[i]
		}
		out[i] = frama
	}
	return vals
}
