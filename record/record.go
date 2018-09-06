package record

import "errors"

type Record struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	TimeStamp   string `json:"timestamp"`
	CallID      string `json:"call_id"`
	Source      int    `json:"source"`
	Sestination int    `json:"destination"`
}

func (r *Record) DataChecks() (err error) {
	if r.ID == "" {
		return errors.New("Record ID cannot be nil")
	}
	if (r.Type == "") || ((r.Type != "start") && (r.Type != "end")) {
		return errors.New("Invalid record type")
	}
	if r.TimeStamp == "" {
		return errors.New("Record time stamp cannot be nil")
	}
	if r.CallID == "" {
		return errors.New("Record call_id cannot be nil")
	}
	return
}
