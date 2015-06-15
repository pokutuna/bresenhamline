bresenhamline
===

[![Build Status on Travis-CI](https://travis-ci.org/pokutuna/bresenhamline.svg?branch=master)](https://travis-ci.org/pokutuna/bresenhamline)
[![Documentation on godoc.org](https://godoc.org/github.com/pokutuna/bresenhamline?status.svg)](https://godoc.org/github.com/pokutuna/bresenhamline)

Package bresenhamline provides simple line methods based on Bresenham's Line Algorithm.


# Example

```go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/pokutuna/bresenhamline"
)

func main() {
	var palette color.Palette = []color.Color{color.White, color.Black}
	img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette)

	bresenhamline.Line(image.Pt(50, 10), image.Pt(10, 90), func(p image.Point) {
		img.Set(p.X, p.Y, palette[1])
	})
	bresenhamline.Line(image.Pt(50, 10), image.Pt(90, 90), func(p image.Point) {
		img.Set(p.X, p.Y, palette[1])
	})

	out, _ := os.Create("out.png")
	defer out.Close()
	png.Encode(out, img)
	out.Sync()
}
```

![example output](https://raw.githubusercontent.com/wiki/pokutuna/bresenhamline/images/out.png)
