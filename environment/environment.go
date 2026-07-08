// Package environment implements functions to read values from environment variables of the OS.
package environment

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// GetEnv reads a string value from the environment with the given key.
// If the key is not found, the default value is returned.
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// getEnvParsed reads a string value from the environment with the given key and
// parses it using the given parse function. If the key is not found or parsing
// fails, the default value is returned (and the reason is logged).
func getEnvParsed[T any](key string, defaultVal T, parse func(string) (T, error)) T {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Nothing found for key '%s'. Using default value: '%v'\n", key, defaultVal)
		return defaultVal
	}

	parsed, err := parse(value)
	if err != nil {
		log.Printf("Could not parse value for key '%s'! Returning default value: '%v'\n", key, defaultVal)
		return defaultVal
	}

	return parsed
}

// GetEnvBool reads a string value from the environment with the given key and returns it as bool.
// If the key is not found or is not a bool, the default value is returned.
func GetEnvBool(key string, defaultVal bool) bool {
	return getEnvParsed(key, defaultVal, strconv.ParseBool)
}

// GetEnvInt reads a string value from the environment with the given key and returns it as int.
// If the key is not found or is not a int, the default value is returned.
func GetEnvInt(key string, defaultVal int) int {
	return getEnvParsed(key, defaultVal, strconv.Atoi)
}

// GetEnvFloat32 reads a string value from the environment with the given key and returns it as float32.
// If the key is not found or is not a float32, the default value is returned.
func GetEnvFloat32(key string, defaultVal float32) float32 {
	return getEnvParsed(key, defaultVal, func(s string) (float32, error) {
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	})
}

// GetEnvFloat64 reads a string value from the environment with the given key and returns it as float64.
// If the key is not found or is not a float64, the default value is returned.
func GetEnvFloat64(key string, defaultVal float64) float64 {
	return getEnvParsed(key, defaultVal, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}

// GetEnvMap reads all environment variables with the given prefix. (CAUTION!! Auto attaches the "_" character!!)
//
// It returns a map with all env variables found.
func GetEnvMap(prefix string, cutoffPrefix bool) map[string]string {
	envMap := make(map[string]string)
	allEnv := os.Environ()

	calcPrefix := prefix
	if !strings.HasSuffix(prefix, "_") {
		calcPrefix += "_"
	}

	for _, env := range allEnv {
		if rawKey, value, ok := cutEnvEntry(env, calcPrefix); ok {
			envMap[cleanKey(rawKey, calcPrefix, cutoffPrefix)] = value
		}
	}

	return envMap
}

// cutEnvEntry splits a raw "KEY=VALUE" environment entry and reports whether the
// key has the given prefix.
func cutEnvEntry(env, prefix string) (key, value string, ok bool) {
	key, value, _ = strings.Cut(env, "=")
	if !strings.HasPrefix(key, prefix) {
		return "", "", false
	}
	return key, value, true
}

func cleanKey(key, prefix string, cutoffPrefix bool) string {
	if cutoffPrefix {
		if trimmed, ok := strings.CutPrefix(key, prefix); ok {
			return trimmed
		}
	}

	return key
}
