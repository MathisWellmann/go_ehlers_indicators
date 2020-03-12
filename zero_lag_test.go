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

// TestZeroLagGraphStep tests the zero lag filter on a step function by graphing it
func TestZeroLagGraphStep(t *testing.T) {
	vals := GenerateStepFunction(1024, 500, 100)

	filename := fmt.Sprintf("img/TestZeroLagGraphStep.png")
	err := Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
