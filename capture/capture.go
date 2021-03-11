package capture

import (
	"github.com/sigmonsays/screenshot2/config"
	"github.com/urfave/cli/v2"
)

func getConfig(cfgfile string) *config.AppConfig {
	cfg := config.GetDefaultConfig()
	//cfg.LoadYaml(cfgfile)
	return cfg
}

func GalleryCapture(c *cli.Context) error {
	cfg := getConfig(c.String("config"))

	shortname := NewShortname(c.String("shortname"))
	log.Tracef("GalleryCapture shortname:%s", shortname.Value)

	cap := &Command{}
	err := cap.Capture(cfg, shortname)
	if err != nil {
		return err
	}
	return nil
}
