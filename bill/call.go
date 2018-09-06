package bill

type Call struct {
	Destination   int     `json:"destination"`
	CallStartDate int     `json:"start_date"`
	CallStartTime string  `json:"start_time"`
	CallDuration  string  `json:"duration"` // 0h22m12s
	CallPrice     float64 `json:"price"`
}
