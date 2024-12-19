package animators_test

import (
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/SnareChops/nengine/animators"
	"github.com/stretchr/testify/assert"
)

func TestGeneralAnimator(t *testing.T) {
	nengine.EnableMocks()

	i1 := nengine.NewImage(1, 1)
	i2 := nengine.NewImage(1, 1)
	frame1 := animators.NewGeneralFrame(10, i1)
	frame2 := animators.NewGeneralFrame(10, i2)

	anim := new(animators.GeneralAnimator).Init([]animators.GeneralFrame{frame1, frame2})
	assert.Equal(t, i1, anim.Image())
	anim.Update(10)
	assert.Equal(t, i2, anim.Image())
	anim.Update(10)
	assert.Equal(t, i1, anim.Image())

	i3 := nengine.NewImage(1, 1)
	i4 := nengine.NewImage(1, 1)
	frame3 := animators.NewGeneralFrame(10, i3)
	frame4 := animators.NewGeneralFrame(10, i4)

	anim.Add("test", false, 0, []animators.GeneralFrame{frame3, frame4})
	anim.Play("test")
	assert.Equal(t, i3, anim.Image())
	anim.Update(10)
	assert.Equal(t, i4, anim.Image())
	anim.Update(10)
	assert.Equal(t, i1, anim.Image())

	anim.Play("test")
	assert.Equal(t, i3, anim.Image())
	anim.Update(10)
	assert.Equal(t, i4, anim.Image())
	anim.Stop()
	assert.Equal(t, i1, anim.Image())
}
