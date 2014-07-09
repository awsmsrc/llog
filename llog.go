package llog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[93m"
	CYAN    = "\033[36m"
	MAGENTA = "\033[95m"
)

const (
	LevelNull = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

var (
	level int = 1
)

type Logger struct {
	io.Writer
}

func write(w io.Writer, color string, s string, a ...interface{}) {
	fmt.Fprintf(
		w,
		"%v[%v] %v\n",
		color,
		time.Now().Format("2006-01-02  15:04:05"),
		fmt.Sprintf(s, a...),
	)
}

func SetLevel(i int) error {
	if i < 1 || i > 4 {
		return errors.New("log level out of range")
	}
	level = i
	return nil
}

func (l *Logger) Debug(s string) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, YELLOW, s)
	}
}

func (l *Logger) FDebug(s string, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, YELLOW, s, a)
	}
}

func (l *Logger) Info(s string) {
	if !(level > LevelInfo || level == LevelNull) {
		write(l, CYAN, s)
	}
}

func (l *Logger) FInfo(s string, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, CYAN, s, a)
	}
}

func (l *Logger) Warn(s string) {
	if !(level > LevelWarn || level == LevelNull) {
		write(l, MAGENTA, s)
	}
}

func (l *Logger) FWarn(s string, a ...interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(l, MAGENTA, s, a)
	}
}

func (l *Logger) Error(s string) {
	write(l, RED, s)
}

func (l *Logger) FError(s string, a ...interface{}) {
	write(l, RED, s, a)
}

func (l *Logger) Success(s string) {
	write(l, GREEN, s)
}

func (l *Logger) FSuccess(s string, a ...interface{}) {
	write(l, GREEN, s, a)
}

func Debug(s string) {
	if !(level > LevelDebug || level == LevelNull) {
		write(os.Stdout, YELLOW, s)
	}
}

func FDebug(s string, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(os.Stdout, YELLOW, s, a)
	}
}

func Info(s string) {
	if !(level > LevelInfo || level == LevelNull) {
		write(os.Stdout, CYAN, s)
	}
}

func FInfo(s string, a ...interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(os.Stdout, CYAN, s, a)
	}
}

func Warn(s string) {
	if !(level > LevelWarn || level == LevelNull) {
		write(os.Stdout, MAGENTA, s)
	}
}

func FWarn(s string, a ...interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(os.Stdout, MAGENTA, s, a)
	}
}

func Error(s string) {
	write(os.Stdout, RED, s)
}

func FError(s string, a ...interface{}) {
	write(os.Stdout, RED, s, a)
}

func Success(s string) {
	write(os.Stdout, GREEN, s)
}

func FSuccess(s string, a ...interface{}) {
	write(os.Stdout, GREEN, s, a)
}
