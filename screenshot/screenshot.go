package screenshot

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func ScreenshotAction(c *cli.Context) error {
	err := TakeScreenshot(c)
	if err != nil {
		fmt.Printf("error %s\n", err)
		os.Exit(1)
	}
	return nil
}

func TakeScreenshot(c *cli.Context) error {
	shortname := c.String("shortname")
	log.Tracef("TakeScreenshot shortname:%s", shortname)
	return nil
}
