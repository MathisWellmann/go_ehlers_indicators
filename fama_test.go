package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestFAMAGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)

	fama := FAMADefault(vals)
	filename := fmt.Sprintf("img/fama.png")
	err := Plt(fama, filename)
	if err != nil {
		t.Error(err)
	}
}
