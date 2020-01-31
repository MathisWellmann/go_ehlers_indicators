package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestReFlex(t *testing.T) {
	vals := GenerateTimeseries(1024)
	reflex := ReFlex(vals, 16)
	for i := 0; i < len(reflex); i++ {
		fmt.Sprintf("reflex[%d]: %f", i, reflex)
	}
}

func TestReFlexGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)
	reflex := ReFlex(vals, 16)
	err := checkImgDirectory()
	if err != nil {
		t.Error(err)
	}
	filename := fmt.Sprintf("./img/TestReFlexGraph.png")
	err = Plt(reflex, filename)
	if err != nil {
		t.Error(err)
	}
}
