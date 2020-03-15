package EhlersTAIndicators

func LaguerreFilter(vals []float64, gamma float64) []float64 {
	l0s := make([]float64, len(vals))
	l1s := make([]float64, len(vals))
	l2s := make([]float64, len(vals))
	l3s := make([]float64, len(vals))
	filts := make([]float64, len(vals))
	// firs := make([]float64, len(vals))

	for i := 1; i < len(vals); i++ {
		l0s[i] = (1.0-gamma)*vals[i] + gamma*l0s[i-1]
		l1s[i] = -gamma*l0s[i] + l0s[i-1] + gamma*l1s[i-1]
		l2s[i] = -gamma*l1s[i] + l1s[i-1] + gamma*l2s[i-1]
		l3s[i] = -gamma*l2s[i] + l2s[i-1] + gamma*l3s[i-1]

		filts[i] = (l0s[i] + 2*l1s[i] + 2*l2s[i] + l3s[i]) / 6
		// firs[i] = (vals[i] + 2*vals[i-1] + 2*vals[i-2] + vals[i-3]) / 6
	}

	return filts
}

func LaguerreFilterDefault(vals []float64) []float64 {
	return LaguerreFilter(vals, 0.8)
}
