package clipboard

import (
	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Clipboard struct {
}

type Clipper interface {
	CopyToClipboard(fg *config.AppConfig, shortname *core.Shortname) error
}

func (me *Clipboard) CopyToClipboard(cfg *config.AppConfig, shortname *core.Shortname) error {
	log.Tracef("starting clipboard copy shortname:%s", shortname)

	if shortname.Url == "" {
		log.Tracef("No url found, not copying to clipboard")
		return nil
	}

	var clip Clipper
	if cfg.Clipboard.Interface == "xclip" {
		cl := &XClip{}
		clip = cl
	} else {
		cl := &XClip{}
		clip = cl
	}
	err := clip.CopyToClipboard(cfg, shortname)
	if err != nil {
		return err
	}

	return nil
}
