package core

import (
	gologging "github.com/sigmonsays/go-logging"
)

var log gologging.Logger

const LoggerName = "core"

func init() {
	log = gologging.Register(LoggerName, func(log gologging.Logger) {
		log = log
	})
}
