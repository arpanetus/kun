package util

import (
	"fmt"
	"log"

	"github.com/arpanetus/kun/pkg/config"
)

// check if config file exists, if exists then read it,
// if it's empty or invalid then warn user and recreate it with default values
func Config(logger *log.Logger) (*config.KunConfig, error) {
	logger.Println("checking config file path")

	pc := NewPathConfig(logger)
	path, err := pc.ConfigFilePath(DefaultConfigFileName)
	if err != nil {
		return nil, fmt.Errorf("cannot get config file path: %s", err)
	}

	c := config.DefaultConfig()

	logger.Println("checking config file exists")
	if pc.ConfigFileExists(path) {
		logger.Println("config file exists, reading it")
		data, err := pc.ReadConfig(path)
		if err != nil {
			logger.Println("cannot read config file, creating new one with default values")
			data, err = c.Marshal()
			if err != nil {
				return nil, fmt.Errorf("cannot create new config file: %s", err)
			}

			if err = pc.WriteConfig(path, data); err != nil {
				return nil, fmt.Errorf("cannot write new config file: %s", err)
			}
		}
		if err = c.Unmarshal(data); err != nil {
			logger.Println("cannot unmarshal config file, creating new one with default values")
			data, err = c.Marshal()
			if err != nil {
				return nil, fmt.Errorf("cannot create new config file: %s", err)
			}
		}
	} else {
		logger.Println("config file doesn't exist, creating new one with default values")
		data, err := c.Marshal()
		if err != nil {
			return nil, fmt.Errorf("cannot create new config file: %s", err)
		}
		if err = pc.WriteConfig(path, data); err != nil {
			return nil, fmt.Errorf("cannot write new config file: %s", err)
		}
	}

	return c, nil
}
