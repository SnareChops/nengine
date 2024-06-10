package nengine_test

import (
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/SnareChops/nengine/bit"
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
	assert.True(t, bit.IsSet(mask, state))
	state = 0b0100
	assert.False(t, bit.IsSet(mask, state))
}

func TestScreenToRelativePosition(t *testing.T) {
	bounds := new(nengine.RawBounds).Init(10, 20)
	bounds.SetPos2(5, 10)
	rx, ry := nengine.RelativePosition(10.0, 20.0, bounds)
	assert.Equal(t, 5.0, rx)
	assert.Equal(t, 10.0, ry)
}

func TestRelativePosition(t *testing.T) {
	bounds := new(nengine.RawBounds).Init(10, 20)
	bounds.SetPos2(5, 10)

	rx, ry := nengine.RelativePosition(12, 13, bounds)
	assert.Equal(t, 7, rx)
	assert.Equal(t, 3, ry)
}

func TestGridPointsAroundBounds(t *testing.T) {
	bounds := new(nengine.RawBounds).Init(40, 20)
	bounds.SetPos2(30, 30)
	result := nengine.GridPointsAroundBounds(bounds, 10, 10)

	assert.Equal(t, 16, len(result))
	// corner
	assert.Equal(t, 25., result[0].X())
	assert.Equal(t, 25., result[0].Y())
	//
	assert.Equal(t, 35., result[1].X())
	assert.Equal(t, 25., result[1].Y())
	//
	assert.Equal(t, 45., result[2].X())
	assert.Equal(t, 25., result[2].Y())
	//
	assert.Equal(t, 55., result[3].X())
	assert.Equal(t, 25., result[3].Y())
	//
	assert.Equal(t, 65., result[4].X())
	assert.Equal(t, 25., result[4].Y())
	// corner
	assert.Equal(t, 75., result[5].X())
	assert.Equal(t, 25., result[5].Y())
	//
	assert.Equal(t, 75., result[6].X())
	assert.Equal(t, 35., result[6].Y())
	//
	assert.Equal(t, 75., result[7].X())
	assert.Equal(t, 45., result[7].Y())
	// corner
	assert.Equal(t, 75., result[8].X())
	assert.Equal(t, 55., result[8].Y())
	//
	assert.Equal(t, 65., result[9].X())
	assert.Equal(t, 55., result[9].Y())
	//
	assert.Equal(t, 55., result[10].X())
	assert.Equal(t, 55., result[10].Y())
	//
	assert.Equal(t, 45., result[11].X())
	assert.Equal(t, 55., result[11].Y())
	//
	assert.Equal(t, 35., result[12].X())
	assert.Equal(t, 55., result[12].Y())
	// corner
	assert.Equal(t, 25., result[13].X())
	assert.Equal(t, 55., result[13].Y())
	//
	assert.Equal(t, 25., result[14].X())
	assert.Equal(t, 45., result[14].Y())
	//
	assert.Equal(t, 25., result[15].X())
	assert.Equal(t, 35., result[15].Y())
}
