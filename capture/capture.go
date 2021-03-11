package capture

import "github.com/urfave/cli/v2"

func GalleryCapture(c *cli.Context) error {
	shortname := NewShortname(c.String("shortname"))
	log.Tracef("GalleryCapture shortname:%s", shortname.Value)

	cap := &Command{}
	err := cap.Capture(shortname)
	if err != nil {
		return err
	}
	return nil
}
