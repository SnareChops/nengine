package console

import "strings"

type ConsoleResult byte

const (
	ConsoleResultNormal ConsoleResult = iota
	ConsoleResultInfo
	ConsoleResultWarn
	ConsoleResultError
)

type ConsoleFunc = func(args []string) (ConsoleResult, string)

var registered = map[string]ConsoleFunc{}

func ConsoleRegister(name string, fn ConsoleFunc) {
	registered[name] = fn
}

func RunCommand(value string) (ConsoleResult, string) {
	split := strings.Split(value, " ")
	name := split[0]
	args := split[1:]
	if fn, ok := registered[name]; ok {
		return fn(args)
	}
	return ConsoleResultError, "Command " + name + " not registered"
}
