package nengine

type StagedTimer interface {
	AddStage(int) TimerStage
	SetStage(TimerStage, int)
	Timings() map[TimerStage]int
	Stage() TimerStage
	Start(bool)
	Stop()
	Next()
	Update(int)
	Elapsed(TimerStage) bool
	StagePercent() float64
}

type TimerStage int

const TimerStageIdle TimerStage = 1

type Timer struct {
	stage    TimerStage
	timers   map[TimerStage]int
	counters map[TimerStage]int
	elapsed  TimerStage
	looping  bool
}

func (self *Timer) Init() *Timer {
	self.timers = map[TimerStage]int{}
	self.counters = map[TimerStage]int{}
	self.stage = TimerStageIdle
	return self
}

func (self *Timer) AddStage(duration int) TimerStage {
	max := TimerStageIdle
	for stage := range self.timers {
		if stage > max {
			max = stage
		}
	}
	stage := TimerStageIdle << 1
	if max != TimerStageIdle {
		stage = max << 1
	}
	self.timers[stage] = duration
	return stage
}

func (self *Timer) SetStage(stage TimerStage, duration int) {
	self.timers[stage] = duration
}

func (self *Timer) GetStage(stage TimerStage) int {
	return self.timers[stage]
}

func (self *Timer) Timings() map[TimerStage]int {
	return self.timers
}

func (self *Timer) Stage() TimerStage {
	return self.stage
}

func (self *Timer) Start(looping bool) {
	if self.stage == TimerStageIdle {
		self.looping = looping
		for stage, time := range self.timers {
			self.counters[stage] = time
		}
		self.next()
	}
}

func (s *Timer) Stop() {
	s.stage = TimerStageIdle
}

func (self *Timer) Next() {
	if self.stage != TimerStageIdle {
		self.next()
	}
}

// TODO: Change elapsed to only contain Idle on the first activation of
// the timer if the timer is looping

// Moves the timer to the next state
func (self *Timer) next() {
	self.elapsed |= self.stage
	// If there is no next step, unless looping == true
	if _, ok := self.timers[self.stage<<1]; !ok {
		self.stage = TimerStageIdle
		if self.looping {
			self.Start(self.looping)
		}
		return
	}
	// For any other state, move to the next state
	self.stage <<= 1
	if self.timers[self.stage] == 0 {
		self.next()
	}
}

func (self *Timer) Update(delta int) {
	self.elapsed = 0
	self.update(delta)
}

func (self *Timer) update(delta int) {
	if delta <= 0 || self.stage == TimerStageIdle {
		return
	}
	self.counters[self.stage] -= delta
	if self.counters[self.stage] <= 0 {
		rem := self.counters[self.stage] * -1
		self.next()
		self.update(rem)
	}
}

func (self *Timer) StagePercent() float64 {
	time := self.timers[self.stage]
	count := self.counters[self.stage]
	return float64(time-count) / float64(time)
}

// Elapsed returns a bitmask of all timer stages that completed
// during the current update frame
func (self *Timer) Elapsed(stage TimerStage) bool {
	return self.elapsed&stage == stage
}
