package core

import "time"

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
	Value string
}

func (me *Shortname) String() string {
	return me.Value
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
