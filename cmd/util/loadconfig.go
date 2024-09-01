package util

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	DB struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Db    string `yaml:"db"`
		User  string `yaml:"user"`
		Pass  string `yaml:"pass"`
		Extra struct {
			SSL bool `yaml:"ssl"`
		} `yaml:"extra,omitempty"`
	} `yaml:"db"`
}

func LoadConfig() (*DbConfig, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("error while getting current directory: %v", err)
		os.Exit(1)
	}

	configPath := filepath.Join(dir, ".config", "default.yml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("error while reading config file: %v", err)
		os.Exit(1)
	}

	var config DbConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("error while unmarshalling config file: %v", err)
		os.Exit(1)
	}

	return &config, nil
}
