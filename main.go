package main

import log "github.com/sirupsen/logrus"

// LogErr ... Logs an error if there is an error with logrus
func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
