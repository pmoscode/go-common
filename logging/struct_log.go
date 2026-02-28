package logging

import (
	"fmt"
	strings2 "strings"
	"time"

	"github.com/pmoscode/go-common/strings"
)

// OutputFormat defines the format for structured log output.
type OutputFormat int

const (
	// Json formats the struct as JSON.
	Json OutputFormat = iota
	// Yaml formats the struct as YAML.
	Yaml
)

// InfoStruct logs a struct in the given format (JSON or YAML) with INFO severity.
func (l *Logger) InfoStruct(format OutputFormat, obj any) {
	var structString string
	switch format {
	case Json:
		structString = strings.PrettyPrintJson(obj)
	case Yaml:
		structString = strings.PrettyPrintYaml(obj)
	}

	header := l.header("INFO")
	dateStr := time.Now().Format("2006/01/02 15:04:05")

	// Prefix every line with timestamp + header
	lines := strings2.Split(structString, "\n")
	for _, line := range lines {
		formatted := fmt.Sprintf("%s %s%s\n", dateStr, header, line)
		_, _ = l.writer.Write([]byte(formatted))
	}
}
