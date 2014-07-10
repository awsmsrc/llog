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
	llog.FDebug("debug %v", "test string")
	llog.Info("info")
	llog.FInfo("info %v", "test string")
	llog.Warn("warn")
	llog.FWarn("warn %v", "test string")
	llog.Error("error")
	llog.FError("error %v", "test string")
	llog.Success("success")
	llog.FSuccess("success %v", "test string")

	//specify the writer
	logger := &llog.Logger{ioutil.Discard}
	logger.Debug("Piped to a specified io.Writer")
}
