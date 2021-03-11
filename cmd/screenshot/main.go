package main

import (
	"os"
	"path/filepath"

	"github.com/sigmonsays/screenshot2/capture"
	"github.com/sigmonsays/screenshot2/gallery"
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
			Name:   "capture",
			Usage:  "capture a screenshot",
			Action: capture.GalleryCapture,
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
			Action: gallery.GalleryAction,
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
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Tracef("app exited with error: %s", err)
		os.Exit(1)
	}
}
