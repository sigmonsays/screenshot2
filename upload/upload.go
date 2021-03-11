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

	//todo: pick method based on config
	//
	method := "command"
	method = "s3"
	var upper Uploader
	if method == "command" {
		cmd := &Command{}
		upper = cmd
	} else if method == "s3" {
		cmd := &S3{}
		upper = cmd
	} else {
		upper = &Command{}
	}
	err := upper.Upload(cfg, shortname)
	if err != nil {
		return err
	}

	return nil
}
