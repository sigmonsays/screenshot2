package upload

import (
	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Upload struct {
}

type Uploader interface {
	Upload(fg *config.AppConfig, shortname *core.Shortname) error
}

func (me *Upload) Upload(cfg *config.AppConfig, shortname *core.Shortname) error {
	log.Tracef("starting upload shortname:%s", shortname)

	var upper Uploader
	iface := cfg.Upload.Interface

	if iface == "command" {
		cmd := &Command{}
		upper = cmd
	} else if iface == "s3" {
		cmd := &S3{}
		upper = cmd
	} else {
		upper = &Command{}
		log.Warnf("defaulting to command interface")
	}
	err := upper.Upload(cfg, shortname)
	if err != nil {
		return err
	}

	return nil
}
