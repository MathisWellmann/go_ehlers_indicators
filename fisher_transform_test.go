package EhlersTAIndicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestFisherTransformGraph(t *testing.T) {
	vals := go_timeseries_generator.SineWave(1024)
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
