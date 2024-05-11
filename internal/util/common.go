package util

import "os"

// EnvOrDefault returns the value of the environment variable named by the key,
// or default value def.
func EnvOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
