package config

import (
	"encoding/json"
	"os"
)

type NameConfig struct {
	FirstNames struct {
		Male   []string `json:"male"`
		Female []string `json:"female"`
	} `json:"firstNames"`
	LastNames struct {
		Common []string `json:"common"`
	} `json:"lastNames"`
}

func FromFile(path string) (*NameConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg NameConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
