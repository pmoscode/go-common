package meta

import (
	"github.com/iancoleman/strcase"
	"github.com/pmoscode/go-common/environment"
	"reflect"
	"strconv"
	"strings"
)

type Kind int

const (
	nameTag    = "name"
	selfTag    = "self"
	prefixTag  = "prefix"
	cutoffTag  = "cutoff"
	defaultTag = "default"
)

const (
	Name = iota
	Prefix
	None
)

const tagName = "env"

// Cardinality: name, self, prefix
// defaultValue only for name and self
// cutoff only for prefix

type Tag struct {
	name         string
	defaultValue string
	prefix       string
	cutoff       bool
}

func (m *Tag) Parse(field reflect.StructField) (bool, error) {
	tag, ok := field.Tag.Lookup(tagName)

	if !ok {
		return false, nil
	} else {
		if tag == "" {
			tag = selfTag
		}

		for _, val := range strings.Split(tag, ",") {
			items := strings.Split(val, "=")

			switch items[0] {
			case nameTag:
				m.name = items[1]
			case defaultTag:
				m.defaultValue = items[1]
			case selfTag:
				if m.name == "" {
					m.name = strcase.ToScreamingSnake(field.Name)
				}
			case prefixTag:
				m.prefix = items[1]
			case cutoffTag:
				if len(items) == 2 {
					result, err := strconv.ParseBool(items[1])
					if err != nil {
						return true, err
					}
					m.cutoff = result
				} else {
					m.cutoff = true
				}
			}
		}
	}

	return true, nil
}

func (m *Tag) Kind() Kind {
	if m.name != "" {
		return Name
	}

	if m.prefix != "" {
		return Prefix
	}

	return None
}

func (m *Tag) ValueAsInt() int64 {
	return int64(environment.GetEnvInt(m.name, m.defaultAsInt()))
}

func (m *Tag) ValueAsString() string {
	return environment.GetEnv(m.name, m.defaultValue)
}

func (m *Tag) ValueAsBool() bool {
	return environment.GetEnvBool(m.name, m.defaultAsBool())
}

func (m *Tag) ValueAsMap() map[string]string {
	return environment.GetEnvMap(m.prefix, m.cutoff)
}

func (m *Tag) defaultAsInt() int {
	if m.defaultValue == "" {
		return 0
	}

	parseInt, err := strconv.Atoi(m.defaultValue)
	if err != nil {
		panic(err)
	}

	return parseInt
}

func (m *Tag) defaultAsBool() bool {
	if m.defaultValue == "" {
		return false
	}

	parseBool, err := strconv.ParseBool(m.defaultValue)
	if err != nil {
		panic(err)
	}

	return parseBool
}

func NewTagMeta() *Tag {
	return &Tag{
		name:         "",
		defaultValue: "",
		prefix:       "",
		cutoff:       false,
	}
}
