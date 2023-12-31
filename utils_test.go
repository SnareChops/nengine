package nengine_test

import (
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/stretchr/testify/assert"
)

func TestFloats(t *testing.T) {
	a, b := 1, 2
	fa, fb := nengine.Floats(a, b)
	assert.Equal(t, float64(a), fa)
	assert.Equal(t, float64(b), fb)
}

func TestInts(t *testing.T) {
	a, b := 1.5, 2.5
	ia, ib := nengine.Ints(a, b)
	assert.Equal(t, int(a), ia)
	assert.Equal(t, int(b), ib)
}

func TestIsSet(t *testing.T) {
	mask := 0b1010
	state := 0b1000
	assert.True(t, nengine.IsSet(mask, state))
	state = 0b0100
	assert.False(t, nengine.IsSet(mask, state))
}

func TestScreenToRelativePosition(t *testing.T) {
	x, y := 10.0, 20.0
	bounds := new(nengine.RawBounds).Init(10, 20)
	bounds.SetPos2(5, 10)
	rx, ry := nengine.RelativePosition(x, y, bounds)
	assert.Equal(t, 5.0, rx)
	assert.Equal(t, 5.0, ry)
}

func TestRelativePosition(t *testing.T) {
	bounds := new(nengine.RawBounds).Init(10, 20)
	bounds.SetPos2(5, 10)

	rx, ry := nengine.RelativePosition(12, 13, bounds)
	assert.Equal(t, 7., rx)
	assert.Equal(t, 3., ry)
}
