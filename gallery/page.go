package gallery

type ImageSet []*ImageInfo

type Page struct {
   Title string
   PerPage int
   Images []ImageSet
}


// paginate images into pages
func MakePageImages(page *Page, images []*ImageInfo) ([]ImageSet, error) {
   num_images := len(images)
   num_pages := num_images / page.PerPage
   imageset := make([]ImageSet, 0)

   if num_pages == 0 && num_images > 0 {
      imageset = append(imageset, images)
      return imageset, nil
   }

   for pagenum := 0; pagenum<=num_pages; pagenum++ {
      imageset = append(imageset, make([]*ImageInfo, 0))
      for i := 0; i<page.PerPage; i++ {
         idx := (pagenum*page.PerPage)+i
         if idx + 1 > num_images {
            break
         }
         imageset[pagenum] = append(imageset[pagenum], images[idx])
      }
   }
   return imageset, nil
}

