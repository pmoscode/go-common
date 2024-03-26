// Package strings provides functions to operate with strings or where strings are a result of.
package strings

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	strings2 "strings"
)

// PrettyPrintJson returns a JSON string representation of a given struct.
func PrettyPrintJson(obj any) string {
	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println(err)

		return fmt.Sprintf("%+v\n", obj)
	}

	return string(pretty)
}

// PrettyPrintYaml returns a YAML string representation of a given struct.
func PrettyPrintYaml(obj any) string {
	pretty, err := yaml.Marshal(obj)
	if err != nil {
		log.Println(err)

		return fmt.Sprintf("%+v\n", obj)
	}

	return strings2.TrimRight(string(pretty), "\n")
}
