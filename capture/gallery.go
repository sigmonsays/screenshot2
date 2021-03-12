package capture

import (
	"github.com/sigmonsays/screenshot2/clipboard"
	"github.com/sigmonsays/screenshot2/core"
	"github.com/sigmonsays/screenshot2/upload"
	"github.com/urfave/cli/v2"
)

func GalleryCapture(c *cli.Context) error {
	cfg, err := getConfig(c.String("config"))
	if err != nil {
		return err
	}

	shortname := core.NewShortname(c.String("shortname"))
	log.Tracef("GalleryCapture shortname:%s", shortname.Value)

	cap := &Capture{}
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
