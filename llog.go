package llog

import (
	"errors"
	"fmt"
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

func coloredLog(color string, s string) {
	fmt.Fprintf(
		os.Stdout,
		"%v[%v] %v\n",
		color,
		time.Now().Format("2006-01-02  15:04:05"),
		s,
	)
}

func coloredLogF(color string, s string, a ...interface{}) {
	fmt.Fprintf(
		os.Stdout,
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

func Debug(s string) {
	if level > LevelDebug || level == LevelNull {
		return
	}
	coloredLog(YELLOW, s)
}

func FDebug(s string, a ...interface{}) {
	if level > LevelDebug || level == LevelNull {
		return
	}
	coloredLogF(YELLOW, s, a)
}

func Info(s string) {
	if level > LevelInfo || level == LevelNull {
		return
	}
	coloredLog(CYAN, s)
}

func FInfo(s string, a ...interface{}) {
	if level > LevelInfo || level == LevelNull {
		return
	}
	coloredLogF(CYAN, s, a)
}

func Warn(s string) {
	if level > LevelWarn || level == LevelNull {
		return
	}
	coloredLog(MAGENTA, s)
}

func FWarn(s string, a ...interface{}) {
	if level > LevelWarn || level == LevelNull {
		return
	}
	coloredLogF(MAGENTA, s, a)
}

func Error(s string) {
	coloredLog(RED, s)
}

func FError(s string, a ...interface{}) {
	coloredLogF(RED, s, a)
}

func Success(s string) {
	coloredLog(GREEN, s)
}

func FSuccess(s string, a ...interface{}) {
	coloredLogF(GREEN, s, a)
}
