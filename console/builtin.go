package console

import (
	"strings"

	"github.com/SnareChops/nengine/debug"
)

func init() {
	ConsoleRegister("help", Help)
	ConsoleRegister("echo", Echo)
	ConsoleRegister("stats", Stats)
	ConsoleRegister("timers", Timers)
}

func Help(args []string) (ConsoleResult, string) {
	var names []string
	for key := range registered {
		names = append(names, key)
	}
	return ConsoleResultNormal, "Available commands: " + strings.Join(names, ", ")
}

func Echo(args []string) (ConsoleResult, string) {
	return ConsoleResultError, strings.Join(args, " ")
}

func Stats(args []string) (ConsoleResult, string) {
	if len(args) == 0 {
		return ConsoleResultError, "Missing sub-command. Expected one of: show, hide"
	}
	switch args[0] {
	case "show":
		debug.EnableStats(true)
		return ConsoleResultNormal, "Displaying debug stats"
	case "hide":
		debug.EnableStats(false)
		return ConsoleResultNormal, "Removing debug stats"
	default:
		return ConsoleResultError, "Invalid sub-command. Expected one of: show, hide"
	}
}

func Timers(args []string) (ConsoleResult, string) {
	if len(args) == 0 {
		return ConsoleResultError, "Missing sub-command. Expected one of: show, hide"
	}
	switch args[0] {
	case "show":
		debug.EnableTimers(true)
		return ConsoleResultNormal, "Displaying timers"
	case "hide":
		debug.EnableTimers(false)
		return ConsoleResultNormal, "Removed timers"
	default:
		return ConsoleResultError, "Invalid sub-command. Expected one of: show, hide"
	}
}
