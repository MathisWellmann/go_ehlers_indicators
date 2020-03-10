package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestMAMAGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)

	mama := MAMADefault(vals)
	filename := fmt.Sprintf("img/mama.png")
	err := Plt(mama, filename)
	if err != nil {
		t.Error(err)
	}
}
