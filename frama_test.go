package EhlersTAIndicators

import (
	"fmt"
	"github.com/MathisWellmann/timeseries_generator"
	"testing"
)

func TestFRAMAGraph(t *testing.T) {
	candles := timeseries_generator.GaussianOHLCVDefault(1024)
	highs := make([]float64, len(candles))
	lows := make([]float64, len(candles))
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
