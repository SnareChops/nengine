package debug

import "time"

type FrameTimer struct {
	*DebugTimer
	accumulator int64
}

func NewFrameTimer(name string) *FrameTimer {
	return &FrameTimer{DebugTimer: NewDebugTimer(name)}
}

func (self *FrameTimer) End() {
	delta := int64(time.Since(self.start))
	self.accumulator += delta
}

func (self *FrameTimer) EndFrame() {
	self.buffer[self.pointer] = self.accumulator
	self.pointer = (self.pointer + 1) % len(self.buffer)
	if self.accumulator > self.peak {
		self.peak = self.accumulator
	}
	self.accumulator = 0
}
