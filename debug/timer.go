package debug

import (
	"fmt"
	"time"
)

type DebugTimer struct {
	name    string
	buffer  [10]int64
	pointer int
	start   time.Time
	peak    int64
}

func NewDebugTimer(name string) *DebugTimer {
	timer := &DebugTimer{name: name, buffer: [10]int64{}}
	DebugStat(name, func() string {
		return fmt.Sprintf("avg: %d, peak: %d", timer.Average(), timer.Peak())
	})
	return timer
}

func (self *DebugTimer) Start() {
	self.start = time.Now()
}

func (self *DebugTimer) End() {
	delta := int64(time.Since(self.start))
	self.buffer[self.pointer] = delta
	self.pointer = (self.pointer + 1) % len(self.buffer)
	if delta > self.peak {
		self.peak = delta
	}
}

// Average returns the average time in ms
func (self *DebugTimer) Average() int64 {
	var total int64
	for _, delta := range self.buffer {
		total += delta
	}
	return (total / int64(len(self.buffer))) / 1_000_000
}

// Peak returns the peak time in ms
func (self *DebugTimer) Peak() int64 {
	return self.peak / 1_000_000
}
