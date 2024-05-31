package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Read(filePath string) (*Config, error) {
	var cfg Config

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer func() {
		if cerr := f.Close(); cerr != nil {
			// Protokollierung des Fehlers oder eine andere geeignete Aktion
			fmt.Fprintf(os.Stderr, "failed to close file: %v\n", cerr)
		}
	}()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML: %w", err)
	}

	return &cfg, nil
}
