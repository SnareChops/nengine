package debug

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

var stats = []stat{}
var enableStats bool

type stat struct {
	label string
	value func() string
}

func EnableStats(enabled bool) {
	// Enable default stats and early exit
	enableStats = enabled
	if enabled {
		DebugStat("TPS", func() string {
			return fmt.Sprintf("%0.2f", ebiten.ActualTPS())
		})
		DebugStat("FPS", func() string {
			return fmt.Sprintf("%0.2f", ebiten.ActualFPS())
		})
		return
	}
	// Disable default stats
	RemoveStat("TPS")
	RemoveStat("FPS")
}

func DebugStat(label string, value func() string) {
	for i := range stats {
		if stats[i].label == label {
			stats[i].value = value
			return
		}
	}
	stats = append(stats, stat{label, value})
}

func RemoveStat(label string) {
	for i, stat := range stats {
		if stat.label == label {
			stats = append(stats[:i], stats[i+1:]...)
		}
	}
}
