package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Users       UsersConfiguration
	AutoWelcome bool
}
type UsersConfiguration struct {
	OwnerId     string
	Admins      []string
	BlackList   []string
	Permissions map[string][]string
}

var Config Configuration

func Load() {
	log.Println("[config][Load] Loading configuration file")
	configFile, _ := os.Open("config.json")
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln("[config][init] Unable to decode configuration:", err)
	}
	Config = configuration
}

func Save() {
	configFile, err := os.Create("config.json")
	defer configFile.Close()
	if err != nil {
		log.Fatalln("[config][Save] Unable to open configuration file:", err)
	}
	configFile.WriteString(ToJson())
}

func ToJson() string {
	data, _ := json.MarshalIndent(Config, "", "  ")
	return string(data)
}
