package gallery
import (
   "fmt"
   "os"
   "path/filepath"
   "github.com/disintegration/imaging"
   "github.com/sigmonsays/screenshot2/util"
)


func MakeThumbnails(opts *GalleryOptions, images []*ImageInfo) error {

   for _, image := range images {

      img, err := imaging.Open(image.SrcPath)
      if err != nil {
         fmt.Printf("error opening %s: %s\n", image.SrcPath, err)
         continue
      }

      image.ThumbnailPath = "thumbs/" + image.SrcPath
      image.OrigPath = "orig/" + image.SrcPath

      thumb_path := opts.OutPath("thumbs/" + image.SrcPath)
      os.MkdirAll(filepath.Dir(thumb_path), 0722)

      orig_path := opts.OutPath("orig/" + image.SrcPath)
      os.MkdirAll(filepath.Dir(orig_path), 0722)


      img_thumb := imaging.Resize(img, opts.ThumbWidth, opts.ThumbHeight, imaging.Gaussian)

      err = imaging.Save(img_thumb, thumb_path)
      if err != nil {
         fmt.Printf("error saving thumb %s: %s\n", thumb_path, err)
         continue
      }


      err = util.CopyFile(image.SrcPath, orig_path)
      if err != nil {
         fmt.Printf("error copyingc original %s: %s\n", orig_path, err)
         continue
      }

   }

   return nil
}
