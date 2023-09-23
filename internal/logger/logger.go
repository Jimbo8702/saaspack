package logger

import "fmt"

type Logger interface {
	Log(level string, message string, data any)
}
// add a switch to get the datatype and then printf with the correct type

type FMTLogger struct {}

func NewFMTLogger() *FMTLogger {
	return &FMTLogger{}
}

func (fl *FMTLogger) Log(l, m string, d any) {
	fmt.Printf("level=%s message=%s data=%d\n", l, m, d)
}