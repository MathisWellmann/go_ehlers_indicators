package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestLaguerreRSI(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(512)
	lrsi := LaguerreRSIDefault(vals)

	filename := fmt.Sprintf("img/laguerre_rsi.png")
	err := Plt(lrsi, filename)
	if err != nil {
		t.Error(err)
	}
}
