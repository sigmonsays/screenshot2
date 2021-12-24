package gallery

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/sigmonsays/screenshot2/data"
)

// renders the page template
func RenderTemplate(opts *GalleryOptions, page *Page, data *data.Data) error {

	templateFile := "template/default/page.tmpl"

	t, err := data.Open(templateFile)
	if err != nil {
		return err
	}
	defer t.Close()

	tmpl, err := ioutil.ReadAll(t)
	if err != nil {
		return err
	}

	pageTmpl := template.Must(template.New("page.tmpl").Parse(string(tmpl)))

	pageFile := opts.OutPath("index.html")
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
