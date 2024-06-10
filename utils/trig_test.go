package utils_test

import (
	"math"
	"testing"

	"github.com/SnareChops/nengine/utils"
	"github.com/stretchr/testify/assert"
)

func TestDistanceBetweenPoints(t *testing.T) {
	result := utils.DistanceBetweenPoints(2, 3, 2, 4)
	assert.Equal(t, 1., result)

	result = utils.DistanceBetweenPoints(3, 4, 2, 4)
	assert.Equal(t, 1., result)

	result = utils.DistanceBetweenPoints(3, 3, 2, 4)
	assert.Equal(t, 1.4142135623730951, result)
}

func TestAngleBetweenPoints(t *testing.T) {
	result := utils.AngleBetweenPoints(2, 3, 3, 3)
	assert.Equal(t, 0., result)

	result = utils.AngleBetweenPoints(2, 3, 2, 4)
	assert.Equal(t, math.Pi/2, result)

	result = utils.AngleBetweenPoints(2, 3, 1, 3)
	assert.Equal(t, math.Pi, result)

	result = utils.AngleBetweenPoints(2, 3, 2, 2)
	assert.Equal(t, 3*math.Pi/2, result)

	result = utils.AngleBetweenPoints(2, 3, 3, 4)
	assert.Equal(t, math.Pi/4, result)

	// Distance greater than 1 tests
	result = utils.AngleBetweenPoints(2, 3, 4, 3)
	assert.Equal(t, 0., result)

	result = utils.AngleBetweenPoints(2, 3, 2, 5)
	assert.Equal(t, math.Pi/2, result)

	result = utils.AngleBetweenPoints(2, 3, 2, 0)
	assert.Equal(t, 3*math.Pi/2, result)

	result = utils.AngleBetweenPoints(2, 3, 0, 3)
	assert.Equal(t, math.Pi, result)

	result = utils.AngleBetweenPoints(600, 900, 600, 800)
	assert.Equal(t, 3*math.Pi/2, result)

	result = utils.AngleBetweenPoints(256, 384, 266, 384)
	assert.Equal(t, 0., result)
}

func TestPointAtAngleWithLength(t *testing.T) {
	angle := utils.AngleBetweenPoints(2, 3, 2, 4)
	length := utils.DistanceBetweenPoints(2, 3, 2, 4)
	x, y := utils.PointAtAngleWithDistance(2, 3, angle, length)
	assert.Equal(t, 2., x)
	assert.Equal(t, 4., y)

	angle = utils.AngleBetweenPoints(2, 3, 1, 2)
	length = utils.DistanceBetweenPoints(2, 3, 1, 2)
	x, y = utils.PointAtAngleWithDistance(2, 3, angle, length)
	assert.Equal(t, 1., math.Round(x))
	assert.Equal(t, 2., math.Round(y))
}
