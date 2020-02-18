package util

import (
	"log"
	"os"
)

// DirectoryManager manages everything related to directories
type DirectoryManager struct {
	realOS            RealOS
	CachedDirectories map[string]bool
}

// PrepareDirectory makes sure the directory exists
func (dm *DirectoryManager) PrepareDirectory(path string) {
	err := dm.realOS.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal("prepareFolder: ", err)
	} else if dm.CachedDirectories != nil {
		dm.CachedDirectories[path] = true
	}
}

func (dm *DirectoryManager) IsCachedDirectory(key string) bool {
	return dm.CachedDirectories[key]
}
