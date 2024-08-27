package console

import "strings"

func init() {
	Register("echo", Echo)
}

func Echo(args []string) string {
	return strings.Join(args, " ")
}
