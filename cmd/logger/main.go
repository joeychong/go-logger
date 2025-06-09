package main

import (
	"ebizzone.com/logger"
	"ebizzone.com/logger/log"
)

func main() {
	log.CurrentLevel = log.LOG_DEBUG
	console := logger.New("main")
	console.Debug("This is some test")
	console.Info("This is some info")
	console.Debugf("this is some %s", "234")

	log := logger.New("log")
	log.Debug("This is some test")
	log.Info("This is some info")
	log.Debugf("this is some %s", "234")
	log.Debug(123)
	log.Warn("This is warning")
	log.Error("This is error")
	log.Fatal("this is fatal")
}
