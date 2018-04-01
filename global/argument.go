package global

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type ArgumentsStruct struct {
	JobName       string `long:"job-name"`
	ConfigPath    string `long:"config-path" default:"../config/config.json"`
	LogConfigPath string `long:"log-config-path" default:"../config/logger.json"`
	NumOfExecutor int    `long:"num-of-executor" default:"10"`
	SpecificDate  string `long:"specific-date"`
}

var Arguments ArgumentsStruct = ArgumentsStruct{}

func init() {
	_, err := flags.ParseArgs(&Arguments, os.Args)

	if err != nil {
		panic(err)
	}
}
