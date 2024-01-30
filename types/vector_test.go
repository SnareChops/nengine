package types_test

import (
	"testing"

	"github.com/SnareChops/nengine/types"
	"github.com/stretchr/testify/assert"
)

func TestVectorAdd(t *testing.T) {
	result := types.Vector{0, 0}.Add(types.Vector{0, 0})
	assert.Equal(t, 0., result.X)
	assert.Equal(t, 0., result.Y)

	result = types.Vector{1, 2}.Add(types.Vector{3, 4})
	assert.Equal(t, 4., result.X)
	assert.Equal(t, 6., result.Y)
}

func TestVectorSub(t *testing.T) {
	result := types.Vector{0, 0}.Sub(types.Vector{0, 0})
	assert.Equal(t, 0., result.X)
	assert.Equal(t, 0., result.Y)

	result = types.Vector{3, 4}.Sub(types.Vector{1, 2})
	assert.Equal(t, 2., result.X)
	assert.Equal(t, 2., result.Y)

	result = types.Vector{1, 2}.Sub(types.Vector{3, 4})
	assert.Equal(t, -2., result.X)
	assert.Equal(t, -2., result.Y)
}

func TestVectorNormalize(t *testing.T) {
	result := types.Vector{0, 0}.Normalize()
	assert.Equal(t, 0., result.X)
	assert.Equal(t, 0., result.Y)
}
