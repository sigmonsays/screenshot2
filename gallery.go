package screenshot2

import (
	"fmt"
	"os"

	"github.com/sigmonsays/screenshot2/gallery"
	"github.com/urfave/cli/v2"
)

func GalleryAction(c *cli.Context) error {
	err := gallery.BuildGallery(c)
	if err != nil {
		fmt.Printf("error %s\n", err)
		os.Exit(1)
	}
	return nil
}
