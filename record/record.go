package record

type Record struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	TimeStamp   string `json:"timestamp"`
	CallID      string `json:"call_id"`
	Source      int    `json:"source"`
	Sestination int    `json:"destination"`
}
