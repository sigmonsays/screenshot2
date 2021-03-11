package capture

import (
	"os"

	"github.com/sigmonsays/screenshot2/config"
	"github.com/urfave/cli/v2"
)

func getConfig(cfgfile string) (*config.AppConfig, error) {
	cfg := config.GetDefaultConfig()

	st, err := os.Stat(cfgfile)
	if err != nil || st.IsDir() {
		return cfg, nil
	}

	log.Tracef("Loading config %s...", cfgfile)
	err = cfg.LoadYaml(cfgfile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GalleryCapture(c *cli.Context) error {
	cfg, err := getConfig(c.String("config"))
	if err != nil {
		return err
	}

	shortname := NewShortname(c.String("shortname"))
	log.Tracef("GalleryCapture shortname:%s", shortname.Value)

	cap := &Command{}
	err = cap.Capture(cfg, shortname)
	if err != nil {
		return err
	}
	return nil
}
