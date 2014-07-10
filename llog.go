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
	RESET   = "\033[0m"
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

func write(w io.Writer, color string, s interface{}, a ...interface{}) {
	fmt.Fprintf(
		w,
		"%v[%v] %v%v\n",
		color,
		time.Now().Format("2006/01/02-15:04:05"),
		fmt.Sprintf(fmt.Sprint(s), a...),
		RESET,
	)
}

func SetLevel(i int) error {
	if i < 1 || i > 4 {
		return errors.New("log level out of range")
	}
	level = i
	return nil
}

func (l *Logger) Debug(s interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, YELLOW, s)
	}
}

func (l *Logger) Debugf(s interface{}, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, YELLOW, s, a)
	}
}

func (l *Logger) Info(s interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(l, CYAN, s)
	}
}

func (l *Logger) Infof(s interface{}, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, CYAN, s, a)
	}
}

func (l *Logger) Warn(s interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(l, MAGENTA, s)
	}
}

func (l *Logger) Warnf(s interface{}, a ...interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(l, MAGENTA, s, a)
	}
}

func (l *Logger) Error(s interface{}) {
	write(l, RED, s)
}

func (l *Logger) Errorf(s interface{}, a ...interface{}) {
	write(l, RED, s, a)
}

func (l *Logger) Success(s interface{}) {
	write(l, GREEN, s)
}

func (l *Logger) Successf(s interface{}, a ...interface{}) {
	write(l, GREEN, s, a)
}

func (l *Logger) FATAL(err error) {
	write(os.Stdout, RED, err)
	os.Exit(1)
}

func Debug(s interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(os.Stdout, YELLOW, s)
	}
}

func Debugf(s interface{}, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(os.Stdout, YELLOW, s, a)
	}
}

func Info(s interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(os.Stdout, CYAN, s)
	}
}

func Infof(s interface{}, a ...interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(os.Stdout, CYAN, s, a)
	}
}

func Warn(s interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(os.Stdout, MAGENTA, s)
	}
}

func Warnf(s interface{}, a ...interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(os.Stdout, MAGENTA, s, a)
	}
}

func Error(s interface{}) {
	write(os.Stdout, RED, s)
}

func Errorf(s interface{}, a ...interface{}) {
	write(os.Stdout, RED, s, a)
}

func FATAL(err error) {
	write(os.Stdout, RED, err)
	os.Exit(1)
}

func Success(s interface{}) {
	write(os.Stdout, GREEN, s)
}

func Successf(s interface{}, a ...interface{}) {
	write(os.Stdout, GREEN, s, a)
}
