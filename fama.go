package EhlersTAIndicators

// FAMA
func FAMA(vals []float64, fastLimit, slowLimit float64) []float64 {
	return MAMAFAMA(vals, fastLimit, slowLimit, F)
}

// FAMADefault is a wrapper for FAMA with recommended settings for fastLimit and slowLimit
func FAMADefault(vals []float64) []float64 {
	fastLimit := 0.5
	slowLimit := 0.05
	return FAMA(vals, fastLimit, slowLimit)
}
