package strings

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	strings2 "strings"
)

func PrettyPrintJson(obj any) string {
	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println(err)

		return fmt.Sprintf("%+v\n", obj)
	}

	return string(pretty)
}
func PrettyPrintYaml(obj any) string {
	pretty, err := yaml.Marshal(obj)
	if err != nil {
		log.Println(err)

		return fmt.Sprintf("%+v\n", obj)
	}

	return strings2.TrimRight(string(pretty), "\n")
}
