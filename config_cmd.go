package screenshot2

import (
	"github.com/sigmonsays/screenshot2/core"
	"github.com/urfave/cli/v2"
)

func Config(c *cli.Context) error {
	cfg, err := GetConfig(c.String("config"))
	if err != nil {
		return err
	}

	shortname := core.NewShortname(c.String("shortname"))
	shortname.IncludeDate = true
	log.Tracef("GalleryCapture shortname:%s", shortname.GetShortname())

	deleteLocal := shortname.LocalFile != "" && cfg.Capture.KeepLocal == false
	log.Debugf("deleteLocal %v", deleteLocal)

	cfg.PrintConfig()
	return nil
}
