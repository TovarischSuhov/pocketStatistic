package log

import (
	"fmt"
	"time"
)

type logLevel int

const (
	DebugLevel logLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const esc = "\u001B["

var colors = map[logLevel]string{
	DebugLevel: "",
	InfoLevel:  esc + "00;32m",
	WarnLevel:  esc + "00;33m",
	ErrorLevel: esc + "00;31m",
	FatalLevel: esc + "01;31m",
}

const finishColor = esc + "m"

var (
	UseColors = false
	LogLevel  = InfoLevel
)

func printMsg(level logLevel, msg string, args ...interface{}) {
	if level < LogLevel {
		return
	}
	ts := time.Now().Format(time.RFC3339)
	prefix := "[" + level.String() + "]"
	if UseColors {
		prefix = colors[level] + prefix + finishColor
	}
	fmt.Printf("%s %s %s\n", ts, prefix, fmt.Sprintf(msg, args...))
}

func Debug(msg string, args ...interface{}) {
	printMsg(DebugLevel, msg, args...)
}

func Info(msg string, args ...interface{}) {
	printMsg(InfoLevel, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	printMsg(WarnLevel, msg, args...)
}

func Error(msg string, args ...interface{}) {
	printMsg(ErrorLevel, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	printMsg(FatalLevel, msg, args...)
	panic("sends panic")
}
