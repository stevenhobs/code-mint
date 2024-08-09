package config

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/stevenhobs4323/code-mint/internal/base"
)

type ConfData struct {
	AppendPath map[string]int    `json:"AppendPath"`
	OtherEnv   map[string]string `json:"OtherEnv"`
}

var config ConfData

func LoadData() {
	var home string = base.CODE_HOME
	confFile := path.Join(home, "Config", "code-mint.json")
	configData, err := os.ReadFile(confFile)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
}
