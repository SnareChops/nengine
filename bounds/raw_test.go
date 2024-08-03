package bounds_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/stretchr/testify/assert"
)

func TestBoundsOffset(t *testing.T) {
	b := new(bounds.Raw).Init(2, 4)
	b.SetAnchor(bounds.CENTER, bounds.CENTER)
	x, y := b.Offset()
	assert.Equal(t, 1., x)
	assert.Equal(t, 2., y)
}

func TestBoundsAnchor(t *testing.T) {
	result := new(bounds.Raw).Init(2, 3)
	result.SetAnchor(bounds.LEFT, bounds.TOP)

	ox, oy := result.Offset()
	assert.Equal(t, 0., ox)
	assert.Equal(t, 0., oy)

	result.SetAnchor(bounds.CENTER, bounds.CENTER)
	ox, oy = result.Offset()
	assert.Equal(t, 0.5, ox)
	assert.Equal(t, 1., oy)

	result.SetAnchor(bounds.RIGHT, bounds.BOTTOM)
	ox, oy = result.Offset()
	assert.Equal(t, 1., ox)
	assert.Equal(t, 2., oy)
}

// func TestPosOf(t *testing.T) {
// 	result := new(bounds.Raw).Init(10, 10)

// 	x, y := result.PosOf(bounds.LEFT, bounds.TOP)
// 	assert.Equal(t, 0., x)
// 	assert.Equal(t, 0., y)

// 	x, y = result.PosOf(bounds.CENTER, bounds.TOP)
// 	assert.Equal(t, 5., x)
// 	assert.Equal(t, 0., y)

// 	x, y = result.PosOf(bounds.RIGHT, bounds.TOP)
// 	assert.Equal(t, 10., x)
// 	assert.Equal(t, 0., y)

// 	x, y = result.PosOf(bounds.LEFT, bounds.CENTER)
// 	assert.Equal(t, 0., x)
// 	assert.Equal(t, 5., y)

// 	x, y = result.PosOf(bounds.CENTER, bounds.CENTER)
// 	assert.Equal(t, 5., x)
// 	assert.Equal(t, 5., y)

// 	x, y = result.PosOf(bounds.RIGHT, bounds.CENTER)
// 	assert.Equal(t, 10., x)
// 	assert.Equal(t, 5., y)

// 	x, y = result.PosOf(bounds.LEFT, bounds.BOTTOM)
// 	assert.Equal(t, 0., x)
// 	assert.Equal(t, 10., y)

// 	x, y = result.PosOf(bounds.CENTER, bounds.BOTTOM)
// 	assert.Equal(t, 5., x)
// 	assert.Equal(t, 10., y)

// 	x, y = result.PosOf(bounds.RIGHT, bounds.BOTTOM)
// 	assert.Equal(t, 10., x)
// 	assert.Equal(t, 10., y)
// }

func TestRawResize(t *testing.T) {
	// Create bounds with an initial size
	b := new(bounds.Raw).Init(10, 20)
	b.SetPos2(100, 100)
	b.SetAnchor(bounds.CENTER, bounds.CENTER)
	assert.Equal(t, 10, b.Dx())
	assert.Equal(t, 20, b.Dy())
	ox, oy := b.Offset()
	assert.Equal(t, 5., ox)
	assert.Equal(t, 10., oy)
	w, h := b.Size()
	assert.Equal(t, 10, w)
	assert.Equal(t, 20, h)
	x, y := b.Min()
	assert.Equal(t, 95., x)
	assert.Equal(t, 90., y)
	x, y = b.Mid()
	assert.Equal(t, 100., x)
	assert.Equal(t, 100., y)
	x, y = b.Max()
	assert.Equal(t, 104., x)
	assert.Equal(t, 109., y)
	x = b.MaxX()
	assert.Equal(t, 104., x)
	y = b.MaxY()
	assert.Equal(t, 109., y)

	// Resize the bounds and verify all values are correct
	b.Resize(20, 30)
	assert.Equal(t, 20, b.Dx())
	assert.Equal(t, 30, b.Dy())
	ox, oy = b.Offset()
	assert.Equal(t, 10., ox)
	assert.Equal(t, 15., oy)
	w, h = b.Size()
	assert.Equal(t, 20, w)
	assert.Equal(t, 30, h)
	x, y = b.Min()
	assert.Equal(t, 90., x)
	assert.Equal(t, 85., y)
	x, y = b.Mid()
	assert.Equal(t, 100., x)
	assert.Equal(t, 100., y)
	x, y = b.Max()
	assert.Equal(t, 109., x)
	assert.Equal(t, 114., y)
	x = b.MaxX()
	assert.Equal(t, 109., x)
	y = b.MaxY()
	assert.Equal(t, 114., y)
}
