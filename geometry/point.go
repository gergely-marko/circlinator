package geometry

import "math"

type Point struct {
	X float64
	Y float64
}

func Distance(p1 Point, p2 Point) float64 {
	dX := p2.X - p1.X
	dY := p2.Y - p1.Y

	return math.Sqrt(dX*dX + dY*dY)
}
