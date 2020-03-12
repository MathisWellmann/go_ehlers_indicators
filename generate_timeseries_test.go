package EhlersTAIndicators

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateTimeseries(t *testing.T) {
	vals := GenerateTimeseries(1024)
	for i := 0; i < len(vals); i++ {
		fmt.Printf("vals[%d]: %f\n", i, vals[i])
		if vals[i] < 0 {
			t.Error("vals[i] < 0")
		}
	}
}

func TestGenerateTimeseriesGraph(t *testing.T) {
	vals := GenerateTimeseries(1024)

	err := checkImgDirectory()
	if err != nil {
		t.Error(err)
	}
	filename := fmt.Sprintf("img/TestGenerateTimeseriesGraph.png")
	err = Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}

func TestRandChange(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := RandChange()
		fmt.Printf("r: %f\n", r)
		if r > 1 || r < -1 {
			t.Error("r out of desired range")
		}
	}
}

// checkImgDirectory will make sure the img directory is available for any tests using it in the file path
func checkImgDirectory() error {
	// check if img directory already exists
	_, err := os.Stat("./img")
	if err != nil {
		if os.IsNotExist(err) {
			// create directory img
			err := os.Mkdir("./img", os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func TestGenerateSinewave(t *testing.T) {
	sine := GenerateSineWave(1024)

	filename := fmt.Sprintf("img/sinewave.png")
	err := Plt(sine, filename)
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateGaussianTimeseries(t *testing.T) {
	vals := GenerateGaussianTimeseries(1024)

	filename := fmt.Sprintf("img/TestGenerateGaussianTimeseris.png")
	err := Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
