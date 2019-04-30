package log

import "fmt"

type Logger interface {
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
}

type logger struct{}

func Init(path string) {

}

func Info(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func Warn(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func Error(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
