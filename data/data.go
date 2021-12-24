package data

import (
	"io/fs"
	"os"
)

type Data struct {
}

func (me *Data) Open(name string) (fs.File, error) {
	if FileExists(name) {
		log.Tracef("using local file for %s", name)
		return os.Open(name)
	}
	log.Tracef("using embeded file for %s", name)
	return FS.Open(name)
}

func FileExists(path string) bool {
	ret := false
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		ret = false
		return ret
	}

	ret = true
	return ret
}
