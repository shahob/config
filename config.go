package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	Base     string `json:"database"`
}

type Trello struct {
	Key            string `json:"key"`
	Token          string `json:"token"`
	ListTesting    string `json:"listTesting"`
	ListInProgress string `json:"listInProgress"`
}

type Gitlab struct {
	Token     string `json:"token"`
	ProjectId int    `json:"projectId"`
}

type Config struct {
	Database `json:"database"`
	Trello   `json:"trello"`
	Gitlab   `json:"gitlab"`
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
