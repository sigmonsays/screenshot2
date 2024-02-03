package main

import (
	"os"
	"path/filepath"

	"github.com/sigmonsays/screenshot2"
	"github.com/urfave/cli/v2"

	gologging "github.com/sigmonsays/go-logging"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "screenshot"
	app.Before = func(c *cli.Context) error {
		loglevel := c.String("loglevel")
		if loglevel != "" {
			gologging.SetLogLevel(loglevel)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "loglevel",
			Aliases: []string{"l"},
			Usage:   "log level",
			Value:   "INFO",
		},
	}

	home := os.Getenv("HOME")
	cfgfile := filepath.Join(home, ".screenshot.yaml")

	app.Commands = []*cli.Command{
		{
			Name:   "config",
			Usage:  "show screenshot config",
			Action: screenshot2.Config,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "config file",
					Value:   cfgfile,
				},
				&cli.StringFlag{
					Name:    "shortname",
					Aliases: []string{"s"},
					Usage:   "shortname of screenshot",
				},
			},
		},
		{
			Name:   "capture",
			Usage:  "capture a screenshot",
			Action: screenshot2.Capture,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "config file",
					Value:   cfgfile,
				},
				&cli.StringFlag{
					Name:    "shortname",
					Aliases: []string{"s"},
					Usage:   "shortname of screenshot",
				},
				&cli.IntFlag{
					Name:    "delay",
					Aliases: []string{"d"},
					Usage:   "delay before taking screenshot",
				},
			},
		},
		{
			Name:   "gallery",
			Usage:  "make a gallery of images",
			Action: screenshot2.GalleryAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "in",
					Aliases: []string{"i"},
					Usage:   "input directory",
					Value:   ".",
				},
				&cli.StringFlag{
					Name:    "out",
					Aliases: []string{"o"},
					Usage:   "output directory",
					Value:   ".",
				},
				&cli.StringFlag{
					Name:    "title",
					Aliases: []string{"t"},
					Usage:   "title",
					Value:   "Pictures",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Tracef("app exited with error: %s", err)
		os.Exit(1)
	}
}
