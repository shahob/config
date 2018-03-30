package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Base     string `json:"database"`
	} `json:"database"`
	Trello struct {
		Key      string `json:"key"`
		Token    int    `json:"token"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"trello"`
}

func Load(file *string) Config {
	var config Config
	configFile, err := os.Open(*file)
	defer configFile.Close()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
