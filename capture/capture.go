package capture

import (
	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Capture struct {
}

type Capper interface {
	Capture(fg *config.AppConfig, shortname *core.Shortname) error
}

func (me *Capture) Capture(cfg *config.AppConfig, shortname *core.Shortname) error {
	log.Tracef("starting capture shortname:%s", shortname)
	var capper Capper
	iface := cfg.Capture.Interface

	if iface == "command" {
		cmd := &Command{}
		capper = cmd
	} else {
		capper = &Command{}
		log.Warnf("defaulting to command interface")
	}
	err := capper.Capture(cfg, shortname)
	if err != nil {
		return err
	}
	return nil
}
