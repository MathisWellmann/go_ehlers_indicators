package EhlersTAIndicators

import (
	"fmt"
	"math"
)

const (
	M = iota
	F
)

// MAMA from this paper: https://www.mesasoftware.com/papers/MAMA.pdf
func MAMAFAMA(vals []float64, fastLimit, slowLimit float64, which int) []float64 {
	smooth := make([]float64, len(vals))
	period := make([]float64, len(vals))
	detrender := make([]float64, len(vals))
	is := make([]float64, len(vals))
	i2s := make([]float64, len(vals))
	qs := make([]float64, len(vals))
	q2s := make([]float64, len(vals))
	res := make([]float64, len(vals))
	ims := make([]float64, len(vals))
	smoothPeriod := make([]float64, len(vals))
	phase := make([]float64, len(vals))
	mama := make([]float64, len(vals))
	fama := make([]float64, len(vals))

	for i := 0; i < len(vals); i++ {
		if i <= 5 {
			mama[i] = vals[i]
			fama[i] = vals[i]
			continue
		}
		smooth[i] = (4*vals[i] + 3*vals[i-1] + 2*vals[i-2] + vals[i-3]) / 10
		detrender[i] = (0.0962*smooth[i] + 0.5769*smooth[i-2] - 0.5769*smooth[i-4] - 0.0962*smooth[i-6]) * (0.075*period[i-1] + 0.54)

		// compute InPhase and Quadrature components
		q1 := (0.0962*detrender[i] + 0.5769*detrender[i-2] - 0.5769*detrender[i-4] - 0.0962*detrender[i-6]) * (0.075*period[i-1] + 0.54)
		qs[i] = q1
		is[i] = detrender[i-3]

		// Advance the phase of detrender and q1 by 90 Degrees
		jI := (0.0962*is[i] + 0.05769*is[i-2] - 0.5769*is[i-4] - 0.0962*is[i-6]) * (0.075*period[i-1] + 0.54)
		jQ := (0.0962*qs[i] + 0.5769*qs[i-2] - 0.5769*qs[i-4] - 0.0962*qs[i-6]) * (0.075*period[i-1] + 0.54)

		// Phasor addition for 3 bar averaging
		i2 := detrender[i-3] - jQ
		i2s[i] = i2
		q2 := q1 + jI
		q2s[i] = q2

		// smooth the I and Q components befor applying the discriminator
		i2 = 0.2*i2 + 0.8*i2s[i-1]
		i2s[i] = i2
		q2 = 0.2*q2 + 0.8*qs[i-1]
		q2s[i] = q2

		// Homodyne Discriminator
		re := i2*i2s[i-1] + q2*q2s[i-1]
		res[i] = re
		im := i2*q2s[i-1] - q2*i2s[i-1]
		ims[i] = im
		re = 0.2*re + 0.8*res[i-1]
		im = 0.2*im + 0.8*ims[i-1]

		if im != 0 && re != 0 {
			period[i] = 360 / math.Atan(im/re)
		}
		if period[i] > 1.5*period[i-1] {
			period[i] = 1.5 * period[i-1]
		}
		if period[i] < 0.67*period[i-1] {
			period[i] = 0.67 * period[i-1]
		}
		if period[i] < 6 {
			period[i] = 6
		}
		if period[i] > 50 {
			period[i] = 50
		}
		period[i] = 0.2*period[i] + 0.8*period[i-1]
		smoothPeriod[i] = 0.33*period[i] + 0.67*smoothPeriod[i-1]

		if detrender[i-3] != 0 {
			phase[i] = math.Atan(q1 / detrender[i-3])
		}
		deltaPhase := phase[i-1] - phase[i]
		if deltaPhase < 1 {
			deltaPhase = 1
		}
		alpha := fastLimit / deltaPhase
		if alpha < slowLimit {
			alpha = slowLimit
		}
		mama[i] = alpha*vals[i] + (1-alpha)*mama[i-1]
		fama[i] = 0.5*alpha*mama[i] + (1-0.5*alpha)*fama[i-1]

	}
	if which == M {
		return mama
	} else if which == F {
		return fama
	}
	fmt.Printf("Wrong 'which' parameter provided. accepted are M for MAMA or F for MAMA constants")
	return nil
}

// MAMADefault provides a wrapper for MAMA with recommended fastsLimit of 0.5 and slowLimit of 0.05
func MAMADefault(vals []float64) []float64 {
	fastLimit := 0.5
	slowLimit := 0.05
	return MAMA(vals, fastLimit, slowLimit)
}

// MAMA (MESA adaptive moving average)
func MAMA(vals []float64, fastLimit, slowLimit float64) []float64 {
	return MAMAFAMA(vals, fastLimit, slowLimit, M)
}
