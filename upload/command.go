package upload

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type Command struct {
}

func (me *Command) Upload(cfg *config.AppConfig, shortname *core.Shortname) error {

	td := &TemplateData{
		Shortname: shortname,
	}

	cmdline, err := templateCommand(cfg.Upload.CommandTemplate, td)
	if err != nil {
		return err
	}
	log.Tracef("evaluated template, cmdline:%v", cmdline)

	return nil
}

type TemplateData struct {
	Shortname *core.Shortname
}

func templateCommand(templ []string, data *TemplateData) ([]string, error) {
	buf, err := json.Marshal(templ)
	if err != nil {
		return nil, err
	}

	t, err := template.New("command").Parse(string(buf))
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer

	err = t.Execute(&out, data)
	if err != nil {
		return nil, err
	}

	var cmd []string
	err = json.Unmarshal(out.Bytes(), &cmd)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
