package mackerel

type Host struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Tyep   string `json:"type"`
	Status string `json:"status"`
}
