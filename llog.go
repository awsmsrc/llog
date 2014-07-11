//llog is a simple *leveled* logging packages originally writen and used in Go
//services at Sqwiggle (https://www.sqwiggle.com). The API is designed to be
//conceptually similar to the standard library's log package
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

//The available log levels are { levelNull:no logging, levelDebug:log out
//everything, levelInfo: log out Info, Warn, Error, Succes and FATAL, LevelWarn:
//log out Warn, Error, Success and FATAL and finally LevelError: log out Error,
//Success and FATAL}
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

//llog.Logger wraps/embeds a writer so that you can log where ever you like.
//This package defaults to logging to os.Stdout
type Logger struct {
	io.Writer
}

func write(w io.Writer, color string, s string, a ...interface{}) {
	fmt.Fprintf(
		w,
		"%v[%v] %v%v\n",
		color,
		time.Now().Format("2006/01/02-15:04:05"),
		fmt.Sprintf(s, a...),
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
		write(l, YELLOW, fmt.Sprint(s))
	}
}

func (l *Logger) Debugf(s interface{}, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, YELLOW, fmt.Sprint(s), a...)
	}
}

func (l *Logger) Info(s interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(l, CYAN, fmt.Sprint(s))
	}
}

func (l *Logger) Infof(s interface{}, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(l, CYAN, fmt.Sprint(s), a...)
	}
}

func (l *Logger) Warn(s interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(l, MAGENTA, fmt.Sprint(s))
	}
}

func (l *Logger) Warnf(s interface{}, a ...interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(l, MAGENTA, fmt.Sprint(s), a...)
	}
}

func (l *Logger) Error(s interface{}) {
	write(l, RED, fmt.Sprint(s))
}

func (l *Logger) Errorf(s interface{}, a ...interface{}) {
	write(l, RED, fmt.Sprint(s), a...)
}

func (l *Logger) Success(s interface{}) {
	write(l, GREEN, fmt.Sprint(s))
}

func (l *Logger) Successf(s interface{}, a ...interface{}) {
	write(l, GREEN, fmt.Sprint(s), a...)
}

func (l *Logger) FATAL(s interface{}) {
	write(os.Stdout, RED, fmt.Sprint(s))
	os.Exit(1)
}

func Debug(s interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(os.Stdout, YELLOW, fmt.Sprint(s))
	}
}

func Debugf(s interface{}, a ...interface{}) {
	if !(level > LevelDebug || level == LevelNull) {
		write(os.Stdout, YELLOW, fmt.Sprint(s), a...)
	}
}

func Info(s interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(os.Stdout, CYAN, fmt.Sprint(s))
	}
}

func Infof(s interface{}, a ...interface{}) {
	if !(level > LevelInfo || level == LevelNull) {
		write(os.Stdout, CYAN, fmt.Sprint(s), a...)
	}
}

func Warn(s interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(os.Stdout, MAGENTA, fmt.Sprint(s))
	}
}

func Warnf(s interface{}, a ...interface{}) {
	if !(level > LevelWarn || level == LevelNull) {
		write(os.Stdout, MAGENTA, fmt.Sprint(s), a...)
	}
}

func Error(s interface{}) {
	write(os.Stdout, RED, fmt.Sprint(s))
}

func Errorf(s interface{}, a ...interface{}) {
	write(os.Stdout, RED, fmt.Sprint(s), a...)
}

func FATAL(s interface{}) {
	write(os.Stdout, RED, fmt.Sprint(s))
	os.Exit(1)
}

func Success(s interface{}) {
	write(os.Stdout, GREEN, fmt.Sprint(s))
}

func Successf(s interface{}, a ...interface{}) {
	write(os.Stdout, GREEN, fmt.Sprint(s), a...)
}
