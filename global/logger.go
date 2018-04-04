package global

import (
	"fmt"
	"os"
	"path"

	configor "github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

type logConfigStruct struct {
	LogDir   string
	LogName  string
	LogLevel string
}

var Logger *log.Logger = log.New()
var logConfig logConfigStruct = logConfigStruct{}
var logFilePointer *os.File = nil

func ReloadLogConfig() {

	var configPath = Arguments.LogConfigPath

	if err := configor.Load(&logConfig, configPath); err != nil {
		panic(err)
	}
	reloadDefaultLogger()
}

func reloadDefaultLogger() {
	if logLevel, err := log.ParseLevel(logConfig.LogLevel); err == nil {
		Logger.SetLevel(logLevel)
	} else {
		panic(err)
	}

	if oneLogFilePointer, err := CreateLogFile("default"); err == nil {
		Logger.Out = oneLogFilePointer
		if logFilePointer != nil {
			logFilePointer.Close()
		}
		logFilePointer = oneLogFilePointer
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}
}

func CreateLogFile(suffix string) (*os.File, error) {
	if oneFilePointer, err := os.OpenFile(path.Join(logConfig.LogDir, fmt.Sprintf("%s.%s.log", logConfig.LogName, suffix)), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err == nil {
		return oneFilePointer, nil
	} else {
		Logger.Error("Cannot produce log file")
		return nil, err
	}
}
