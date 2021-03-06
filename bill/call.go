package bill

type Call struct {
	BillID        string  `json:"bill_id"`
	Destination   string  `json:"destination"`
	CallStartDate int     `json:"start_date"`
	CallStartTime string  `json:"start_time"`
	CallDuration  string  `json:"duration"` // 0h22m12s
	CallPrice     float64 `json:"price"`
}
