package animators_test

import (
	"testing"

	"github.com/SnareChops/nengine/animators"
	"github.com/stretchr/testify/assert"
)

func TestSimpleAnimator(t *testing.T) {
	a := new(animators.SimpleAnimator).Init([]animators.SimpleFrame{
		{Duration: 20, Image: nil},
		{Duration: 10, Image: nil},
		{Duration: 5, Image: nil},
		{Duration: 6, Image: nil},
	})

	assert.False(t, a.IsActive())
	a.Start(false)
	assert.True(t, a.IsActive())

	assert.Equal(t, 0, a.Index())
	a.Update(10) //10
	assert.Equal(t, 0, a.Index())

	a.Update(10) //20
	assert.Equal(t, 1, a.Index())

	a.Update(10) //30
	assert.Equal(t, 2, a.Index())

	a.Update(10) // 40
	assert.Equal(t, 3, a.Index())

	a.Update(1) // 41
	assert.Equal(t, 0, a.Index())
	assert.False(t, a.IsActive())
}
