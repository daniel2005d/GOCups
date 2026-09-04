package utils

import (
	"github.com/fatih/color"
)

type Logger struct {
	information *color.Color
	err         *color.Color
	warning     *color.Color
}

func NewPrint() *Logger {
	return &Logger{
		information: color.New(color.FgGreen),
		err:         color.New(color.FgRed),
		warning:     color.New(color.FgYellow),
	}
}

func (l *Logger) Info(message string, a ...any) {
	l.information.Printf(message+"\n", a...)
}

func (l *Logger) Error(err error) {

	e := l.err.PrintFunc()
	e(err)
}

func (l *Logger) Warn(message string, a ...any) {
	l.warning.Printf(message+"\n", a...)
}
