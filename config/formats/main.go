package formats

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Parser func(in []byte, out any) error

func parseConfig(filename string, out any, parser Parser) error {
	fileData, err := loadConfigFile(filename)
	if err != nil {
		return err
	}

	return parser(fileData, out)
}

func loadConfigFile(filename string) ([]byte, error) {
	configFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %s\n", err)
	}

	return configFile, nil
}

func ParseYamlConfig(filename string, out any) error {
	return parseConfig(filename, out, yaml.Unmarshal)
}

func ParseJsonConfig(filename string, out any) error {
	return parseConfig(filename, out, json.Unmarshal)
}

func ParseConfig(filename string, out any, parser Parser) error {
	return parseConfig(filename, out, parser)
}
