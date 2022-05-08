package util

import (
	"fmt"
	"log"
	"time"

	"github.com/arpanetus/kun/pkg/config"
	"github.com/kelvins/sunrisesunset"
)

// check if config file exists, if exists then read it,
// if it's empty or invalid then warn user and recreate it with default values
func Config() (*config.KunConfig, error) {
	log.Println("checking config file path")
	path, err := ConfigFilePath(DefaultConfigFileName)
	if err!=nil {
		return nil, fmt.Errorf("cannot get config file path: %s", err)
	}

	c := config.DefaultConfig()

	log.Println("checking config file exists")
	if ConfigFileExists(path) {
		log.Println("config file exists, reading it")
		data, err := ReadConfig(path)
		if err!=nil {
			log.Println("cannot read config file, creating new one with default values")
			data, err = c.Marshal()
			if err!=nil {
				return nil, fmt.Errorf("cannot create new config file: %s", err)
			}

			if err = WriteConfig(path, data); err!=nil {
				return nil, fmt.Errorf("cannot write new config file: %s", err)
			}
		}
		if err = c.Unmarshal(data); err!=nil {
			log.Println("cannot unmarshal config file, creating new one with default values")
			data, err = c.Marshal()
			if err!=nil {
				return nil, fmt.Errorf("cannot create new config file: %s", err)
			}
		}
	} else {
		log.Println("config file doesn't exist, creating new one with default values")
		data, err := c.Marshal()
		if err!=nil {
			return nil, fmt.Errorf("cannot create new config file: %s", err)
		}
		if err = WriteConfig(path, data); err!=nil {
			return nil, fmt.Errorf("cannot write new config file: %s", err)
		}
	}

	return c, nil
}


func sunriseSunset(config *config.KunConfig) (time.Time, time.Time, error) {
	// let's get the offset from	
	ofs, err := UTCOffset(config.Latitude, config.Longitude)
	if err!=nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot get UTC offset: %s", err)
	}

	sunrise, sunset, err := sunrisesunset.GetSunriseSunset(config.Latitude, config.Longitude, float64(ofs), time.Now())
	if err!=nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot get sunrise and sunset: %s", err)
	}

	return sunrise, sunset, nil
}

func SunriseSunset() (time.Time, time.Time, error) {
	c, err := Config()
	if err!=nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot get config: %s", err)
	}

	return sunriseSunset(c)
}


func IsSunDown() (bool, error) {
	log.Println("checking if sun is down")
	config, err := Config()
	if err!=nil {
		return false, fmt.Errorf("cannot get config: %s", err)
	}
	rise, set, err := sunriseSunset(config)
	now := time.Now()
	if now.After(set) {
		return true, nil
	}

	if now.Before(rise) {
		return true, nil
	}

	return false, nil
}