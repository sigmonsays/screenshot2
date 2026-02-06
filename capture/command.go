package capture

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Command struct {
}

func (me *Command) Capture(cfg *config.AppConfig, shortname *core.Shortname) error {
	basename := shortname.Value + ".jpg"
	var cmdline []string

	if len(cfg.Capture.Cmdline) > 0 {
		cmdline = cfg.Capture.Cmdline
		cmdline = append(cmdline, basename)
	} else {
		cmdline = []string{cfg.Capture.Command, basename}
	}

	for i, c := range cmdline {
		c = strings.ReplaceAll(c, "%s", basename)
		cmdline[i] = c
	}

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
