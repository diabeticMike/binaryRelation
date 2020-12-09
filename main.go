package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// l1,l2 - коефіцієнти пропорційності швидкості зростання зброї до кількості зброї у супротивника
// b1,b2 - коефіцієнти, що характеризують швидкість старіння власної зброї,
// y1,y2 - функції, що описують рівень взаємної недовіри конкурентів, який вважається незалежним від кількості озброєнь
// m10,m20 - об'єми озброєнь
func armRace(l1, l2, b1, b2, y1, y2, m10, m20, t, n float64) ([]float64, []float64) {
	k := int(n / t)
	m1 := []float64{}
	m2 := []float64{}

	for i := 0; i < k; i++ {
		if i == 0 {
			m1 = append(m1, m10)
			m2 = append(m2, m20)
		} else {
			m1 = append(m1, (t*(l1*m2[i-1]+y1)+m1[i-1])/(1-t*b1))
			m2 = append(m2, (t*(l2*m1[i-1]+y2)+m2[i-1])/(1-t*b2))
		}
	}
	return m1, m2
}

func main() {

	M1, M2 := armRace(0.75, 2.2, 0.3, 0.16, 0.02, 1.2, 233, 340, 0.5, 5)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	m1 := make(plotter.XYs, len(M1))
	m2 := make(plotter.XYs, len(M2))

	for i := 0; i < len(M1); i++ {
		m1[i].X = float64(i)
		m1[i].Y = M1[i]
	}

	for i := 0; i < len(M2); i++ {
		m2[i].X = float64(i)
		m2[i].Y = M2[i]
	}

	p.Title.Text = "Arm race"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"m1", m1,
		"m2", m2)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "arm_race.png"); err != nil {
		panic(err)
	}
}
