package logging

import (
	"fmt"
	"strings"
	"time"

	commonstrings "github.com/pmoscode/go-common/strings"
)

// OutputFormat defines the format for structured log output.
type OutputFormat int

const (
	// Json formats the struct as JSON.
	Json OutputFormat = iota
	// Yaml formats the struct as YAML.
	Yaml
)

// String implements fmt.Stringer for OutputFormat.
func (f OutputFormat) String() string {
	switch f {
	case Json:
		return "Json"
	case Yaml:
		return "Yaml"
	default:
		return fmt.Sprintf("OutputFormat(%d)", int(f))
	}
}

// InfoStruct logs a struct in the given format (JSON or YAML) with INFO severity.
func (l *Logger) InfoStruct(format OutputFormat, obj any) {
	var structString string
	switch format {
	case Json:
		structString = commonstrings.PrettyPrintJson(obj)
	case Yaml:
		structString = commonstrings.PrettyPrintYaml(obj)
	default:
		structString = fmt.Sprintf("unsupported OutputFormat: %s", format)
	}

	header := l.header("INFO")
	dateStr := time.Now().Format("2006/01/02 15:04:05")

	// Prefix every line with timestamp + header, then write everything in a
	// single Write call so concurrent log lines from other calls cannot get
	// interleaved with this one.
	var sb strings.Builder
	for _, line := range strings.Split(structString, "\n") {
		sb.WriteString(dateStr)
		sb.WriteByte(' ')
		sb.WriteString(header)
		sb.WriteString(line)
		sb.WriteByte('\n')
	}

	l.writeLine(sb.String())
}
