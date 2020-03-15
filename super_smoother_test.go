package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestSuperSmoother(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	smooth := SuperSmoother(vals, 16)

	filename := fmt.Sprintf("img/super_smoother.png")
	err := Plt(smooth, filename)
	if err != nil {
		t.Error(err)
	}
}
