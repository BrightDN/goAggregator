package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL 				string `json:"db_url"`
	CurrentUserName  	string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

    file, err := os.Open(path)
    if err != nil {
        return Config{}, err
    }
    defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
    	return Config{}, err
    }

    return cfg, nil
}

func (c *Config) SetUser(name string) error {
    c.CurrentUserName = name

    path, err := getConfigFilePath()
    if err != nil {
        return err
    }

    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(c)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFileName), nil
}