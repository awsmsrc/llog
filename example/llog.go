package main

import (
	"github.com/lukeroberts1990/llog"
	"io/ioutil"
)

func main() {
	//Set the global log level
	_ = llog.SetLevel(llog.LevelDebug)

	//simple functions that pipe to Stdout
	llog.Debug("debug")
	llog.Debugf("debug %v", "test string")
	llog.Info("info")
	llog.Infof("info %v", "test string")
	llog.Warn("warn")
	llog.Warnf("warn %v", "test string")
	llog.Error("error")
	llog.Errorf("error %v", "test string")
	llog.Success("success")
	llog.Successf("success %v", "test string")

	//specify the writer
	logger := &llog.Logger{ioutil.Discard}
	logger.Debug("Piped to a specified io.Writer")
}
