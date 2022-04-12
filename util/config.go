package util

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	var configFile string
	flag.StringVar(&configFile, "config.file", "", "config file path")
	flag.Parse()
	if configFile == "" {
		curPath, err := ExecPath()
		if err != nil {
			log.Panicf("Get current path error: %w", err)
		}
		configFile = curPath + string(os.PathSeparator) + "config/common.toml"
	}
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Fatal error config file: %w", err)
	}
}
