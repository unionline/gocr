package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"gocr/utils"
)

// Setting ...
var Setting Config

const (
	ConfigPath     = "config/config.toml"
	ConfigPath_win = "config/config_windows.toml"
)

// Init ...
func Init() {
	var configPath string
	if utils.IsLinuxPlatform() {
		configPath = ConfigPath
	} else {
		configPath = ConfigPath_win
	}

	if err := LoadConfigs(configPath); err != nil {
		panic(err)
	}

}

// LoadConfigs ...
func LoadConfigs(files ...string) error {
	for _, file := range files {
		if _, err := toml.DecodeFile(file, &Setting); err != nil {
			logrus.Error("decode error: ", err)
			return err
		}
	}

	return nil
}
