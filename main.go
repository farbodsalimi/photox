package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"photox/src/cli"
)

func init() {
	if _, exists := os.LookupEnv("PHOTOX_DEBUG"); exists {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	cli.Parse()
}
