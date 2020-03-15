package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestLaguerreFilter(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	l := LaguerreFilterDefault(vals)

	filename := fmt.Sprintf("img/laguerre_filter.png")
	err := Plt(l, filename)
	if err != nil {
		t.Error(err)
	}
}
