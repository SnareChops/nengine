package debug

import "time"

var FrameTimers = []*FrameTimer{}

type FrameTimer struct {
	*DebugTimer
	accumulator int64
}

func NewFrameTimer(name string, auto bool) *FrameTimer {
	timer := &FrameTimer{DebugTimer: NewDebugTimer(name)}
	if auto == true {
		FrameTimers = append(FrameTimers, timer)
	}
	return timer
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
