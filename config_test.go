package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	testConfig := Config{
		Database{
			"localhost",
			5432,
			"pwd",
			"usr",
			"db",
		},
		Trello{
			"key",
			"token",
			"id",
			"id",
		},
		Gitlab{
			"token",
			1,
		},
	}

	filePath := "config.json"
	config := Load(&filePath)

	if config != testConfig {
		t.Errorf("Error")
	}
}