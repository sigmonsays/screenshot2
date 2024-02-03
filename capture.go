package screenshot2

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sigmonsays/screenshot2/capture"
	"github.com/sigmonsays/screenshot2/clipboard"
	"github.com/sigmonsays/screenshot2/core"
	"github.com/sigmonsays/screenshot2/upload"
	"github.com/urfave/cli/v2"
)

func Capture(c *cli.Context) error {
	cfg, err := GetConfig(c.String("config"))
	if err != nil {
		return err
	}

	shortname := core.NewShortname(c.String("shortname"))
	shortname.IncludeDate = true
	delay := c.Int("delay")
	log.Tracef("GalleryCapture shortname:%s delay:%d",
		shortname.GetShortname(), delay)

	if delay > 0 {
		log.Tracef("delaying %d seconds before capture", delay)
		time.Sleep(time.Duration(delay) * time.Second)
	}

	cap := &capture.Capture{}
	err = cap.Capture(cfg, shortname)
	if err != nil {
		return err
	}

	up := &upload.Upload{}
	err = up.Upload(cfg, shortname)
	if err != nil {
		return err
	}

	clip := clipboard.Clipboard{}
	err = clip.CopyToClipboard(cfg, shortname)
	if err != nil {
		return err
	}

	if shortname.Url != "" {
		fmt.Printf("%s\n", shortname.Url)
		logfile := filepath.Join(cfg.DataDir, "screenshot-url.txt")
		err := WriteLogFile(logfile, shortname.Url)
		if err != nil {
			log.Warnf("WriteLogFile %s; %s: %s", logfile, shortname.Url, err)
		}
	}

	if shortname.LocalFile != "" && cfg.Capture.KeepLocal == false {
		log.Debugf("Delete local file %s", shortname.LocalFile)
		err = os.Remove(shortname.LocalFile)
		if err != nil {
			log.Warnf("Remove %s: %s", shortname.LocalFile, err)
		}
	}

	return nil
}

func WriteLogFile(logfile string, url string) error {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", url)
	return nil
}
