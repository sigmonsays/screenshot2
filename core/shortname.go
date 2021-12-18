package core

import (
	"bytes"
	"fmt"
	"time"
)

func NewShortname(v string) *Shortname {
	s := &Shortname{
		Value: v,
	}
	if s.IsEmpty() {
		s.FillWithDate()
	}
	return s
}

type Shortname struct {
	Value       string
	LocalFile   string
	Url         string
	IncludeDate bool
}

// returns shortname formatted
func (me *Shortname) GetShortname() string {
	var out bytes.Buffer
	now := time.Now()
	if me.IncludeDate {
		x := now.Format("2006-01-02/15:04:05.999999999")
		fmt.Fprintf(&out, "%s/", x)
	}
	fmt.Fprintf(&out, "%s", me.Value)
	return out.String()
}

func (me *Shortname) String() string {
	return me.GetShortname()
}

func (me *Shortname) IsEmpty() bool {
	return me.Value == ""
}

func (me *Shortname) FillWithDate() {
	n := time.Now()
	// Mon Jan 2 15:04:05 MST 2006
	me.Value = n.Format("20060102150405")
	log.Tracef("shortname set to date %s", me.Value)
}
