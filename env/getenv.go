package env

import (
	"os"
)

func GetEnv(name string) string {
	result := os.Getenv(name)
	if result == "" {
		panic(result)
	}
	return result
}
