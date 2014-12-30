package gallery
import (
   "os"
   "fmt"
   "path/filepath"
   "github.com/codegangsta/cli"
)

func GalleryAction(c *cli.Context) {
   err := BuildGallery(c)
   if err != nil {
      fmt.Printf("error %s\n", err)
      os.Exit(1)
   }
}

type GalleryOptions struct {
   OutDir string
   ThumbWidth, ThumbHeight int
}
func (o *GalleryOptions) OutPath(path string) string {
   return filepath.Join(o.OutDir, path)
}

func BuildGallery(c *cli.Context) error {

   opts := &GalleryOptions{
      OutDir: "out/",
      ThumbWidth: 250,
      ThumbHeight: 250,
   }

   images, err := FindImages("ex")
   if err != nil {
      return err
   }

   err = MakeThumbnails(opts, images)
   if err != nil {
      return err
   }

   for i, image := range images {
      fmt.Printf("%d. %+v\n", i, image)
   }

   page := &Page{
      Title: "Pictures",
      PerPage: 5,
   }
   page.Images, err = MakePageImages(page, images)
   if err != nil {
      return err
   }

   err = RenderTemplate(opts, page)
   if err != nil {
      return err
   }

   fmt.Printf("page %+v\n", page)
   return nil
}
