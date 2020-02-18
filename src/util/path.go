package util

import (
	"fmt"
	"time"
)

// MakePathByTakenDateTime makes a path in string by using datetime from taken field
func MakePathByTakenDateTime(basePath string, taken time.Time) string {
	return fmt.Sprintf("%s/%d-%d-%d", basePath, taken.Year(), taken.Month(), taken.Day())
}
