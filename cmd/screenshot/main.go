package main

import (
	"fmt"
	"os"

	"github.com/sigmonsays/screenshot2/gallery"
	"github.com/sigmonsays/screenshot2/screenshot"
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
			Name: "capture",
			// ShortName: "c",
			Usage: "capture a screenshot",
			Action: func(c *cli.Context) error {
				fmt.Println("capture not implemented")
				return nil
			},
		},
		{
			Name: "gallery",
			// ShortName: "g",
			Usage:  "make a gallery of images",
			Action: gallery.GalleryAction,
		},
		{
			Name: "screenshot",
			// ShortName: "s",
			Usage:  "take a screenshot",
			Action: screenshot.ScreenshotAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "shortname",
					Value: "shortname",
					Usage: "shortname for screenshot",
				},
			},
		},
	}
	app.Run(os.Args)
}
