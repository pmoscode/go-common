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

		return fmt.Sprintf("%+v", obj)
	}

	return string(pretty)
}

// PrettyPrintYaml returns a YAML string representation of a given struct.
//
// Unlike encoding/json, gopkg.in/yaml.v3 panics (instead of returning an error)
// for genuinely unsupported types (e.g. channels or functions). This is
// recovered internally and falls back to the same "%+v" representation used
// for regular marshalling errors, so PrettyPrintYaml itself never panics.
func PrettyPrintYaml(obj any) (result string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered from panic while marshalling to YAML: %v", r)
			result = fmt.Sprintf("%+v", obj)
		}
	}()

	pretty, err := yaml.Marshal(obj)
	if err != nil {
		log.Println(err)

		return fmt.Sprintf("%+v", obj)
	}

	return strings2.TrimRight(string(pretty), "\n")
}
