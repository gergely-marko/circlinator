package fit

import (
	. "circlinator/geometry"

	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOnCircle(t *testing.T) {
	source := CreateCircle(200, 100, 50)

	noise := 2.0

	points := generatePoints(source, 0, 360, 100, noise)

	result := Fit(points)

	assert.InDelta(t, Distance(source.Center, result.Center), 0.0, 1.0, "Distance of centers")
	assert.InDelta(t, source.Radius, result.Radius, 3.0, "Radii of circles")
}

func TestOnArc(t *testing.T) {
	source := CreateCircle(200, 100, 50)

	noise := 2.0

	points := generatePoints(source, 0, 45, 100, noise)

	result := Fit(points)

	assert.InDelta(t, Distance(source.Center, result.Center), 0.0, 1.0, "Distance of centers")
	assert.InDelta(t, source.Radius, result.Radius, 3.0, "Radii of circles")
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

		result = append(result, Point{X: x, Y: y})
	}

	return result
}
