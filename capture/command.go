package capture

import "os/exec"

type Command struct {
}

func (me *Command) Capture(shortname *Shortname) error {
	cmdline := []string{"import", shortname.String()}
	log.Tracef("cmdline %v", cmdline)
	c := exec.Command(cmdline[0], cmdline[1:]...)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil

}
