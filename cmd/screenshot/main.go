package main

import (
	"fmt"
	"os"

	"github.com/sigmonsays/screenshot2/capture"
	"github.com/sigmonsays/screenshot2/gallery"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "screenshot"
	app.Action = func(c *cli.Context) error {
		fmt.Println("screenshot")
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:   "capture",
			Usage:  "capture a screenshot",
			Action: capture.GalleryCapture,
			Flags: []cli.Flag{
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
					Name:    "shortname",
					Aliases: []string{"s"},
					Usage:   "shortname of screenshot",
				},
			},
		},
	}
	app.Run(os.Args)
}
