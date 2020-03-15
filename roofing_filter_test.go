package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestRoofingFilter(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	filt := RoofingFilter(vals)

	filename := fmt.Sprintf("img/roofing_filter.png")
	err := Plt(filt, filename)
	if err != nil {
		t.Error(err)
	}
}
