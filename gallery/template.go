package gallery
import (
   "os"
   "path/filepath"
   "text/template"
)

// renders the page template
func RenderTemplate(opts *GalleryOptions, page *Page) error {
   pageFile := opts.OutPath("index.html")
   pageTmplFile := filepath.Join(opts.TemplatePath, "page.tmpl")

   pageTmpl := template.Must(template.New("page.tmpl").ParseFiles(pageTmplFile))

   f, err := os.Create(pageFile)
   if err != nil {
      return err
   }
   defer f.Close()

   err = pageTmpl.Execute(f, page)
   if err != nil {
      return err
   }

   return nil
}
