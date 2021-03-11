package capture

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Command struct {
}

func (me *Command) Capture(cfg *config.AppConfig, shortname *core.Shortname) error {
	basename := shortname.Value + ".jpg"
	cmdline := []string{cfg.Capture.Command, basename}
	c := exec.Command(cmdline[0], cmdline[1:]...)
	c.Dir = cfg.DataDir
	debug := true
	if debug {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}
	log.Tracef("cmdline %v (%s)", cmdline, c.Dir)
	err := c.Run()
	if err != nil {
		return err
	}
	shortname.LocalFile = filepath.Join(cfg.DataDir, basename)
	return nil
}
