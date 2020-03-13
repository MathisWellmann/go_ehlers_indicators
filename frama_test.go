package EhlersTAIndicators

import (
	"fmt"
	"github.com/MathisWellmann/timeseries_generator"
	"testing"
)

func TestFRAMAGraph(t *testing.T) {
	highs, lows := timeseries_generator.GaussianHL(1024)
	frama, err := FRAMA(highs, lows, 16)
	if err != nil {
		t.Error(err)
	}

	filename := fmt.Sprint("img/frama.png")
	err = Plt(frama, filename)
	if err != nil {
		t.Error(err)
	}
}
