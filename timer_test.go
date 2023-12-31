package nengine_test

import (
	"fmt"
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/stretchr/testify/assert"
)

func TestSkillTimersStatusProgression(t *testing.T) {
	timer := new(nengine.Timer).Init()
	a := timer.AddStage(30)
	b := timer.AddStage(40)
	c := timer.AddStage(20)
	d := timer.AddStage(10)

	assert.Equal(t, nengine.TimerStageIdle, timer.Stage())
	assert.False(t, timer.Elapsed(nengine.TimerStageIdle))
	timer.Start(false)
	assert.Equal(t, a, timer.Stage())
	assert.True(t, timer.Elapsed(nengine.TimerStageIdle))

	timer.Update(10) // 10
	assert.Equal(t, a, timer.Stage())

	timer.Update(10) // 20
	assert.Equal(t, a, timer.Stage())

	assert.False(t, timer.Elapsed(a))
	timer.Update(10) // 30
	assert.Equal(t, b, timer.Stage())
	assert.True(t, timer.Elapsed(a))

	timer.Update(10) // 10
	assert.Equal(t, b, timer.Stage())

	timer.Update(10) // 20
	assert.Equal(t, b, timer.Stage())

	timer.Update(10) // 30
	assert.Equal(t, b, timer.Stage())

	assert.False(t, timer.Elapsed(b))
	timer.Update(10) // 40
	assert.Equal(t, c, timer.Stage())
	assert.True(t, timer.Elapsed(b))

	timer.Update(10) // 10
	assert.Equal(t, c, timer.Stage())

	assert.False(t, timer.Elapsed(c))
	timer.Update(10) // 20
	assert.Equal(t, d, timer.Stage())
	assert.True(t, timer.Elapsed(c))

	assert.False(t, timer.Elapsed(d))
	timer.Update(10) // 10
	assert.Equal(t, nengine.TimerStageIdle, timer.Stage())
	assert.True(t, timer.Elapsed(d))

	timer.Update(10)
	assert.Equal(t, nengine.TimerStageIdle, timer.Stage())
}

func TestSkillTimerLooping(t *testing.T) {
	timer := new(nengine.Timer).Init()
	a := timer.AddStage(10)

	assert.Equal(t, nengine.TimerStageIdle, timer.Stage())
	assert.False(t, timer.Elapsed(nengine.TimerStageIdle))
	assert.False(t, timer.Elapsed(a))
	timer.Start(true)
	assert.Equal(t, a, timer.Stage())
	assert.True(t, timer.Elapsed(nengine.TimerStageIdle))
	assert.False(t, timer.Elapsed(a))

	timer.Update(5)
	assert.Equal(t, a, timer.Stage())
	assert.False(t, timer.Elapsed(nengine.TimerStageIdle))
	assert.False(t, timer.Elapsed(a))

	timer.Update(15)
	assert.True(t, timer.Elapsed(nengine.TimerStageIdle))
	assert.True(t, timer.Elapsed(a))
	assert.Equal(t, a, timer.Stage())

	timer.Update(5)
	assert.Equal(t, a, timer.Stage())
	assert.False(t, timer.Elapsed(nengine.TimerStageIdle))
	assert.False(t, timer.Elapsed(a))

	timer.Update(5)
	assert.True(t, timer.Elapsed(nengine.TimerStageIdle))
	assert.True(t, timer.Elapsed(a))
	assert.Equal(t, a, timer.Stage())

}

func TestSkillTimerUpdateWithDeltaWhileIdle(t *testing.T) {
	timer := new(nengine.Timer).Init()
	timer.AddStage(150)
	timer.Update(151)

	assert.Equal(t, nengine.TimerStageIdle, timer.Stage())
}

func TestTimerStagePercent(t *testing.T) {
	timer := new(nengine.Timer).Init()
	timer.AddStage(10)
	timer.Start(false)

	fmt.Printf("%#v\n", timer)
	assert.Equal(t, 0., timer.StagePercent())

	timer.Update(5)
	assert.Equal(t, .5, timer.StagePercent())

	timer.Update(2)
	assert.Equal(t, .7, timer.StagePercent())
}

// func TestSkillTimersWindupOnlyProgression(t *testing.T) {
// 	timers := new(core.Timer).Init(10, 0, 0, 0)

// 	windup := false
// 	timers.OnWindup(func() {
// 		windup = true
// 	})

// 	duration := false
// 	timers.OnDuration(func() {
// 		duration = true
// 	})

// 	winddown := false
// 	timers.OnWinddown(func() {
// 		winddown = true
// 	})

// 	cooldown := false
// 	timers.OnCooldown(func() {
// 		cooldown = true
// 	})

// 	idle := false
// 	timers.OnIdle(func() {
// 		idle = true
// 	})

// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.False(t, windup)
// 	timers.Start(false)
// 	assert.Equal(t, skills.WINDUP, timers.Status)
// 	assert.True(t, windup)

// 	assert.False(t, duration)
// 	assert.False(t, winddown)
// 	assert.False(t, cooldown)
// 	assert.False(t, idle)
// 	timers.Update(10)
// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.True(t, duration)
// 	assert.True(t, winddown)
// 	assert.True(t, cooldown)
// 	assert.True(t, idle)
// }

// func TestSkillTimersDurationOnlyProgression(t *testing.T) {
// 	timers := new(core.Timer).Init(0, 10, 0, 0)

// 	windup := false
// 	timers.OnWindup(func() {
// 		windup = true
// 	})

// 	duration := false
// 	timers.OnDuration(func() {
// 		duration = true
// 	})

// 	winddown := false
// 	timers.OnWinddown(func() {
// 		winddown = true
// 	})

// 	cooldown := false
// 	timers.OnCooldown(func() {
// 		cooldown = true
// 	})

// 	idle := false
// 	timers.OnIdle(func() {
// 		idle = true
// 	})

// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.False(t, windup)
// 	assert.False(t, duration)
// 	timers.Start(false)
// 	assert.Equal(t, skills.DURATION, timers.Status)
// 	assert.True(t, windup)
// 	assert.True(t, duration)

// 	assert.False(t, winddown)
// 	assert.False(t, cooldown)
// 	assert.False(t, idle)
// 	timers.Update(10)
// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.True(t, winddown)
// 	assert.True(t, cooldown)
// 	assert.True(t, idle)
// }

// func TestSkillTimersWinddownOnlyProgression(t *testing.T) {
// 	timers := new(core.Timer).Init(0, 0, 10, 0)

// 	windup := false
// 	timers.OnWindup(func() {
// 		windup = true
// 	})

// 	duration := false
// 	timers.OnDuration(func() {
// 		duration = true
// 	})

// 	winddown := false
// 	timers.OnWinddown(func() {
// 		winddown = true
// 	})

// 	cooldown := false
// 	timers.OnCooldown(func() {
// 		cooldown = true
// 	})

// 	idle := false
// 	timers.OnIdle(func() {
// 		idle = true
// 	})

// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.False(t, windup)
// 	assert.False(t, duration)
// 	assert.False(t, winddown)
// 	timers.Start(false)
// 	assert.Equal(t, skills.WINDDOWN, timers.Status)
// 	assert.True(t, windup)
// 	assert.True(t, duration)
// 	assert.True(t, winddown)

// 	assert.False(t, cooldown)
// 	assert.False(t, idle)
// 	timers.Update(10)
// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.True(t, cooldown)
// 	assert.True(t, idle)
// }

// func TestSkillTimersCooldownOnlyProgression(t *testing.T) {
// 	timers := new(core.Timer).Init(0, 0, 0, 10)

// 	windup := false
// 	timers.OnWindup(func() {
// 		windup = true
// 	})

// 	duration := false
// 	timers.OnDuration(func() {
// 		duration = true
// 	})

// 	winddown := false
// 	timers.OnWinddown(func() {
// 		winddown = true
// 	})

// 	cooldown := false
// 	timers.OnCooldown(func() {
// 		cooldown = true
// 	})

// 	idle := false
// 	timers.OnIdle(func() {
// 		idle = true
// 	})

// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.False(t, windup)
// 	assert.False(t, duration)
// 	assert.False(t, winddown)
// 	assert.False(t, cooldown)
// 	timers.Start(false)
// 	assert.Equal(t, skills.COOLDOWN, timers.Status)
// 	assert.True(t, windup)
// 	assert.True(t, duration)
// 	assert.True(t, winddown)
// 	assert.True(t, cooldown)

// 	assert.False(t, idle)
// 	timers.Update(10)
// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.True(t, idle)
// }

// func TestSkillTimersInstantProgression(t *testing.T) {
// 	timers := new(core.Timer).Init(0, 0, 0, 0)

// 	windup := false
// 	timers.OnWindup(func() {
// 		windup = true
// 	})

// 	duration := false
// 	timers.OnDuration(func() {
// 		duration = true
// 	})

// 	winddown := false
// 	timers.OnWinddown(func() {
// 		winddown = true
// 	})

// 	cooldown := false
// 	timers.OnCooldown(func() {
// 		cooldown = true
// 	})

// 	idle := false
// 	timers.OnIdle(func() {
// 		idle = true
// 	})

// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.False(t, windup)
// 	assert.False(t, duration)
// 	assert.False(t, winddown)
// 	assert.False(t, cooldown)
// 	assert.False(t, idle)
// 	timers.Start(false)
// 	assert.Equal(t, skills.IDLE, timers.Status)
// 	assert.True(t, windup)
// 	assert.True(t, duration)
// 	assert.True(t, winddown)
// 	assert.True(t, cooldown)
// 	assert.True(t, idle)
// }
