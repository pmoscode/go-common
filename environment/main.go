// Package environment implements functions to read values from environment variables of the OS.
package environment

import (
	"fmt"
	"os"
	"strconv"
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
