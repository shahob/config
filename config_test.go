package config

import (
	"testing"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	Base     string `json:"database"`
}

type Config struct {
	Database `json:"database"`
}

const DefaultConfigFilePath = "config.json"

func TestLoad(t *testing.T) {
	testConfig := Config{
		Database{
			"localhost",
			5432,
			"pwd",
			"usr",
			"db",
		},
	}

	configuration := Config{}

	err := Load(DefaultConfigFilePath, &configuration)

	if err != nil {
		t.Errorf(err.Error())
	}

	if configuration != testConfig {
		t.Errorf("Incorrect configuration file")
	}
}
