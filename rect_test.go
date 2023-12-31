package nengine_test

import (
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/stretchr/testify/assert"
)

func TestCombineRects(t *testing.T) {
	rects := []nengine.Rect{
		{0, 0, 1, 1},
		{0, 1, 1, 2},
		{1, 0, 2, 1},
		{1, 1, 2, 2},
	}

	result := nengine.CombineRects(rects)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, 0, result[0].MinX)
	assert.Equal(t, 0, result[0].MinY)
	assert.Equal(t, 2, result[0].MaxX)
	assert.Equal(t, 2, result[0].MaxY)

	bounds := nengine.RectsToBounds(result)

	assert.Equal(t, 1, len(bounds))
	w, h := bounds[0].Size()
	x, y := bounds[0].Pos2()
	assert.Equal(t, 2, w)
	assert.Equal(t, 2, h)
	assert.Equal(t, 0., x)
	assert.Equal(t, 0., y)
}
