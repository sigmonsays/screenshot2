package main

import (
	"os"
	"path/filepath"

	"github.com/sigmonsays/screenshot2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "screenshot"

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
