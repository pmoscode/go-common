package logging

import (
	"github.com/pmoscode/go-common/strings"
	strings2 "strings"
)

type OutputFormat int

const (
	Json OutputFormat = iota
	Yaml
)

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
