package main

import (
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type FunctionGlyph struct {
	points plotter.XYs
}

func (s FunctionGlyph) DrawGlyph(c *draw.Canvas, sty draw.GlyphStyle, pt vg.Point) {
	c.SetColor(sty.Color)
	c.SetLineWidth(sty.Radius / 2)

	// Translate to the desired point
	c.Push()
	c.Translate(pt)

	// Construct a path using the calculated points
	p := vg.Path{}
	p.Move(convertXYToVGPoint(s.points[0], sty.Radius))
	for i := 1; i < len(s.points); i++ {
		p.Line(convertXYToVGPoint(s.points[i], sty.Radius))
	}
	p.Close()

	c.Fill(p)
	c.Pop()
}

func convertXYToVGPoint(point plotter.XY, radius vg.Length) vg.Point {
	return vg.Point{X: vg.Length(point.X) * radius, Y: vg.Length(point.Y) * radius}
}
