package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestTrendFlex(t *testing.T) {
	vals := GenerateTimeseries(1024)
	tf := TrendFlex(vals, 16)
	for i := 0; i < len(tf); i++ {
		fmt.Printf("trendflex[%d]: %f\n", i, tf[i])
	}
}

func TestTrendFlexGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)
	tf := TrendFlex(vals, 16)
	err := checkImgDirectory()
	if err != nil {
		t.Error(err)
	}
	filename := fmt.Sprintf("./img/TestTrendFlexGraph.png")
	err = Plt(tf, filename)
	if err != nil {
		t.Error(err)
	}
}
