package aco

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func (a *Aco) PlotAnt(ant *Ant, outputFileName string) {
	p := plot.New()
	p.Title.Text = "Ant path"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	citiesXYs := citiesToXys(a.enviroment.cities)
	pathXYs := pathToXys(ant.Path, a.enviroment.cities)
	s, _ := plotter.NewScatter(citiesXYs)
	l, _, _ := plotter.NewLinePoints(pathXYs)

	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(s)
	p.Add(l)
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprintf("%v.png", outputFileName)); err != nil {
		panic(err)
	}
}

func pathToXys(path []int, cities *[]City) plotter.XYs {
	pts := make(plotter.XYs, len(path)+1)

	for i := range path {
		pts[i].X = (*cities)[path[i]].X
		pts[i].Y = (*cities)[path[i]].Y
	}

	pts[len(pts)-1].X = pts[0].X
	pts[len(pts)-1].Y = pts[0].Y

	return pts
}

func citiesToXys(cities *[]City) plotter.XYs {
	pts := make(plotter.XYs, len(*cities))

	for i := range pts {
		pts[i].X = (*cities)[i].X
		pts[i].Y = (*cities)[i].Y
	}

	return pts
}
