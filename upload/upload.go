package upload

import (
	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Upload struct {
}

func (me *Upload) Upload(cfg *config.AppConfig, shortname *core.Shortname) error {
	log.Tracef("starting upload shortname:%s", shortname)

	//todo: pick method based on config
	//
	cmd := &Command{}
	err := cmd.Upload(cfg, shortname)
	if err != nil {
		return err
	}

	return nil
}
