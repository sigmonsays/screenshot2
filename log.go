package screenshot2

import (
	gologging "github.com/sigmonsays/go-logging"
)

var log gologging.Logger

const LoggerName = "screenshot2"

func init() {
	log = gologging.Register(LoggerName, func(log gologging.Logger) {
		log = log
	})
}
