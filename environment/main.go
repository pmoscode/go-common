// Package environment implements functions to read values from environment variables of the OS.
package environment

import (
	"fmt"
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

// GetEnvBool reads a string value from the environment with the given key and returns it as bool.
// If the key is not found or is not a bool, the default value is returned.
func GetEnvBool(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		} else {
			fmt.Printf("Could not parse '%s' for key '%s' as bool! Returning default value: '%t'\n", value, key, defaultVal)
		}
	} else {
		fmt.Printf("Nothing found for key '%s'. Using default value: '%t'\n", key, defaultVal)
	}

	return defaultVal
}

// GetEnvInt reads a string value from the environment with the given key and returns it as int.
// If the key is not found or is not a int, the default value is returned.
func GetEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		} else {
			fmt.Printf("Could not parse '%s' for key '%s' as int! Returning default value: '%d'\n", value, key, defaultVal)
		}
	} else {
		fmt.Printf("Nothing found for key '%s'. Using default value: '%d'\n", key, defaultVal)
	}

	return defaultVal
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
		if strings.HasPrefix(env, calcPrefix) {
			split := strings.Split(env, "=")
			key := cleanKey(split[0], calcPrefix, cutoffPrefix)
			value := split[1]

			envMap[key] = value
		}
	}

	return envMap
}

func cleanKey(key, prefix string, cutoffPrefix bool) string {
	if cutoffPrefix {
		key = strings.TrimLeft(key, prefix)
	}

	return key
}
