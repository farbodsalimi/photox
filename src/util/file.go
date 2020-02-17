package util

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"

	"path/filepath"
	"photox/src/config"
)

// CopyFile copies files using cp command
func CopyFile(srcFile string, destFile string) {
	cpCmd := exec.Command("cp", srcFile, destFile)
	err := cpCmd.Run()
	if err != nil {
		log.Fatal("copyFile: ", err)
	}
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.Replace(strings.ToUpper(filepath.Ext(path)), ".", "", 1)

		if config.SupportedExt[ext] {
			*files = append(*files, path)
		}

		return nil
	}
}

// ListSupportedFiles returns a list of files in the given directory
func ListSupportedFiles(root string, files *[]string) error {
	return filepath.Walk(root, visit(files))
}
