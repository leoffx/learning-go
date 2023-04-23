package main

import (
	"math"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type HeartGlyph struct{}

func (HeartGlyph) DrawGlyph(c *draw.Canvas, sty draw.GlyphStyle, pt vg.Point) {
	c.SetColor(sty.Color)
	c.SetLineWidth(sty.Radius / 2)

	// Scale the radius by a factor of 0.75 to make the heart shape more proportional
	r := sty.Radius * 0.75

	// Translate to the desired point
	c.Push()
	c.Translate(pt)

	// Calculate the heart shape points using the applyHeartEquation function
	points := make([]vg.Point, 0)
	t := 0.0
	for t <= 2*math.Pi {
		x, y := applyHeartEquation(t)
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
