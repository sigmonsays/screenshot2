package postprocess

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type PostProcess struct {
}

func (me *PostProcess) Run(cfg *config.AppConfig, shortname *core.Shortname) error {
	log.Tracef("starting postprocess shortname:%s", shortname)

	td := map[string]any{
		"shortname": shortname.GetShortname(),
		"filename":  shortname.LocalFile,
		"url":       shortname.Url,
	}

	for _, pp := range cfg.PostProcess {

		// parse command as template
		t, err := template.New("").Parse(pp.Command)
		if err != nil {
			return err
		}

		var out bytes.Buffer

		err = t.Execute(&out, td)
		if err != nil {
			return err
		}
		command := out.String()

		log.Tracef("run post process %s", command)
		cmdline := []string{
			"bash",
			"-c",
			command,
		}
		c := exec.Command("/usr/bin/env", cmdline...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err = c.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
