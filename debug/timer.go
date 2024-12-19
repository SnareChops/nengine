package debug

import (
	"fmt"
	"time"
)

var timers = []*DebugTimer{}
var enableTimers bool

func EnableTimers(enable bool) {
	enableTimers = enable
}

type DebugTimer struct {
	name    string
	buffer  [10]int64
	pointer int
	start   time.Time
	peak    int64
}

func NewDebugTimer(name string) *DebugTimer {
	timer := &DebugTimer{name: name, buffer: [10]int64{}}
	timers = append(timers, timer)
	return timer
}

func (self *DebugTimer) Start() {
	if enableTimers {
		self.start = time.Now()
	}
}

func (self *DebugTimer) End() {
	if enableTimers {
		delta := int64(time.Since(self.start))
		self.buffer[self.pointer] = delta
		self.pointer = (self.pointer + 1) % len(self.buffer)
		if delta > self.peak {
			self.peak = delta
		}
	}
}

func (self *DebugTimer) Value() string {
	return fmt.Sprintf("avg: %0.2f, peak: %0.2f", self.Average(), self.Peak())
}

// Average returns the average time in ms
func (self *DebugTimer) Average() float64 {
	var total float64
	for _, delta := range self.buffer {
		total += float64(delta)
	}
	return (total / float64(len(self.buffer))) / 1_000_000
}

// Peak returns the peak time in ms
func (self *DebugTimer) Peak() float64 {
	return float64(self.peak) / 1_000_000
}
