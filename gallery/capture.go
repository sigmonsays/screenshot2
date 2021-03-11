package gallery

import "github.com/urfave/cli/v2"

func GalleryCapture(c *cli.Context) error {
	shortname := c.String("shortname")

	log.Tracef("GalleryCapture shortname:%s", shortname)
	return nil
}
