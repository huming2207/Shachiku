package common

import (
	"gopkg.in/ini.v1"
	"log"
)

var cfgFile *ini.File

func GetConfig() *ini.File {
	if cfgFile != nil {
		return cfgFile
	} else {
		var err error
		cfgFile, err = ini.Load("config.ini")
		if err != nil {
			log.Fatalln("Failed to open configuration file: %w", err)
		}
	}

	return cfgFile
}
