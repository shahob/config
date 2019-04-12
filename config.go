package config

import (
	"encoding/json"
	"os"
)

func Load(filePath string, configuration interface{}) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&configuration)
	return nil
}
