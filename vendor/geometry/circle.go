package geometry

type Circle struct {
	Center Point
	Radius float64
}

func CreateCircle(x float64, y float64, r float64) Circle {
	return Circle{
		Center: Point{X: x, Y: y},
		Radius: r}
}
