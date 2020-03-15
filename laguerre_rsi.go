package go_ehlers_indicators

// LaguerreRSI from paper: http://mesasoftware.com/papers/TimeWarp.pdf
func LaguerreRSI(vals []float64, gamma float64) []float64 {
	l0s := make([]float64, len(vals))
	l1s := make([]float64, len(vals))
	l2s := make([]float64, len(vals))
	l3s := make([]float64, len(vals))
	rsi := make([]float64, len(vals))

	for i := 1; i < len(vals); i++ {
		l0s[i] = (1-gamma)*vals[i] + gamma*l0s[i-1]
		l1s[i] = -gamma*l0s[i] + l0s[i-1] + gamma*l1s[i-1]
		l2s[i] = -gamma*l1s[i] + l1s[i-1] + gamma*l2s[i-1]
		l3s[i] = -gamma*l2s[i] + l2s[i-1] + gamma*l3s[i-1]

		var cu float64
		var cd float64
		if l0s[i] >= l1s[i] {
			cu = l0s[i] - l1s[i]
		} else {
			cd = l1s[i] - l0s[i]
		}
		if l1s[i] >= l2s[i] {
			cu += +l1s[i] - l2s[i]
		} else {
			cd += l2s[i] - l1s[i]
		}
		if l2s[i] >= l3s[i] {
			cu += l2s[i] - l3s[i]
		} else {
			cd += l3s[i] - l2s[i]
		}

		if cu+cd != 0 {
			rsi[i] = cu / (cu + cd)
		}
	}
	return rsi
}

func LaguerreRSIDefault(vals []float64) []float64 {
	return LaguerreRSI(vals, 0.5)
}
