package EhlersTAIndicators

// CenterOfGravity Oscillator or CGOscillator by John Ehlers
//
func CenterOfGravity(vals []float64, windowLen int) []float64 {
	out := make([]float64, len(vals))

	for i := windowLen; i < len(out); i++ {
		var denom float64
		var num float64
		for j := 0; j < windowLen; j++ {
			weight := 1.0 + float64(windowLen-1-j)
			num += weight * vals[i-j]
			denom += vals[i-j]
		}
		var cgo float64
		if denom != 0 {
			cgo = -num/denom + (float64(windowLen)+1.0)/2.0 // add term to center cgo around zero
		}
		out[i] = cgo
	}

	return out
}
