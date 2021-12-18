package screenshot2

import (
	"github.com/sigmonsays/screenshot2/capture"
	"github.com/sigmonsays/screenshot2/clipboard"
	"github.com/sigmonsays/screenshot2/core"
	"github.com/sigmonsays/screenshot2/upload"
	"github.com/urfave/cli/v2"
)

func Capture(c *cli.Context) error {
	cfg, err := GetConfig(c.String("config"))
	if err != nil {
		return err
	}

	shortname := core.NewShortname(c.String("shortname"))
	shortname.IncludeDate = true
	log.Tracef("GalleryCapture shortname:%s", shortname.GetShortname())

	cap := &capture.Capture{}
	err = cap.Capture(cfg, shortname)
	if err != nil {
		return err
	}

	up := &upload.Upload{}
	err = up.Upload(cfg, shortname)
	if err != nil {
		return err
	}

	clip := clipboard.Clipboard{}
	err = clip.CopyToClipboard(cfg, shortname)
	if err != nil {
		return err
	}
	return nil
}
