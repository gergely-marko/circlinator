package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	p1 := Point{X: 2.0, Y: 2.0}
	p2 := Point{X: 5.0, Y: 6.0}
	dist := Distance(p1, p2)

	assert.InDelta(t, 5.0, dist, 0.00001)
}
