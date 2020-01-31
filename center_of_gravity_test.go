package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestCenterOfGravity(t *testing.T) {
	vals := GenerateTimeseries(1024)
	cgo := CenterOfGravity(vals, 16)
	for i := 0; i < len(cgo); i++ {
		fmt.Printf("cgo[%d]: %f\n", i, cgo[i])
	}
}

func TestCenterOfGravityGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)
	cgo := CenterOfGravity(vals, 16)
	err := checkImgDirectory()
	if err != nil {
		t.Error(err)
	}
	filename := fmt.Sprintf("./img/TestCenterOfGravityGraph.png")
	err = Plt(cgo, filename)
	if err != nil {
		t.Error(err)
	}
}
