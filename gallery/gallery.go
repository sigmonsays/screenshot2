package gallery

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sigmonsays/screenshot2/data"
	"github.com/sigmonsays/screenshot2/util"
	"github.com/urfave/cli/v2"
)

type GalleryOptions struct {
	Title                   string
	PerPage                 int
	OutDir                  string
	TemplatePath            string
	ThumbWidth, ThumbHeight int
}

func (o *GalleryOptions) OutPath(path string) string {
	return filepath.Join(o.OutDir, path)
}

func BuildGallery(c *cli.Context) error {

	opts := &GalleryOptions{
		Title:        "Pictures",
		PerPage:      5,
		OutDir:       "out/",
		ThumbWidth:   250,
		ThumbHeight:  250,
		TemplatePath: "data/template/default",
	}

	err := PrepareOutput(opts)
	if err != nil {
		return err
	}

	images, err := FindImages("ex")
	if err != nil {
		return err
	}

	err = MakeThumbnails(opts, images)
	if err != nil {
		return err
	}

	page := &Page{
		Title:   opts.Title,
		PerPage: opts.PerPage,
	}
	page.Images, err = MakePageImages(page, images)
	if err != nil {
		return err
	}

	data := &data.Data{}

	err = RenderTemplate(opts, page, data)
	if err != nil {
		return err
	}

	fmt.Printf("page %+v\n", page)
	return nil
}

func PrepareOutput(opts *GalleryOptions) error {

	os.MkdirAll(opts.OutDir, 0722)

	// copy everything from the template directory in the output directory
	offset := len(opts.TemplatePath)
	err := filepath.Walk(opts.TemplatePath, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".tmpl" {
			return nil
		}
		newpath := filepath.Join(opts.OutDir, path[offset:])
		os.MkdirAll(filepath.Dir(newpath), 0722)
		util.CopyFile(path, newpath)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
