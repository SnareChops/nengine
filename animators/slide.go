package animators

import (
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type SlideAnimator struct {
	slider   int
	duration int
	frame    int
	frames   []*ebiten.Image
}

func (self *SlideAnimator) Init(duration int, frames []*ebiten.Image) *SlideAnimator {
	self.frames = frames
	self.duration = duration
	return self
}

func (self *SlideAnimator) Update(advance bool, delta int) bool {
	prev := self.frame
	if advance {
		self.slider += delta
	} else {
		self.slider -= delta
	}
	self.slider = utils.Clamp(self.slider, 0, self.duration)
	self.frame = int(float64(self.slider) / float64(self.duration) * float64(len(self.frames)-1))
	return prev != self.frame
}

func (self *SlideAnimator) Image() *ebiten.Image {
	return self.frames[self.frame]
}
