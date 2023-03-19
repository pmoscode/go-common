package yamlconfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

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
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}

	return yamlFile, nil
}

func parseYaml(yamlFileData []byte, out interface{}) error {
	err := yaml.Unmarshal(yamlFileData, out)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return err
	}

	return nil
}
