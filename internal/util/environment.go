package util

import (
	"os"
	"strconv"
)

// GetStringOrDefault check if a env variable exists for the given name. If so
// it returns the string value of the env var. If not it will return the given default value.
func GetStringOrDefault(name, defaultV string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}

	return v
}

// GetIntOrDefault check if a env variable exists for the given name. If so
// it tries to convert the string into an integer value, if the system can do that it will return the converted
// integer value if not the system will return the given default integer value
func GetIntOrDefault(name string, defaultV int) int {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}

	vAsInt, err := strconv.Atoi(v)
	if err != nil {
		return defaultV
	}

	return vAsInt
}
