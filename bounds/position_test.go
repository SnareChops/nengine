package bounds_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/stretchr/testify/assert"
)

func TestGridAlign(t *testing.T) {
	point1 := bounds.Point(0, 0)
	point1.GridAlign(64, 64)
	assert.Equal(t, 0., point1.X())
	assert.Equal(t, 0., point1.Y())

	point := bounds.Point(10, 10)
	point.GridAlign(64, 64)
	assert.Equal(t, 0., point.X())
	assert.Equal(t, 0., point.Y())

	point = bounds.Point(65, 65)
	point.GridAlign(64, 64)
	assert.Equal(t, 64., point.X())
	assert.Equal(t, 64., point.Y())
}
