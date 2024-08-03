package utils_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/utils"
	"github.com/stretchr/testify/assert"
)

func TestDoesCollide(t *testing.T) {
	a := bounds.NewBox(4, 4)
	b := bounds.NewBox(4, 4)
	a.SetPos2(8, 8)
	b.SetPos2(20, 20)
	assert.False(t, utils.DoesCollide(a, b))

	b.SetPos2(6, 5)
	assert.True(t, utils.DoesCollide(a, b))

	b.SetPos2(5, 6)
	assert.True(t, utils.DoesCollide(a, b))

	b.SetPos2(10, 5)
	assert.True(t, utils.DoesCollide(a, b))

	b.SetPos2(5, 10)
	assert.True(t, utils.DoesCollide(a, b))

	b.SetPos2(7, 7)
	assert.True(t, utils.DoesCollide(a, b))
}

func TestIsWithin(t *testing.T) {
	box := bounds.NewBox(4, 8)
	assert.False(t, utils.IsWithin(box, -1, 4))
	assert.False(t, utils.IsWithin(box, 2, -1))
	assert.False(t, utils.IsWithin(box, 4, 4))
	assert.False(t, utils.IsWithin(box, 2, 8))

	assert.True(t, utils.IsWithin(box, 2, 0))
	assert.True(t, utils.IsWithin(box, 0, 4))
	assert.True(t, utils.IsWithin(box, 3, 4))
	assert.True(t, utils.IsWithin(box, 2, 7))
}
