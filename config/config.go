package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func (c *AppConfig) LoadDefault() {
	*c = *GetDefaultConfig()
}

func (c *AppConfig) LoadYaml(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(nil)
	_, err = b.ReadFrom(f)
	if err != nil {
		return err
	}

	if err := c.LoadYamlBuffer(b.Bytes()); err != nil {
		return err
	}

	if err := c.FixupConfig(); err != nil {
		return err
	}

	return nil
}

func (c *AppConfig) LoadYamlBuffer(buf []byte) error {
	err := yaml.Unmarshal(buf, c)
	if err != nil {
		return err
	}
	return nil
}

func GetDefaultConfig() *AppConfig {
	cfg := &AppConfig{}
	err := cfg.LoadYamlBuffer([]byte(defaultConf))
	if err != nil {
		log.Warnf("LoadYamlBuffer: %s", err)
	}
	return cfg
}

// after loading configuration this gives us a spot to "fix up" any configuration
// or abort the loading process
func (c *AppConfig) FixupConfig() error {
	// var emptyConfig AppConfig
	home := os.Getenv("HOME")

	if home != "" {

		if strings.HasPrefix(c.DataDir, "/") == false {
			c.DataDir = filepath.Join(home, c.DataDir)
			log.Debugf("Setting datadir relative to %s: datadir=%s",
				home, c.DataDir)
		}
	}

	return nil
}

func (me *AppConfig) PrintConfig() {
	d, err := yaml.Marshal(me)
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}
	fmt.Println("-- Configuration --")
	fmt.Println(string(d))
}
