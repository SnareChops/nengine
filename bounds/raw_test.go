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

func TestBoundsScale(t *testing.T) {
	result := new(bounds.Raw).Init(200, 300)
	result.SetScale(2.3, 4.5)
	x, y := result.Scale()
	assert.Equal(t, 2.3, x)
	assert.Equal(t, 4.5, y)
}

func TestBoundsRotation(t *testing.T) {
	result := new(bounds.Raw).Init(200, 300)
	result.SetRotation(2.3)
	r := result.Rotation()
	assert.Equal(t, 2.3, r)
}

func TestBoundsAnchor(t *testing.T) {
	result := new(bounds.Raw).Init(2, 3)
	result.SetAnchor(bounds.LEFT, bounds.TOP)

	ax, ay := result.Anchor()
	ox, oy := result.Offset()
	assert.Equal(t, bounds.LEFT, ax)
	assert.Equal(t, bounds.TOP, ay)
	assert.Equal(t, 0., ox)
	assert.Equal(t, 0., oy)

	result.SetAnchor(bounds.CENTER, bounds.CENTER)
	ax, ay = result.Anchor()
	ox, oy = result.Offset()
	assert.Equal(t, bounds.CENTER, ax)
	assert.Equal(t, bounds.CENTER, ay)
	assert.Equal(t, 1., ox)
	assert.Equal(t, 1.5, oy)

	result.SetAnchor(bounds.RIGHT, bounds.BOTTOM)
	ax, ay = result.Anchor()
	ox, oy = result.Offset()
	assert.Equal(t, bounds.RIGHT, ax)
	assert.Equal(t, bounds.BOTTOM, ay)
	assert.Equal(t, 2., ox)
	assert.Equal(t, 3., oy)
}

func TestBoundsSize(t *testing.T) {
	result := new(bounds.Raw).Init(2, 3)
	w, h := result.Size()
	assert.Equal(t, 2, w)
	assert.Equal(t, 3, h)

	result = new(bounds.Raw).Init(2, 3)
	result.SetScale(2, 2)
	w, h = result.Size()
	assert.Equal(t, 4, w)
	assert.Equal(t, 6, h)
}

func TestBoundsIsWithin(t *testing.T) {
	bounds := new(bounds.Raw).Init(2, 3)
	result := bounds.IsWithin(1, 2)
	assert.True(t, result)
}

func TestBoundsDoesCollide(t *testing.T) {
	bounds1 := new(bounds.Raw).Init(4, 4)
	bounds1.SetPos2(8, 8)
	bounds2 := new(bounds.Raw).Init(4, 4)
	bounds2.SetPos2(20, 20)
	assert.False(t, bounds1.DoesCollide(bounds2))

	bounds1 = new(bounds.Raw).Init(4, 4)
	bounds2 = new(bounds.Raw).Init(4, 4)
	bounds1.SetPos2(8, 8)
	bounds2.SetPos2(6, 4)
	assert.True(t, bounds1.DoesCollide(bounds2))

	bounds1 = new(bounds.Raw).Init(4, 4)
	bounds2 = new(bounds.Raw).Init(4, 4)
	bounds1.SetPos2(8, 8)
	bounds2.SetPos2(4, 6)
	assert.True(t, bounds1.DoesCollide(bounds2))

	bounds1 = new(bounds.Raw).Init(4, 4)
	bounds2 = new(bounds.Raw).Init(4, 4)
	bounds1.SetPos2(8, 8)
	bounds2.SetPos2(10, 4)
	assert.True(t, bounds1.DoesCollide(bounds2))

	bounds1 = new(bounds.Raw).Init(4, 4)
	bounds2 = new(bounds.Raw).Init(4, 4)
	bounds1.SetPos2(8, 8)
	bounds2.SetPos2(4, 10)
	assert.True(t, bounds1.DoesCollide(bounds2))

	bounds1 = new(bounds.Raw).Init(4, 4)
	bounds2 = new(bounds.Raw).Init(6, 6)
	bounds1.SetPos2(8, 8)
	bounds2.SetPos2(7, 7)
	assert.True(t, bounds1.DoesCollide(bounds2))
}

func TestPosOf(t *testing.T) {
	result := new(bounds.Raw).Init(10, 10)

	x, y := result.PosOf(bounds.LEFT, bounds.TOP)
	assert.Equal(t, 0., x)
	assert.Equal(t, 0., y)

	x, y = result.PosOf(bounds.CENTER, bounds.TOP)
	assert.Equal(t, 5., x)
	assert.Equal(t, 0., y)

	x, y = result.PosOf(bounds.RIGHT, bounds.TOP)
	assert.Equal(t, 10., x)
	assert.Equal(t, 0., y)

	x, y = result.PosOf(bounds.LEFT, bounds.CENTER)
	assert.Equal(t, 0., x)
	assert.Equal(t, 5., y)

	x, y = result.PosOf(bounds.CENTER, bounds.CENTER)
	assert.Equal(t, 5., x)
	assert.Equal(t, 5., y)

	x, y = result.PosOf(bounds.RIGHT, bounds.CENTER)
	assert.Equal(t, 10., x)
	assert.Equal(t, 5., y)

	x, y = result.PosOf(bounds.LEFT, bounds.BOTTOM)
	assert.Equal(t, 0., x)
	assert.Equal(t, 10., y)

	x, y = result.PosOf(bounds.CENTER, bounds.BOTTOM)
	assert.Equal(t, 5., x)
	assert.Equal(t, 10., y)

	x, y = result.PosOf(bounds.RIGHT, bounds.BOTTOM)
	assert.Equal(t, 10., x)
	assert.Equal(t, 10., y)
}

func TestRelativeBoundsAnchor(t *testing.T) {
	parent := new(bounds.Raw).Init(10, 10)
	result := new(bounds.Relative).Init(parent, 8, 8)

	result.SetAnchor(bounds.CENTER, bounds.CENTER)

	x, y := result.RawPos()
	assert.Equal(t, -4., x)
	assert.Equal(t, -4., y)
}

func TestScaleTo(t *testing.T) {
	bounds := new(bounds.Raw).Init(10, 10)
	bounds.ScaleTo(20, 20)
	x, y := bounds.Scale()
	assert.Equal(t, 2., x)
	assert.Equal(t, 2., y)

	bounds.ScaleTo(30, 30)
	x, y = bounds.Scale()
	assert.Equal(t, 3., x)
	assert.Equal(t, 3., y)

	bounds.ScaleTo(20, 10)
	x, y = bounds.Scale()
	assert.Equal(t, 1., x)
	assert.Equal(t, 1., y)
}
