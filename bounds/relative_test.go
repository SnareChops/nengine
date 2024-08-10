package bounds_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/stretchr/testify/assert"
)

func TestRelativeBoundsAnchor(t *testing.T) {
	parent := new(bounds.Raw).Init(10, 10)
	result := new(bounds.Relative).Init(parent, 8, 8)

	result.SetAnchor(bounds.CENTER, bounds.CENTER)

	x, y := result.Min()
	assert.Equal(t, -4., x)
	assert.Equal(t, -4., y)
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
