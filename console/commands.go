package console

import "strings"

type CommandFunc = func(args []string) string

var registered = map[string]CommandFunc{}

func Register(name string, fn CommandFunc) {
	registered[name] = fn
}

func RunCommand(value string) string {
	split := strings.Split(value, " ")
	name := split[0]
	args := split[1:]
	if fn, ok := registered[name]; ok {
		return fn(args)
	}
	return "Command " + name + " not registered"
}
