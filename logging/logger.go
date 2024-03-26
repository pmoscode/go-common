package logging

import (
	"fmt"
	"github.com/pmoscode/go-common/shutdown"
	"io"
	"log"
)

type Level string

const (
	INFO    Level = "info"
	DEBUG   Level = "debug"
	WARNING Level = "warning"
	TRACE   Level = "trace"
)

type Logger struct {
	name            string
	nameSpacing     int
	severitySpacing int
	extend          string
	debug           bool
	trace           bool
	writer          *io.Writer
}

func (l *Logger) Info(format string, params ...any) {
	l.log("INFO", format, params...)
}

func (l *Logger) Warning(format string, params ...any) {
	l.log("WARNING", format, params...)
}

func (l *Logger) Debug(format string, params ...any) {
	if l.debug {
		l.log("DEBUG", format, params...)
	}
}

func (l *Logger) Trace(format string, params ...any) {
	if l.trace {
		l.log("TRACE", format, params...)
	}
}

func (l *Logger) Error(err error, format string, params ...any) {
	l.log("ERROR", format, params...)
	if err != nil {
		log.Println(err)
	}
}

func (l *Logger) Fatal(err error, format string, params ...any) {
	l.log("ERROR", format, params...)
	if err != nil {
		log.Println(err)
	}
	shutdown.Exit(1)
}

func (l *Logger) Panic(err error, format string, params ...any) {
	l.log("Panic", format, params...)
	panic(err)
}

func (l *Logger) log(severity string, format string, params ...any) {
	entry := fmt.Sprintf(format, params...)
	line := fmt.Sprintf("%s%s", l.header(severity), entry)

	_, err := (*l.writer).Write([]byte(line))
	if err != nil {
		return
	}
}

func (l *Logger) header(severity string) string {
	return fmt.Sprintf("%-*s [%*s] %s ### ", l.severitySpacing, severity, l.nameSpacing, l.name, l.extend)
}

func (l *Logger) IdDebug() bool {
	return l.debug
}

func (l *Logger) IdTrace() bool {
	return l.trace
}
