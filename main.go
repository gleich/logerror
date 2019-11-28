package main

import (
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

// Log ... Logs an error if there is an error with logrus
func Log(err error, msg string) {
	if err != nil {
		if msg != "" {
			color.Red(msg)
		}
		log.Fatal(err)
	}
}
