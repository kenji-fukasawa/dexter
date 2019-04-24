package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/kenji-fukasawa/dexter/cmd"
)

func main() {
	// set log format & level
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.InfoLevel)

	cmd.Execute()
}
