package debug

import (
	"fmt"
	"time"
)

var timers = []*DebugTimer{}

func EnableTimers(enable bool) {
	// Enable timer stats, early exit
	if enable {
		enableTimers()
		return
	}
	// Remove timer stats
	disableTimers()
}

func enableTimers() {
	for _, timer := range timers {
		timer.enabled = true
		DebugStat(timer.name, timer.Value)
	}
}

func disableTimers() {
	for _, timer := range timers {
		timer.enabled = false
		RemoveStat(timer.name)
	}
}

type DebugTimer struct {
	name    string
	buffer  [10]int64
	pointer int
	start   time.Time
	peak    int64
	enabled bool
}

func NewDebugTimer(name string) *DebugTimer {
	return &DebugTimer{name: name, buffer: [10]int64{}}
}

func (self *DebugTimer) Start() {
	if self.enabled {
		self.start = time.Now()
	}
}

func (self *DebugTimer) End() {
	if self.enabled {
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
