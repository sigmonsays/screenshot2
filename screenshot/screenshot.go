package screenshot

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
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
	fmt.Println("TODO", shortname)
	return nil
}
