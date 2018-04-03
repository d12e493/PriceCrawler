package global

import (
	"fmt"

	configor "github.com/jinzhu/configor"
)

type DatabaseConfig struct {
	Host     string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Port     int    `required:"true"`
	Dbname   string `required:"true"`
}

type ApiConfig struct {
	Host string `required:"true"`
	Port int    `required:"true"`
}

type ConfigStruct struct {
	Database DatabaseStruct
	Api      ApiConfig
}

type DatabaseStruct struct {
	Demo DatabaseConfig
}

var Config ConfigStruct = ConfigStruct{}

func init() {
	ReloadConfig()
}

func ReloadConfig() {

	var configPath = Arguments.ConfigPath

	err := configor.Load(&Config, configPath)

	if err != nil {
		fmt.Println(err)
	}
}
