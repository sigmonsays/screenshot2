package gallery
import (
   "os"
   "strings"
   "path/filepath"
)

var ImageExtensions = map[string]bool {
   ".jpg": true,
   ".jpeg": true,
   ".gif": true,
   ".png": true,
   ".bmp": true,
   ".xcf": false,
}

func IsValidImage(path string) bool {
   ext := strings.ToLower(filepath.Ext(path))
   enabled, ok := ImageExtensions[ext]
   valid := (enabled == true && ok == true)
   return valid
}

func FindImages(imgpath, outdir string) ([]*ImageInfo, error) {
   images := make([]*ImageInfo, 0)
   err := filepath.Walk(imgpath, func(path string, info os.FileInfo, err error) error {
      if strings.HasPrefix(path, outdir) {
         return nil
      }
      if IsValidImage(path) {
         images = append(images, &ImageInfo{SrcPath:path})
      }
      return nil
   })
   if err != nil {
      return images, err
   }
   return images, nil
}


