package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	plot := plot.New()
	plot.Title.Text = "What I think about Go"

	points := createRange(0, 2*math.Pi, 0.1)
	functionPoints := applyFunctionToPoints(points)

	scatter := &plotter.Scatter{
		XYs: functionPoints,
		GlyphStyle: draw.GlyphStyle{
			Color:  color.RGBA{R: 255, A: 255},
			Radius: vg.Points(0.3),
			Shape:  HeartGlyph{},
		},
	}

	plot.Add(scatter)

	if err := plot.Save(10*vg.Centimeter, 10*vg.Centimeter, "heart.png"); err != nil {
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

func applyFunctionToPoints(points []float64) plotter.XYs {
	functionPoints := make(plotter.XYs, len(points))
	for i, point := range points {
		x, y := applyHeartFunction(point)
		functionPoints[i].X = x
		functionPoints[i].Y = y
	}
	return functionPoints
}

func applyHeartFunction(point float64) (float64, float64) {
	x := 16 * math.Pow(math.Sin(point), 3)
	y := 13*math.Cos(point) - 5*math.Cos(2*point) - 2*math.Cos(3*point) - math.Cos(4*point)
	return x, y
}

type HeartGlyph struct{}

func (HeartGlyph) DrawGlyph(c *draw.Canvas, sty draw.GlyphStyle, pt vg.Point) {
	c.SetColor(sty.Color)
	c.SetLineWidth(sty.Radius / 2)

	// Scale the radius by a factor of 0.75 to make the heart shape more proportional
	r := sty.Radius * 0.75

	// Translate to the desired point
	c.Push()
	c.Translate(pt)

	// Calculate the heart shape points using the applyHeartFunction function
	points := make([]vg.Point, 0)
	t := 0.0
	for t <= 2*math.Pi {
		x, y := applyHeartFunction(t)
		points = append(points, vg.Point{X: r * vg.Length(x), Y: r * vg.Length(y)})
		t += 0.1
	}

	// Construct a path using the calculated points
	p := vg.Path{}
	p.Move(points[0])
	for i := 1; i < len(points); i++ {
		p.Line(points[i])
	}
	p.Close()

	c.Fill(p)
	c.Pop()
}
