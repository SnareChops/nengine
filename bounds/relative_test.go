package bounds_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/stretchr/testify/assert"
)

func TestRelativeBoundsAnchor(t *testing.T) {
	parent := new(bounds.Raw).Init(10, 10)
	result := new(bounds.Relative).Init(parent, 7, 7)

	result.SetAnchor(bounds.CENTER, bounds.CENTER)

	x, y := result.Min()
	assert.Equal(t, -3.5, x)
	assert.Equal(t, -3.5, y)
}

func TestRelativeBoundsPosition(t *testing.T) {
	parent := new(bounds.Raw).Init(10, 10)
	parent.SetPos2(5, 5)
	result := new(bounds.Relative).Init(parent, 8, 8)
	result.SetPos2(2, 3)

	x, y := result.Min()
	assert.Equal(t, 7., x)
	assert.Equal(t, 8., y)
}

func TestRelativeBoundsBoxInBox(t *testing.T) {
	parent := new(bounds.Raw).Init(10, 12)
	parent.SetAnchor(bounds.CENTER, bounds.BOTTOM)
	parent.SetPos2(20, 20)

	box := new(bounds.Relative).Init(parent, 5, 8)
	box.SetAnchor(bounds.CENTER, bounds.BOTTOM)
	box.SetPos2(-1, -2)

	x, y := box.Min()
	assert.Equal(t, 16.5, x)
	assert.Equal(t, 11., y)
}

func TestRelativeMin(t *testing.T) {
	parent := new(bounds.Raw).Init(10, 12)
	parent.SetPos2(20, 20)

	box := new(bounds.Relative).Init(parent, 2, 4)
	box.SetPos2(-5, -3)

	x, y := box.Min()
	assert.Equal(t, 15., x)
	assert.Equal(t, 17., y)
	assert.Equal(t, 15., box.MinX())
	assert.Equal(t, 17., box.MinY())

	// With offsets
	parent.SetAnchor(bounds.LEFT, bounds.TOP)
	box.SetAnchor(bounds.RIGHT, bounds.BOTTOM)

	x, y = box.Min()
	assert.Equal(t, 14., x)
	assert.Equal(t, 14., y)
}
