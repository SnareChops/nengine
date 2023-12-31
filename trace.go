package nengine

import (
	"os"
	"runtime/trace"
)

func Trace(fn func()) {
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	fn()
}
