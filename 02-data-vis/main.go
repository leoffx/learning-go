package main

import (
	"image/color"
	"math"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	points := createRange(0, 2*math.Pi, 0.1)
	heartPoints := applyFunctionToPoints(points, applyHeartEquation)

	createPlot("What I think about Go", color.RGBA{R: 255, A: 255}, 0.25, heartPoints)
}

func createPlot(title string, color color.RGBA, scale float64, functionPoints plotter.XYs) {
	plot := plot.New()
	plot.Title.Text = title

	scatter := &plotter.Scatter{
		XYs: functionPoints,
		GlyphStyle: draw.GlyphStyle{
			Color:  color,
			Radius: vg.Points(scale),
			Shape:  FunctionGlyph{functionPoints},
		},
	}

	plot.Add(scatter)

	var escapedTitle = title
	escapedTitle = strings.ReplaceAll(escapedTitle, " ", "_")
	escapedTitle = strings.ToLower(escapedTitle)
	escapedTitle += ".png"
	if err := plot.Save(10*vg.Centimeter, 10*vg.Centimeter, escapedTitle); err != nil {
		panic(err)
	}
}

func createRange(start, end, step float64) []float64 {
	n := int((end - start) / step)
	result := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = start + float64(i)*step
	}
	return result
}

func applyFunctionToPoints(points []float64, callable func(float64) (float64, float64)) plotter.XYs {
	functionPoints := make(plotter.XYs, len(points))
	for i, point := range points {
		x, y := callable(point)
		functionPoints[i].X = x
		functionPoints[i].Y = y
	}
	return functionPoints
}

func applyHeartEquation(point float64) (float64, float64) {
	x := 16 * math.Pow(math.Sin(point), 3)
	y := 13*math.Cos(point) - 5*math.Cos(2*point) - 2*math.Cos(3*point) - math.Cos(4*point)
	return x, y
}
