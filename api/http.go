package api

import (
	"bakery34/model"
	"bakery34/service"
	"encoding/json"
	"net/http"
	"strings"
)

func New(s *service.Service) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/records", records(s))
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); w.Write([]byte("ok")) })
	return m
}
func records(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			var x model.Record
			if json.NewDecoder(r.Body).Decode(&x) != nil || s.RegisterRecord(x) != nil {
				http.Error(w, "invalid", 400)
				return
			}
			json.NewEncoder(w).Encode(x)
			return
		}
		all, e := s.ListRecords()
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(all)
	}
}
func ParsePath(path string) (string, string) {
	p := strings.Split(strings.Trim(path, "/"), "/")
	if len(p) >= 2 {
		return p[0], p[1]
	}
	return "", ""
}
