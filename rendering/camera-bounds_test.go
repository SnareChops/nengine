package rendering_test

import (
	"testing"

	"github.com/SnareChops/nengine/rendering"
	"github.com/stretchr/testify/assert"
)

func TestCameraBounds(t *testing.T) {
	bounds := new(rendering.CameraBounds).Init(100, 80)
	x, y := bounds.Pos()
	assert.Equal(t, 0., x)
	assert.Equal(t, 0., y)
	assert.Equal(t, 0., bounds.X())
	assert.Equal(t, 0., bounds.Y())
	w, h := bounds.Size()
	assert.Equal(t, 100, w)
	assert.Equal(t, 80, h)
	assert.Equal(t, 100, bounds.Dx())
	assert.Equal(t, 80, bounds.Dy())
	x, y = bounds.Min()
	assert.Equal(t, -50., x)
	assert.Equal(t, -40., y)
	assert.Equal(t, -50., bounds.MinX())
	assert.Equal(t, -40., bounds.MinY())
	x, y = bounds.Max()
	assert.Equal(t, 50., x)
	assert.Equal(t, 40., y)
	assert.Equal(t, 50., bounds.MaxX())
	assert.Equal(t, 40., bounds.MaxY())
	assert.True(t, bounds.IsWithin(20, 30))
	assert.False(t, bounds.IsWithin(55, 45))

	bounds.SetPos(100, 100)
	x, y = bounds.Pos()
	assert.Equal(t, 100., x)
	assert.Equal(t, 100., y)
	assert.Equal(t, 100., bounds.X())
	assert.Equal(t, 100., bounds.Y())
	w, h = bounds.Size()
	assert.Equal(t, 100, w)
	assert.Equal(t, 80, h)
	assert.Equal(t, 100, bounds.Dx())
	assert.Equal(t, 80, bounds.Dy())
	x, y = bounds.Min()
	assert.Equal(t, 50., x)
	assert.Equal(t, 60., y)
	assert.Equal(t, 50., bounds.MinX())
	assert.Equal(t, 60., bounds.MinY())
	x, y = bounds.Max()
	assert.Equal(t, 150., x)
	assert.Equal(t, 140., y)
	assert.Equal(t, 150., bounds.MaxX())
	assert.Equal(t, 140., bounds.MaxY())
	assert.True(t, bounds.IsWithin(60, 70))
	assert.False(t, bounds.IsWithin(45, 55))
}
