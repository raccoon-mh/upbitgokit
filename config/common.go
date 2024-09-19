package config

import (
	"log"
	"os"
	"os/user"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	configPathENV := os.Getenv("RACCOON_UPBIT_TRADER_CONFIG_PATH")
	if configPathENV != "" {
		configPath = configPathENV
	} else {
		configPath = userBaseDir() + "/.raccoon-upbit-trader"
	}
}

func userBaseDir() string {
	home, _ := os.UserHomeDir()
	if len(home) > 0 {
		return home
	}
	currUser, _ := user.Current()
	if currUser != nil {
		home = currUser.HomeDir
	}
	return home
}

func getAllConfig(filePath string, s any) any {
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Config file read Err: %v", err)
		return nil
	}
	err = viper.Unmarshal(&s)
	if err != nil {
		log.Printf("Config file unmarshal Err: %v", err)
		return nil
	}
	return s
}
