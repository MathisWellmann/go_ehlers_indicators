package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestMAMAGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)

	mama := MAMADefault(vals)
	filename := fmt.Sprintf("img/mama.png")
	err := Plt(mama, filename)
	if err != nil {
		t.Error(err)
	}
}
