package fit

import (
	. "geometry"
	"math"
)

// This is a dummy implementation
func Fit(points []Point) Circle {
	xMin := math.MaxFloat64
	xMax := -math.MaxFloat64

	yMin := math.MaxFloat64
	yMax := -math.MaxFloat64

	for _, point := range points {
		xMin = math.Min(xMin, point.X)
		xMax = math.Max(xMax, point.X)

		yMin = math.Min(yMin, point.Y)
		yMax = math.Max(yMax, point.Y)
	}

	x := (xMax + xMin) / 2.0
	y := (yMax + yMin) / 2.0
	r1 := (xMax - xMin) / 2.0
	r2 := (yMax - yMin) / 2.0
	r := (r1 + r2) / 2.0

	return CreateCircle(x, y, r)
}
