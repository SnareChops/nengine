package animators

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ga struct {
	name   string
	loop   bool // Should the animation repeat until stopped or switched
	repeat int  // How many times the animation should play (only if loop is false)
	frames []GeneralFrame
	// TODO: pingpong bool // Should the animation play forwards then backwards
}

type GeneralFrame struct {
	duration int
	image    *ebiten.Image
}

func NewGeneralFrame(duration int, image *ebiten.Image) GeneralFrame {
	return GeneralFrame{
		duration: duration,
		image:    image,
	}
}

type GeneralAnimator struct {
	animations map[string]ga
	idle       ga
	activeName string
	active     ga
	frame      int
	cooldown   int
}

func (self *GeneralAnimator) Init(idle []GeneralFrame) *GeneralAnimator {
	self.idle = ga{loop: true, frames: idle}
	self.animations = map[string]ga{}
	self.activate(self.idle)
	return self
}

func (self *GeneralAnimator) Add(name string, loop bool, repeat int, frames []GeneralFrame) {
	self.animations[name] = ga{
		name:   name,
		loop:   loop,
		repeat: repeat,
		frames: frames,
	}
}

func (self *GeneralAnimator) Play(name string) {
	if self.active.name == name {
		return
	}
	if anim, ok := self.animations[name]; ok {
		self.activate(anim)
	}
}

func (self *GeneralAnimator) Stop() {
	self.activate(self.idle)
}

func (self *GeneralAnimator) Update(delta int) {
	for delta > 0 {
		delta = self.update(delta)
	}
}

func (self *GeneralAnimator) activate(anim ga) {
	self.active = anim
	self.frame = 0
	self.cooldown = self.active.frames[self.frame].duration
}

func (self *GeneralAnimator) update(delta int) int {
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

func (self *GeneralAnimator) next() {
	self.frame++
	if self.frame >= len(self.active.frames) {
		self.frame = 0
		if !self.active.loop {
			self.activate(self.idle)
			return
		}
	}
	self.cooldown = self.active.frames[self.frame].duration
}

func (self *GeneralAnimator) Image() *ebiten.Image {
	return self.active.frames[self.frame].image
}
