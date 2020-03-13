package EhlersTAIndicators

import (
	"fmt"
	"github.com/MathisWellmann/timeseries_generator"
	"testing"
)

func TestFAMAGraph(t *testing.T) {
	vals := timeseries_generator.GaussianProcess(1024)

	fama := FAMADefault(vals)
	filename := fmt.Sprintf("img/fama.png")
	err := Plt(fama, filename)
	if err != nil {
		t.Error(err)
	}
}
