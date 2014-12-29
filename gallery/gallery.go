package gallery
import (
   "fmt"
   "github.com/codegangsta/cli"
)

func GalleryAction(c *cli.Context) {
   fmt.Println("gallery", c.Args())
}
