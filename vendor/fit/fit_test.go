package fit

import (
	"geometry"
	. "geometry"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestOnCircle(t *testing.T) {
	source := CreateCircle(200, 100, 50)

	noise := 2.0

	points := generatePoints(source, 0, 360, 100, noise)

	result := Fit(points)

	assertDistance(source.Center, result.Center, 3.0, t)
	assertAbsFloat(source.Radius, result.Radius, 6.0, t)
}

func TestOnArc(t *testing.T) {
	source := CreateCircle(200, 100, 50)

	noise := 2.0

	points := generatePoints(source, 0, 45, 100, noise)

	result := Fit(points)

	assertDistance(source.Center, result.Center, 3.0, t)
	assertAbsFloat(source.Radius, result.Radius, 6.0, t)
}

func assertDistance(expected Point, actual Point, threshold float64, t *testing.T) {
	assertAbsFloat(0.0, Distance(expected, actual), threshold, t)
}

func assertAbsFloat(expected float64, actual float64, threshold float64, t *testing.T) {
	delta := math.Abs(actual - expected)

	if delta > threshold {
		t.Errorf("Delta (%0.2f) is greater than threshold (%0.2f)", delta, threshold)
	}
}

func generatePoints(circle Circle, startAngle float64, endAngle float64, count int, noise float64) []Point {
	var result []Point

	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	for i := 0; i < count; i++ {
		radius := rnd.NormFloat64()*noise + circle.Radius

		alpha := startAngle + rnd.Float64()*(endAngle-startAngle)
		alpha *= math.Pi / 180.0

		x := circle.Center.X + radius*math.Cos(alpha)
		y := circle.Center.Y + radius*math.Sin(alpha)

		result = append(result, geometry.Point{X: x, Y: y})
	}

	return result
}
