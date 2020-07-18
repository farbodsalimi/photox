package util

import (
	"fmt"
	"log"
	"os/user"
	"time"
)

// MakePathByTakenDateTime makes a path in string by using datetime from taken field
func MakePathByTakenDateTime(basePath string, taken time.Time) string {
	return fmt.Sprintf("%s/%d-%d-%d", basePath, taken.Year(), taken.Month(), taken.Day())
}

// GetHomeDir returns user home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}
