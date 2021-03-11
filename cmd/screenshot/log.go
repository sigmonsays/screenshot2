package main

import (
	gologging "github.com/sigmonsays/go-logging"
)

var log gologging.Logger

const LoggerName = "main"

func init() {
	log = gologging.Register(LoggerName, func(log gologging.Logger) {
		log = log
	})
}
