package EhlersTAIndicators

import (
	"fmt"
	"github.com/MathisWellmann/timeseries_generator"
	"testing"
)

func TestFisherTransformGraph(t *testing.T) {
	vals := timeseries_generator.SineWave(1024)
	fish, err := FisherTransform(vals)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("fish: %v", fish)

	filename := fmt.Sprintf("img/fisher_transform.png")
	err = Plt(fish, filename)
	if err != nil {
		t.Error(err)
	}
}
