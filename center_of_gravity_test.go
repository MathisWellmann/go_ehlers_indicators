package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestCenterOfGravity(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cgo := CenterOfGravity(vals, 16)
	for i := 0; i < len(cgo); i++ {
		fmt.Printf("cgo[%d]: %f\n", i, cgo[i])
	}
}

func TestCenterOfGravityGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cgo := CenterOfGravity(vals, 16)
	filename := fmt.Sprintf("./img/center_of_gravity.png")
	err := Plt(cgo, filename)
	if err != nil {
		t.Error(err)
	}
}
