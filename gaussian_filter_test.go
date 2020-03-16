package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestGaussianFilter(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)

	for poles := 1; poles < 5; poles++ {
		filt := GaussianFilter(vals, 16, poles)
		filename := fmt.Sprintf("img/gaussian_filter_p%d.png", poles)
		err := Plt(filt, filename)
		if err != nil {
			t.Error(err)
		}
	}
}
