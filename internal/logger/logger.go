package logger

import "fmt"

type Logger interface {
	Log(level string, message string, data any)
}

type FMTLogger struct {}

func NewFMTLogger() *FMTLogger {
	return &FMTLogger{}
}

func (fl *FMTLogger) Log(l, m string, d any) {
	fmt.Printf("level=%s message=%s data=%d\n", l, m, d)
}