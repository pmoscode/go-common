package strings

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(obj any) string {
	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println(err)

		return fmt.Sprintf("%+v\n", obj)
	}

	return string(pretty)
}
