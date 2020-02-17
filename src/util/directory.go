package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

// PrepareDirectory makes sure the directory exists
func PrepareDirectory(directories map[string]bool, path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal("prepareFolder: ", err)
	} else if directories != nil {
		directories[path] = true
	}
}

// MakePathByTakenDateTime makes a path in string by using datetime from taken field
func MakePathByTakenDateTime(basePath string, taken time.Time) string {
	return fmt.Sprintf("%s/%d-%d-%d", basePath, taken.Year(), taken.Month(), taken.Day())
}
