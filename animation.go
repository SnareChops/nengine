package nengine

import "github.com/hajimehoshi/ebiten/v2"

// Animation represents timings for a full animation
type Animation []AnimationFrame

// AnimationFrame represents timings for a single animation frame
type AnimationFrame struct {
	Start    int
	Duration int
	Frame    int
}

// AddFrame adds a new frame to an animation
func (self Animation) AddFrame(start, duration, index int) Animation {
	self = append(self, AnimationFrame{start, duration, index})
	return self
}

// Animator top level struct for managing animations
// Tip: Use in combination with Bounds to create an animated sprite
type Animator struct {
	frames     []*ebiten.Image
	animations map[string]Animation
	active     Animation
	frame      int
	elapsed    int
	loop       bool
}

// Init sets the initial state for the Animator
func (self *Animator) Init(frames []*ebiten.Image) *Animator {
	self.frames = frames
	self.animations = map[string]Animation{}
	return self
}

// AddAnimation Adds a new named animation to the Animator
func (self *Animator) AddAnimation(name string, animation Animation) {
	self.animations[name] = animation
}

// ClearAnimation Clears the currently active animation
// and returns the image to the default image
func (self *Animator) ClearAnimation() {
	self.active = nil
	self.loop = false
	self.elapsed = 0
}

// StartAnimation Starts an animation by it's name
func (self *Animator) StartAnimation(name string, loop bool) {
	self.active = self.animations[name]
	self.loop = loop
	self.elapsed = 0
}

// Image Returns the current active image for the animation
func (self *Animator) Image() *ebiten.Image {
	return self.frames[self.frame]
}

// Update Call this on every frame to "run" the animation
func (self *Animator) Update(delta int) {
	if self.active == nil {
		return
	}
	last := self.active[len(self.active)-1]
	total := last.Start + last.Duration
	self.elapsed += delta
	if self.elapsed >= total && !self.loop {
		self.ClearAnimation()
	}
	self.elapsed %= total
	for _, frame := range self.active {
		if self.elapsed > frame.Start {
			self.frame = frame.Frame
		}
	}
}
