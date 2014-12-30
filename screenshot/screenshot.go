package screenshot
import (
   "os"
   "fmt"
   "github.com/codegangsta/cli"
)

func ScreenshotAction(c *cli.Context) {
   err := TakeScreenshot(c)
   if err != nil {
      fmt.Printf("error %s\n", err)
      os.Exit(1)
   }
}

func TakeScreenshot(c *cli.Context) error {
   shortname := c.String("shortname")
   fmt.Println("TODO", shortname)
   return nil
}
