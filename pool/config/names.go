package config

import (
	"encoding/json"
	"os"
)

type FirstNames struct {
	Male   []string `json:"male"`
	Female []string `json:"female"`
}
type LastNames struct {
	Common []string `json:"common"`
}

type NameConfig struct {
	FirstNames FirstNames `json:"firstNames"`
	LastNames  LastNames  `json:"lastNames"`
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
