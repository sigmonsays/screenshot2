package screenshot2

import (
	"os"

	"github.com/sigmonsays/screenshot2/config"
)

func GetConfig(cfgfile string) (*config.AppConfig, error) {
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

	os.MkdirAll(cfg.DataDir, 0755)
	return cfg, nil
}
