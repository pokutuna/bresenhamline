// Package bresenhamline provides simple line methods based on Bresenham's Line Algorithm.
package bresenhamline

import (
	"image"
	"math"
)

// Line takes 2 image.Point (start & end) and a callback function given line points.
func Line(a, b image.Point, point func(image.Point)) {
	steep := math.Abs(float64(b.Y-a.Y)) > math.Abs(float64(b.X-a.X))
	if steep {
		a.X, a.Y = a.Y, a.X
		b.X, b.Y = b.Y, b.X
	}
	if a.X > b.X {
		a.X, b.X = b.X, a.X
		a.Y, b.Y = b.Y, a.Y
	}

	dX := b.X - a.X
	dY := int(math.Abs(float64(b.Y - a.Y)))
	er := dX / 2

	y := a.Y
	var yStep int
	if a.Y < b.Y {
		yStep = 1
	} else {
		yStep = -1
	}

	for x := a.X; x <= b.X; x++ {
		var pos image.Point
		if steep {
			pos = image.Pt(y, x)
		} else {
			pos = image.Pt(x, y)
		}

		point(pos)

		er = er - dY
		if er < 0 {
			y = y + yStep
			er = er + dX
		}
	}
}

// LinePoints returns points of line, inclusive start and end points.
func LinePoints(a, b image.Point) []image.Point {
	linePoints := []image.Point{}
	Line(a, b, func(p image.Point) {
		linePoints = append(linePoints, p)
	})
	return linePoints
}
