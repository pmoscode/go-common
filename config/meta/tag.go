package meta

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pmoscode/go-common/environment"
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

// Tag holds the parsed metadata from an "env" struct tag.
// Cardinality: name, self, prefix
// defaultValue only for name and self
// cutoff only for prefix
type Tag struct {
	name         string
	defaultValue string
	prefix       string
	cutoff       bool
}

// Parse extracts the "env" tag from the given struct field and populates the Tag fields.
// Returns true if a tag was found, false otherwise.
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

// Kind returns the kind of the parsed tag (Name, Prefix, or None).
func (m *Tag) Kind() Kind {
	if m.name != "" {
		return Name
	}

	if m.prefix != "" {
		return Prefix
	}

	return None
}

// ValueAsInt returns the environment variable value as int64.
func (m *Tag) ValueAsInt() (int64, error) {
	defaultVal, err := m.defaultAsInt()
	if err != nil {
		return 0, err
	}
	return int64(environment.GetEnvInt(m.name, defaultVal)), nil
}

// ValueAsString returns the environment variable value as string.
func (m *Tag) ValueAsString() string {
	return environment.GetEnv(m.name, m.defaultValue)
}

// ValueAsBool returns the environment variable value as bool.
func (m *Tag) ValueAsBool() (bool, error) {
	defaultVal, err := m.defaultAsBool()
	if err != nil {
		return false, err
	}
	return environment.GetEnvBool(m.name, defaultVal), nil
}

// ValueAsMap returns the environment variables matching the prefix as a map.
func (m *Tag) ValueAsMap() map[string]string {
	return environment.GetEnvMap(m.prefix, m.cutoff)
}

func (m *Tag) defaultAsInt() (int, error) {
	if m.defaultValue == "" {
		return 0, nil
	}

	parseInt, err := strconv.Atoi(m.defaultValue)
	if err != nil {
		return 0, fmt.Errorf("could not parse default value '%s' as int: %w", m.defaultValue, err)
	}

	return parseInt, nil
}

func (m *Tag) defaultAsBool() (bool, error) {
	if m.defaultValue == "" {
		return false, nil
	}

	parseBool, err := strconv.ParseBool(m.defaultValue)
	if err != nil {
		return false, fmt.Errorf("could not parse default value '%s' as bool: %w", m.defaultValue, err)
	}

	return parseBool, nil
}

// NewTagMeta creates a new Tag instance with default values.
func NewTagMeta() *Tag {
	return &Tag{
		name:         "",
		defaultValue: "",
		prefix:       "",
		cutoff:       false,
	}
}
