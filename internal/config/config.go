package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	res, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(res, &config)
	if err != nil {
		fmt.Printf("%v", err)
		return Config{}, err
	}
	return config, nil
}

func (c Config) SetUser(name string) {
	c.CurrentUserName = name
	write(c)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := homeDir + "/" + configFileName
	return path, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	writeable, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, writeable, 0644)
	return nil
}
