// Package yamlconfig provides functions to load yaml configs from a file into a struct.
//
// Deprecated: Use the config package with [config/formats.ParseYamlConfig] instead.
package yamlconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadConfig loads a yaml config file into a given struct.
//
// Deprecated: Use [config/formats.ParseYamlConfig] instead.
func LoadConfig(filename string, out interface{}) error {
	yamlFileData, err := loadYaml(filename)
	if err != nil {
		return err
	}

	err = parseYaml(yamlFileData, out)
	if err != nil {
		return err
	}

	return nil
}

func loadYaml(filename string) ([]byte, error) {
	yamlFile, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	return yamlFile, nil
}

func parseYaml(yamlFileData []byte, out interface{}) error {
	err := yaml.Unmarshal(yamlFileData, out)
	if err != nil {
		return fmt.Errorf("error parsing YAML file: %w", err)
	}

	return nil
}
