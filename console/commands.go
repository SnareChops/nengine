package console

type Result byte

const (
	ResultNormal Result = iota
	ResultInfo
	ResultWarn
	ResultError
	ResultContinue
)

type ConsoleResult struct {
	code    Result
	message string
	cont    ConsoleContinueFunc
}

type ConsoleContinueFunc = func(x, y int) ConsoleResult
type ConsoleFunc = func(args []string) ConsoleResult

var registered = map[string]ConsoleFunc{}

func ConsoleRegister(name string, fn ConsoleFunc) {
	registered[name] = fn
}

func ConsoleUnregister(name string) {
	delete(registered, name)
}

func RunCommand(value string) {
	split := parseArguments(value)
	name := split[0]
	args := split[1:]
	if fn, ok := registered[name]; ok {
		result := fn(args)
		handleResult(result)
		return
	}
	handleResult(NewConsoleResult(ResultError, "Command "+name+" not registered", nil))
}

func handleResult(result ConsoleResult) {
	switch result.code {
	case ResultContinue:
		setHint(result.message)
		setContinue(result.cont)
	default:
		addResult(result)
	}
}

func parseArguments(value string) []string {
	result := []string{}
	arg := []rune{}
	inQuote := false
	for _, char := range value {
		switch char {
		case '"':
			if inQuote {
				result = append(result, string(arg))
				arg = []rune{}
				break
			}
			inQuote = true
		case ' ':
			if inQuote {
				arg = append(arg, char)
				break
			}
			result = append(result, string(arg))
			arg = []rune{}
		default:
			arg = append(arg, char)
		}
	}
	if len(arg) > 0 {
		result = append(result, string(arg))
	}
	return result
}

func NewConsoleResult(code Result, message string, contFunc ConsoleContinueFunc) ConsoleResult {
	return ConsoleResult{
		code:    code,
		message: message,
		cont:    contFunc,
	}
}
