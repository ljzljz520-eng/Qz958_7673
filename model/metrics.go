package model

type Summary struct {
	Total      int     `json:"total"`
	Pending    int     `json:"pending"`
	Processing int     `json:"processing"`
	Approved   int     `json:"approved"`
	Archived   int     `json:"archived"`
	Cancelled  int     `json:"cancelled"`
	Quantity   float64 `json:"quantity"`
}

func (s *Summary) Add(r Record) {
	s.Total++
	s.Quantity += r.Quantity
	switch r.DisplayStatus() {
	case "pending":
		s.Pending++
	case "processing":
		s.Processing++
	case "approved":
		s.Approved++
	case "archived":
		s.Archived++
	case "cancelled":
		s.Cancelled++
	}
}
func (s Summary) Complete() bool { return s.Total > 0 && s.Archived == s.Total }
