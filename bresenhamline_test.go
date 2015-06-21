package bresenhamline

import (
	"image"
	"sort"
	"testing"
)

func TestLine(t *testing.T) {
	start, end := image.Pt(1, 1), image.Pt(11, 5)
	var linePoints []image.Point
	Line(start, end, func(p image.Point) {
		linePoints = append(linePoints, p)
	})

	expect := []image.Point{
		image.Pt(1, 1), image.Pt(2, 1), image.Pt(3, 2), image.Pt(4, 2),
		image.Pt(5, 3), image.Pt(6, 3), image.Pt(7, 3), image.Pt(8, 4),
		image.Pt(9, 4), image.Pt(10, 5), image.Pt(11, 5),
	}

	if !pointsEq(linePoints, expect) {
		t.Errorf("got %v, want %v", linePoints, expect)
	}

	if start.X != 1 || start.Y != 1 || end.X != 11 || end.Y != 5 {
		t.Errorf("Line function has side effects!")
	}
}

func TestLineDirection(t *testing.T) {

	t.Log("++")
	func() {
		start, end := image.Pt(0, 0), image.Pt(1, 1)
		linePoints := LinePoints(start, end)

		expect := []image.Point{image.Pt(0, 0), image.Pt(1, 1)}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()

	t.Log("+-")
	func() {
		start, end := image.Pt(0, 0), image.Pt(1, -1)
		linePoints := LinePoints(start, end)

		expect := []image.Point{image.Pt(0, 0), image.Pt(1, -1)}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()

	t.Log("-+")
	func() {
		start, end := image.Pt(0, 0), image.Pt(-1, 1)
		linePoints := LinePoints(start, end)

		expect := []image.Point{image.Pt(0, 0), image.Pt(-1, 1)}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()

	t.Log("--")
	func() {
		start, end := image.Pt(0, 0), image.Pt(-1, -1)
		linePoints := LinePoints(start, end)

		expect := []image.Point{image.Pt(0, 0), image.Pt(-1, -1)}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()
}

func TestLineWithSamePoint(t *testing.T) {

	t.Log("same x & y")
	func() {
		start, end := image.Pt(10, 10), image.Pt(10, 10)
		linePoints := LinePoints(start, end)

		expect := []image.Point{image.Pt(10, 10)}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()

	t.Log("same x")
	func() {
		start, end := image.Pt(10, 10), image.Pt(10, 15)
		linePoints := LinePoints(start, end)

		expect := []image.Point{
			image.Pt(10, 10), image.Pt(10, 11), image.Pt(10, 12),
			image.Pt(10, 13), image.Pt(10, 14), image.Pt(10, 15),
		}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()

	t.Log("same y")
	func() {
		start, end := image.Pt(10, 10), image.Pt(15, 10)
		linePoints := LinePoints(start, end)

		expect := []image.Point{
			image.Pt(10, 10), image.Pt(11, 10), image.Pt(12, 10),
			image.Pt(13, 10), image.Pt(14, 10), image.Pt(15, 10),
		}
		if !pointsEq(linePoints, expect) {
			t.Errorf("got %v, want %v", linePoints, expect)
		}
	}()
}

// util methods

type points []image.Point // for sorting

func (p points) Len() int {
	return len(p)
}
func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p points) Less(i, j int) bool {
	xdiff := p[i].X - p[j].X
	if xdiff < 0 {
		return true
	} else if xdiff > 0 {
		return false
	} else {
		return p[i].Y < p[j].Y
	}
}

func pointsEq(a, b []image.Point) bool {
	if len(a) != len(b) {
		return false
	}

	var as, bs points = a, b
	sort.Sort(as)
	sort.Sort(bs)

	for i := range as {
		if !(as[i].Eq(bs[i])) {
			return false
		}
	}
	return true
}
