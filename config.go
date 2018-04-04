package config

import (
	"encoding/json"
	"os"
)

func Load(filePath string, configuration interface{}) error {
	configFile, err := os.Open(filePath)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&configuration)
	return nil
}
