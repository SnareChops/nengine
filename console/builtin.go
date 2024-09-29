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

func Help(args []string) ConsoleResult {
	var names []string
	for key := range registered {
		names = append(names, key)
	}
	return NewConsoleResult(
		ResultNormal,
		"Available commands: "+strings.Join(names, ", "),
		nil,
	)
}

func Echo(args []string) ConsoleResult {
	return NewConsoleResult(
		ResultError,
		strings.Join(args, " "),
		nil,
	)
}

func Stats(args []string) ConsoleResult {
	if len(args) == 0 {
		return NewConsoleResult(
			ResultError,
			"Missing sub-command. Expected one of: show, hide",
			nil,
		)
	}
	switch args[0] {
	case "show":
		debug.EnableStats(true)
		return NewConsoleResult(
			ResultNormal,
			"Displaying debug stats",
			nil,
		)
	case "hide":
		debug.EnableStats(false)
		return NewConsoleResult(
			ResultNormal,
			"Removing debug stats",
			nil,
		)
	default:
		return NewConsoleResult(
			ResultError,
			"Invalid sub-command. Expected one of: show, hide",
			nil,
		)
	}
}

func Timers(args []string) ConsoleResult {
	if len(args) == 0 {
		return NewConsoleResult(
			ResultError,
			"Missing sub-command. Expected one of: show, hide",
			nil,
		)
	}
	switch args[0] {
	case "show":
		debug.EnableTimers(true)
		return NewConsoleResult(
			ResultNormal,
			"Displaying timers",
			nil,
		)
	case "hide":
		debug.EnableTimers(false)
		return NewConsoleResult(
			ResultNormal,
			"Removed timers",
			nil,
		)
	default:
		return NewConsoleResult(
			ResultError,
			"Invalid sub-command. Expected one of: show, hide",
			nil,
		)
	}
}
