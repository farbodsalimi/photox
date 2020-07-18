package photox

import (
	"fmt"
	"os"
	"path"

	"github.com/rwcarlsen/goexif/exif"
	log "github.com/sirupsen/logrus"

	"photox/pkg/config"
	"photox/pkg/util"
)

// Run run photox
func Run(fromPath string, toPath string) {
	dm := util.DirectoryManager{CachedDirectories: make(map[string]bool)}

	// Prepare a folder for undefined photos
	undefinedPath := path.Join(toPath, util.GetHomeDir(), config.BasePath, config.UndefinedPath)
	dm.PrepareDirectory(undefinedPath)

	// List all the supported files
	var files []string
	err := util.ListSupportedFiles(fromPath, &files)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Processing %d files ...", len(files))

	for index, file := range files {
		fmt.Printf("\n↻ Processing file: %d %s\n", index, file)

		// Read file
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		// Decode exif
		x, err := exif.Decode(f)
		if err != nil {
			log.Warnf("⚠︎ Unable to read exchangeable image file format!")
			log.Debugf("exif.Decode: %v", err)
			util.CopyFile(file, undefinedPath)
			continue
		}

		// Extract taken field from exif
		taken, err := x.DateTime()
		if err != nil {
			log.Warnf("⚠︎ Taken field not found!")
			log.Debugf("x.DateTime: %v", err)
			util.CopyFile(file, undefinedPath)
			continue
		}

		// Make a path and directory for the taken datetime
		p := util.MakePathByTakenDateTime(path.Join(toPath, config.BasePath), taken)
		if !dm.IsCachedDirectory(p) {
			fmt.Printf("❐ Making directory: %s\n", p)
			dm.PrepareDirectory(p)
		}

		// Copy the file into the taken directory
		util.CopyFile(file, p)
		fmt.Printf("✓ %d File %s is done!\n", index, file)
	}
}
