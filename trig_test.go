package nengine_test

import (
	"math"
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/stretchr/testify/assert"
)

func TestDistanceBetweenPoints(t *testing.T) {
	result := nengine.DistanceBetweenPoints(2, 3, 2, 4)
	assert.Equal(t, 1., result)

	result = nengine.DistanceBetweenPoints(3, 4, 2, 4)
	assert.Equal(t, 1., result)

	result = nengine.DistanceBetweenPoints(3, 3, 2, 4)
	assert.Equal(t, 1.4142135623730951, result)
}

func TestAngleBetweenPoints(t *testing.T) {
	result := nengine.AngleBetweenPoints(2, 3, 3, 3)
	assert.Equal(t, 0., result)

	result = nengine.AngleBetweenPoints(2, 3, 2, 4)
	assert.Equal(t, math.Pi/2, result)

	result = nengine.AngleBetweenPoints(2, 3, 1, 3)
	assert.Equal(t, math.Pi, result)

	result = nengine.AngleBetweenPoints(2, 3, 2, 2)
	assert.Equal(t, 3*math.Pi/2, result)

	result = nengine.AngleBetweenPoints(2, 3, 3, 4)
	assert.Equal(t, math.Pi/4, result)

	// Distance greater than 1 tests
	result = nengine.AngleBetweenPoints(2, 3, 4, 3)
	assert.Equal(t, 0., result)

	result = nengine.AngleBetweenPoints(2, 3, 2, 5)
	assert.Equal(t, math.Pi/2, result)

	result = nengine.AngleBetweenPoints(2, 3, 2, 0)
	assert.Equal(t, 3*math.Pi/2, result)

	result = nengine.AngleBetweenPoints(2, 3, 0, 3)
	assert.Equal(t, math.Pi, result)

	result = nengine.AngleBetweenPoints(600, 900, 600, 800)
	assert.Equal(t, 3*math.Pi/2, result)

	result = nengine.AngleBetweenPoints(256, 384, 266, 384)
	assert.Equal(t, 0., result)
}

func TestPointAtAngleWithLength(t *testing.T) {
	angle := nengine.AngleBetweenPoints(2, 3, 2, 4)
	length := nengine.DistanceBetweenPoints(2, 3, 2, 4)
	x, y := nengine.PointAtAngleWithDistance(2, 3, angle, length)
	assert.Equal(t, 2., x)
	assert.Equal(t, 4., y)

	angle = nengine.AngleBetweenPoints(2, 3, 1, 2)
	length = nengine.DistanceBetweenPoints(2, 3, 1, 2)
	x, y = nengine.PointAtAngleWithDistance(2, 3, angle, length)
	assert.Equal(t, 1., math.Round(x))
	assert.Equal(t, 2., math.Round(y))
}
