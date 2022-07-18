package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	//log.SetOutput(os.Stdout)

	file, _ := os.OpenFile("./log.log", os.O_WRONLY|os.O_APPEND, 0666)

	w := io.MultiWriter(file, os.Stdout)

	log.SetOutput(w)

	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"level": "debug",
	}).Debug("debug")

	log.WithFields(log.Fields{
		"level": "info",
	}).Info("info")

	log.WithFields(log.Fields{
		"level": "warn",
	}).Warn("warn")

	log.WithFields(log.Fields{
		"level": "error",
	}).Error("error")

	logger := log.New()
	file, _ := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 06666)
	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetLevel(log.InfoLevel)

	logger.WithFields(log.Fields{
		"level": "error",
	}).Error("error")

	logger.WithFields(log.Fields{
		"level": "info",
	}).Info("info")

	logger.AddHook(new(myHook))

	logger.WithFields(log.Fields{
		"level": "info",
	}).Info("info") //不会触发hook

	logger.WithFields(log.Fields{
		"level": "error",
	}).Error("error")

}

type myHook struct {
}

func (h *myHook) Levels() []log.Level {
	return []log.Level{log.ErrorLevel}
}

func (h *myHook) Fire(entry *log.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}
