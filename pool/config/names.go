package config

import (
	"encoding/json"
	"os"
)

type Firstnames struct {
	Male   []string `json:"male"`
	Female []string `json:"female"`
}
type Lastnames struct {
	Common []string `json:"common"`
}

type NameConfig struct {
	Firstnames Firstnames `json:"firstnames"`
	Lastnames  Lastnames  `json:"lastnames"`
}

func FromFile(path string) (*NameConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FromBytes(data)
}

func FromBytes(data []byte) (*NameConfig, error) {
	var cfg NameConfig
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
