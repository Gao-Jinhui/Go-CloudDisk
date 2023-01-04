package system

import (
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	Logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}
	Logger.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
}

func GetLoggerEntry(traceId string) *logrus.Entry {
	entry := Logger.WithFields(logrus.Fields{
		"trace_id": traceId,
	})
	return entry
}
