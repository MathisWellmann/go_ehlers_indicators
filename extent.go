package EhlersTAIndicators

// Extent returns the min and max values of values
func Extent(vals []float64) (float64, float64) {
	high := vals[0]
	low := vals[0]
	for i := 0; i < len(vals); i++ {
		if vals[i] > high {
			high = vals[i]
		} else if vals[i] < low {
			low = vals[i]
		}
	}
	return low, high
}
