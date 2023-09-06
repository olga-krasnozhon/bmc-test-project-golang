package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	UseCSV     bool `json:"UseCSV"`
	UseSQLite  bool `json:"UseSQLite"`
	ServerPort int  `json:"server_port"`
	Database   struct {
		Driver           string `json:"driver"`
		ConnectionString string `json:"connection_string"`
	} `json:"database"`
	APIKey    string `json:"api_key"`
	DebugMode bool   `json:"debug_mode"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
