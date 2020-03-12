package EhlersTAIndicators

import (
	"fmt"
	"testing"
)

func TestFRAMAGraph(t *testing.T) {
	highs, lows := GenerateGaussianHL(1024)
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
