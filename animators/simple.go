package animators

import (
	"github.com/SnareChops/nengine/types"
)

type SimpleFrame struct {
	Duration int
	Image    types.Image
}

type SimpleAnimator struct {
	frames   []SimpleFrame
	loop     bool
	active   bool
	cooldown int
	index    int
}

func (self *SimpleAnimator) Init(frames []SimpleFrame) *SimpleAnimator {
	self.frames = frames
	return self
}

func (self *SimpleAnimator) Start(loop bool) {
	self.loop = loop
	self.cooldown = self.frames[0].Duration
	self.active = true
}

func (self *SimpleAnimator) IsActive() bool {
	return self.active
}

func (self *SimpleAnimator) Index() int {
	return self.index
}

func (self *SimpleAnimator) Update(delta int) {
	if !self.active {
		return
	}
	for delta > 0 {
		delta = self.update(delta)
	}
}

func (self *SimpleAnimator) update(delta int) int {
	self.cooldown -= delta
	if self.cooldown <= 0 {
		rem := self.cooldown * -1
		self.next()
		if rem > 0 {
			return rem
		}
	}
	return 0
}

func (self *SimpleAnimator) next() {
	self.index++
	if self.index >= len(self.frames) {
		self.index = 0
		if !self.loop {
			self.active = false
			return
		}
	}
	self.cooldown = self.frames[self.index].Duration
}

func (self *SimpleAnimator) Image() types.Image {
	return self.frames[self.index].Image
}
