package util

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	curPath, err := ExecPath()
	if err != nil {
		log.Panicf("Get current path error: %w\n", err)
	}
	viper.AddConfigPath(curPath + string(os.PathSeparator) + "config")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Fatal error config file: %w\n", err)
	}
}
