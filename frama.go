package EhlersTAIndicators

import (
	"errors"
	"math"
)

// FRAMA (fractal adaptive moving average) from paper: http://mesasoftware.com/papers/FRAMA.pdf
func FRAMA(highs, lows []float64, length int) ([]float64, error) {
	if math.Mod(float64(length), 2) != 0 {
		return nil, errors.New("length must be an even number")
	}
	if len(highs) != len(lows) {
		return nil, errors.New("len(highs) != len(lows)")
	}
	out := make([]float64, len(highs))

	for i := 0; i < len(out); i++ {
		if i < length+1 {
			avgPrice := (highs[i] + lows[i]) / 2
			out[i] = avgPrice
			continue
		}
		// highest high over last n datapoints, n being length
		_, h := Extent(highs[i-length : i])
		// lowest low over last n datapoints, n being length
		l, _ := Extent(lows[i-length : i])

		n3 := (h - l) / float64(length)

		hh := highs[i]
		ll := lows[i]
		for c := 0; c < (length/2)-1; c++ {
			if highs[c] > hh {
				hh = highs[c]
			}
			if lows[c] < ll {
				ll = lows[c]
			}
		}
		n1 := (hh - ll) / (float64(length) / 2.0)

		hh = highs[(i-length)/2]
		ll = lows[(i-length)/2]
		for c := length / 2; c < length-1; c++ {
			if highs[i] > hh {
				hh = highs[i]
			}
			if lows[i] < ll {
				ll = lows[i]
			}
		}
		n2 := (hh - ll) / (float64(length) / 2.0)

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

		avgPrice := (highs[i] + lows[i]) / 2
		out[i] = alpha*avgPrice + (1-alpha)*out[i-1]
	}

	return out, nil
}
