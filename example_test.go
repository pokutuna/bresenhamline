package bresenhamline_test

import (
	"fmt"
	"image"

	"github.com/pokutuna/bresenhamline"
)

func ExampleLine() {
	a := image.Pt(0, 0)
	b := image.Pt(3, 2)
	bresenhamline.Line(a, b, func(p image.Point) {
		fmt.Println(p)
	})
	// Output:
	// (0,0)
	// (1,1)
	// (2,1)
	// (3,2)
}

func ExampleLinePoints() {
	a := image.Pt(0, 0)
	b := image.Pt(3, 2)
	fmt.Println(bresenhamline.LinePoints(a, b))
	// Output:
	// [(0,0) (1,1) (2,1) (3,2)]
}
