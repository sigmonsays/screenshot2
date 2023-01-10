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
	objectName := shortname.GetShortname() + ".jpg" // name in remote storage
	clipUrl := shortname.Url

	log.Tracef("objectName:%s clipboard url %s", objectName, clipUrl)

	buf := bytes.NewBufferString(clipUrl)

	c := exec.Command("xclip", "-i")
	c.Stdin = buf
	err := c.Run()
	if err != nil {
		return err
	}
	log.Tracef("successfully copied %s to clipboard", clipUrl)

	fmt.Printf("%s\n", clipUrl)
	return nil
}
