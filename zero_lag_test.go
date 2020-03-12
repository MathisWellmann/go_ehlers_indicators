package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestZeroLagGraph(t *testing.T) {
	vals := GenerateGaussianTimeseries(1024)
	zl := ZeroLagDefault(vals, 16)

	filename := fmt.Sprintf("img/zero_lag.png")
	err := Plt(zl, filename)
	if err != nil {
		t.Error(err)
	}
}
