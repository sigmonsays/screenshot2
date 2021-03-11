package upload

import (
	gologging "github.com/sigmonsays/go-logging"
)

var log gologging.Logger

const LoggerName = "upload"

func init() {
	log = gologging.Register(LoggerName, func(log gologging.Logger) {
		log = log
	})
}
