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

func TestCombineRectsExample(t *testing.T) {
	rects := []nengine.Rect{
		{0, 3520, 64, 3584},
		{64, 3520, 128, 3584},
		{128, 3520, 192, 3584},
		{192, 3520, 256, 3584},
		{256, 3520, 320, 3584},
		{320, 3520, 384, 3584},
		{0, 3584, 64, 3648},
		{64, 3584, 128, 3648},
		{128, 3584, 192, 3648},
		{192, 3584, 256, 3648},
		{256, 3584, 320, 3648},
		{320, 3584, 384, 3648},
	}

	result := nengine.CombineRects(rects)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, 0, result[0].MinX)
	assert.Equal(t, 3520, result[0].MinY)
	assert.Equal(t, 384, result[0].MaxX)
	assert.Equal(t, 3648, result[0].MaxY)
}
