package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestCyberCycle(t *testing.T) {
	vals := GenerateTimeseries(1024)
	cc := CyberCycle(vals, 16)
	for i := 0; i < len(cc); i++ {
		fmt.Printf("cc[%d]: %f\n", i, cc[i])
	}
}

func TestCyberCycleGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)
	cc := CyberCycle(vals, 16)

	err := checkImgDirectory()
	if err != nil {
		t.Error(err)
	}

	filename := fmt.Sprintf("./img/TestCyberCycleGraph.png")
	err = Plt(cc, filename)
	if err != nil {
		t.Error(err)
	}
}
