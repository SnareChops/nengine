package nengine_test

import (
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/stretchr/testify/assert"
)

func TestIndexAt(t *testing.T) {
	grid := new(nengine.SpriteGrid).Init(100, 100, 10, 10)
	result := grid.IndexAt(0, 0)
	assert.Equal(t, 0, result)

	result = grid.IndexAt(10, 10)
	assert.Equal(t, 11, result)

	result = grid.IndexAt(25, 25)
	assert.Equal(t, 22, result)
}
