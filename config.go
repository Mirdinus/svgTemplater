package main

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigType struct {
	Weather struct {
		City    string `json:"city"`
		BaseUrl string `json:"baseUrl"`
		ApiKey  string `json:"apiKey"`
	} `json:"weather"`
	Todoist struct {
		BaseUrl   string `json:"baseUrl"`
		ApiKey    string `json:"apiKey"`
		ProjectID int    `json:"projectID"`
	} `json:"todoist"`
	Calendar struct {
		ICSUrl        string `json:"icsUrl"`
		LocalTimezone string `json:"localTimezone"`
		NamedayUrl    string `json:"namedayUrl"`
	} `json:"calendar"`
	Server struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"server"`
}

var Config ConfigType
var configFile = "config.json"

func loadConfig() {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		log.Fatalf("Config file not found: %v", err)
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}
}

func saveConfig() {
	file, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling config file: %v", err)
	}

	err = os.WriteFile(configFile, file, 0644)
	if err != nil {
		log.Fatalf("Error writing config file: %v", err)
	}

}
