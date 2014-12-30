package main
import (
   "fmt"
   "os"
   "github.com/codegangsta/cli"

   "github.com/sigmonsays/screenshot2/gallery"
   "github.com/sigmonsays/screenshot2/screenshot"
)

func main() {
  app := cli.NewApp()
  app.Name = os.Args[0]
  app.Usage = "screenshot"
  app.Action = func(c *cli.Context) {
      fmt.Println("screenshot")
  }
  app.Commands = []cli.Command{
     {
       Name:      "capture",
       ShortName: "c",
       Usage:     "capture a screenshot",
       Action: func(c *cli.Context) {
         fmt.Println("capture not implemented")
       },
     },
     {
       Name:      "gallery",
       ShortName: "g",
       Usage:     "make a gallery of images",
       Action: gallery.GalleryAction,
     },
     {
       Name:      "screenshot",
       ShortName: "s",
       Usage:     "take a screenshot",
       Action: screenshot.ScreenshotAction,
     },
  }
  app.Run(os.Args)
}

