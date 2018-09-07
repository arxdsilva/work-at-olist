package record

import (
	"errors"
	"strconv"
	"time"
)

type Record struct {
	ID          string
	Type        string `json:"type"`
	TimeStamp   string `json:"timestamp"`
	CallID      string `json:"call_id"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Month       int
}

func (r *Record) DataChecks() (err error) {
	if (r.Type == "") || ((r.Type != "start") && (r.Type != "end")) {
		return errors.New("Invalid record type")
	}
	if r.TimeStamp == "" {
		return errors.New("Record time stamp cannot be nil")
	}
	if r.CallID == "" {
		return errors.New("Record call_id cannot be nil")
	}
	if (r.Type == "start") && ((r.Source == "") || (r.Destination == "")) {
		return errors.New("Record start cannot have source or destination nil")
	}
	if (r.Type == "start") && (invalidPhone(r.Source) || invalidPhone(r.Destination)) {
		return errors.New("Invalid record start source or destination numbers")
	}
	t, err := time.Parse(time.RFC3339, r.TimeStamp)
	if err != nil {
		return
	}
	r.Month = int(t.Month())
	return
}

func invalidPhone(p string) bool {
	if _, err := strconv.Atoi(p); err != nil {
		return true
	}
	if (len(p) > 11) || (len(p) < 10) {
		return true
	}
	return false
}
