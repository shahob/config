package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(filePath string, configuration interface{}) interface{} {
	configFile, err := os.Open(filePath)
	defer configFile.Close()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&configuration)
	return configuration
}
