package config

import (
	"encoding/json"
	"os"

	"resume-website/internal/models"
)

func LoadConfig(path string) (*models.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg models.Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}
