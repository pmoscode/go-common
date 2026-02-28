package logging

import (
	strings2 "strings"

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
	switch format {
	case Json:
		l.Info(l.addHeader("INFO", strings.PrettyPrintJson(obj)), "")
	case Yaml:
		l.Info(l.addHeader("INFO", strings.PrettyPrintYaml(obj)), "")
	}
}

func (l *Logger) addHeader(severity string, structString string) string {
	header := l.header(severity)

	return strings2.ReplaceAll(structString, "\n", "\n"+header)
}
