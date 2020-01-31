package EhlersTAIndicators

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type (
	Values struct {
		Xs []float64
		Ys []float64
	}
)

func (v *Values) Len() int {
	return len(v.Xs)
}

func (v *Values) XY(index int) (float64, float64) {
	return v.Xs[index], v.Ys[index]
}

// Plt will create a line plot with given values and filename
func Plt(vals []float64, filename string) error {
	xs := make([]float64, len(vals))
	ys := make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		xs[i] = float64(i)
		ys[i] = vals[i]
	}
	values := &Values{
		Xs: xs,
		Ys: ys,
	}

	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = "line plot"

	line, err := plotter.NewLine(values)
	if err != nil {
		return err
	}
	p.Add(line)

	if err := p.Save(297*vg.Millimeter, 210*vg.Millimeter, filename); err != nil {
		return err
	}
	return nil
}
