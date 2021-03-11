package capture

import (
	"os/exec"

	"github.com/sigmonsays/screenshot2/config"
)

type Command struct {
}

func (me *Command) Capture(cfg *config.AppConfig, shortname *Shortname) error {
	cmdline := []string{"import", shortname.String() + ".jpg"}
	c := exec.Command(cmdline[0], cmdline[1:]...)
	c.Dir = cfg.DataDir
	log.Tracef("cmdline %v (%s)", cmdline, c.Dir)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil

}
