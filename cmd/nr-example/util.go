package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// mustGetEnv forces an environment var to be set, or panic
func mustGetEnv(key string) string {
	val := os.Getenv(key)

	if val != "" {
		return val
	}

	panic(fmt.Sprintf("environment variable %s unset", key))
}

// logIfError does just that.
func logIfError(err error) {
	if err != nil {
		log.Errorf("Caught error: %v", err)
	}
}
