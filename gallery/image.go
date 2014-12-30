package gallery

type ImageInfo struct {
   // original path from listing
   SrcPath string

   // paths in output directory
   OrigPath string
   ThumbnailPath string

   link string
   thumbSrc string
}

func (i *ImageInfo) Link() string {
   return i.OrigPath
}

func (i *ImageInfo) ThumbSrc() string {
   return i.ThumbnailPath
}
