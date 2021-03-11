package upload

import (
	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Upload struct {
}

func (me *Upload) Upload(cfg *config.AppConfig, shortname *core.Shortname) error {
	log.Tracef("starting upload shortname:%s", shortname)
	return nil
}
