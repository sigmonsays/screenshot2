package clipboard

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/sigmonsays/screenshot2/config"
	"github.com/sigmonsays/screenshot2/core"
)

type XClip struct {
}

func (me *XClip) CopyToClipboard(cfg *config.AppConfig, shortname *core.Shortname) error {
	objectName := shortname.Value + ".jpg" // name in remote storage

	clipUrl := fmt.Sprintf("http://%s.s3.amazonaws.com/%s", cfg.Upload.Bucket, objectName)
	log.Tracef("clipboard url %s", clipUrl)

	buf := bytes.NewBufferString(clipUrl)

	c := exec.Command("xclip", "-i")
	c.Stdin = buf
	err := c.Run()
	if err != nil {
		return err
	}
	log.Tracef("successfully copied %s to clipboard", clipUrl)
	return nil
}
