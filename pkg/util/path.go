package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	defaultConfigDirName = ".config"
	defaultKunConfigDirName = "kun"
	DefaultConfigFileName = "kun.yaml"
)


func ConfigDirPath(path string) (string, error) { 
	// get current working directory 
	path, err := os.UserHomeDir()
	if err!=nil {
		return "", fmt.Errorf("cannot get user home directory: %s", err)
	}

	// append config directory name 
	path = filepath.Join(path, defaultConfigDirName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("creating config directory")
		
		//if not exists, create it
		if err = os.Mkdir(path, 0755); err!=nil {
			return "", fmt.Errorf("cannot create config directory: %s", err)
		}
	}
	
	// create kun config directory
	path = filepath.Join(path, defaultKunConfigDirName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("creating kun config directory")
		
		//if not exists, create it
		if err = os.Mkdir(path, 0755); err!=nil {
			return "", fmt.Errorf("cannot create kun config directory: %s", err)
		}
	}

	return path, nil
}

func ConfigFileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ConfigFilePath(path string) (string, error) {
	path, err := ConfigDirPath(path)
	if err!=nil {
		return "", fmt.Errorf("cannot get config directory path: %s", err)
	}
	path = filepath.Join(path, DefaultConfigFileName)
	return path, nil
}


func WriteConfig(path string, configData []byte) error {
	var file *os.File
	var err error
	file, err = os.Create(path);
	if err!=nil {
		return fmt.Errorf("cannot create config file: %s", err)
	}
	defer file.Close()

	file.Write(configData)
	log.Println("default values are set")

	return nil
}

func ReadConfig(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err!=nil {
		return []byte{}, fmt.Errorf("cannot open config file: %s", err)
	}
	defer file.Close()
	var buffer []byte
	_, err = file.Read(buffer)
	if err!=nil {
		return []byte{}, fmt.Errorf("cannot read config file: %s", err)
	}
	return buffer, nil
}