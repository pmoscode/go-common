// Package config can be used to load data to a struct which is defined by the "env" tag.
//
// Data may come from a file and/or the os environment
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pmoscode/go-common/config/formats"
	"github.com/pmoscode/go-common/config/meta"
	"gopkg.in/yaml.v3"
	"reflect"
	"strconv"
)

// Format defines the currently possible config file format (YAML or JSON)
type Format int

const (
	YAML Format = iota
	JSON
)

var functions = []formats.Parser{yaml.Unmarshal, json.Unmarshal}

// LoadFromFile loads config from a file with the given format parser.
//
//   - filename: Path + filename of the config file
//   - target: target struct to load the config to (MUST be a pointer to a struct)
//   - format: format type (YAML or JSON)
func LoadFromFile(filename string, target any, format Format) error {
	err := formats.ParseConfig(filename, target, functions[format])
	if err != nil {
		return err
	}

	return LoadFromEnvironment(target)
}

// LoadFromEnvironment loads environment data into the target struct.
//
// target must be a pointer, otherwise it will return an error.
func LoadFromEnvironment(target any) error {
	pointer := reflect.TypeOf(target)

	if pointer.Kind() != reflect.Pointer || target == nil {
		return fmt.Errorf("target must be a valid pointer")
	}

	typeOf := pointer.Elem()
	valueOf := reflect.ValueOf(target).Elem()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		if !field.IsExported() {
			continue
		}

		err := resolveField(field, valueOf, i)
		if err != nil {
			return err
		}
	}

	return nil
}

func resolveField(field reflect.StructField, valueOf reflect.Value, fieldIndex int) error {
	metaTag := meta.NewTagMeta()

	hasTag, err := metaTag.Parse(field)
	if err != nil {
		return err
	}

	if hasTag {
		element := valueOf.Field(fieldIndex)
		fieldType := field.Type

		switch metaTag.Kind() {
		case meta.Name:
			return resolveNameTag(fieldType, element, metaTag)
		case meta.Prefix:
			err = resolvePrefixTag(field, element, metaTag)
			if err != nil {
				return err
			}
		case meta.None:
			return errors.New("misconfigured tag for field: " + field.Name)
		}
	}

	return nil
}

func resolveNameTag(fieldType reflect.Type, element reflect.Value, metaTag *meta.Tag) error {
	switch fieldType.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		element.SetInt(metaTag.ValueAsInt())
	case reflect.String:
		element.SetString(metaTag.ValueAsString())
	case reflect.Bool:
		element.SetBool(metaTag.ValueAsBool())
	default:
		return errors.New("Invalid field type: " + fieldType.String())
	}

	return nil
}

func resolvePrefixTag(field reflect.StructField, element reflect.Value, metaTag *meta.Tag) error {
	if element.Kind() != reflect.Map {
		return fmt.Errorf("field '%s is not a map", field.Name)
	}

	valueType := element.Type()
	keyType := valueType.Key()

	if element.IsNil() {
		element.Set(reflect.MakeMap(valueType))
	}

	envs := metaTag.ValueAsMap()
	for key, value := range envs {
		err := resolvePrefixTagEnvItem(key, value, keyType, valueType, element)
		if err != nil {
			return err
		}
	}

	return nil
}

func resolvePrefixTagEnvItem(key string, value string, keyType reflect.Type, valueType reflect.Type, element reflect.Value) error {
	keyValue, valueValue := reflect.ValueOf(key), reflect.ValueOf(value)
	if !keyValue.Type().ConvertibleTo(keyType) {
		return fmt.Errorf("can't convert key to type %v", keyType.Kind())
	}

	keyValue = keyValue.Convert(keyType)
	valueElement := valueType.Elem()

	if !valueValue.Type().ConvertibleTo(valueElement) {
		switch valueElement.Kind() {
		case reflect.Int:
			valueValue = getPrimitive(strconv.Atoi(value))
		case reflect.Int32:
			valueValue = getPrimitive(strconv.ParseInt(value, 10, 32))
		case reflect.Int64:
			valueValue = getPrimitive(strconv.ParseInt(value, 10, 64))
		case reflect.Float32:
			valueValue = getPrimitive(strconv.ParseFloat(value, 32))
		case reflect.Float64:
			valueValue = getPrimitive(strconv.ParseFloat(value, 64))
		default:
			return fmt.Errorf("can't assign value to type %v", keyType.Kind())
		}
	}

	element.SetMapIndex(keyValue, valueValue.Convert(valueElement))
	return nil
}

func getPrimitive(val any, err error) reflect.Value {
	if err != nil {
		panic(err)
	}

	return reflect.ValueOf(val)
}
