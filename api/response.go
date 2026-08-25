package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	OK    bool   `json:"ok"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, Response{OK: false, Error: err.Error()})
}
func WriteOK(w http.ResponseWriter, v any) { WriteJSON(w, http.StatusOK, Response{OK: true, Data: v}) }
func MethodAllowed(w http.ResponseWriter, methods ...string) {
	w.Header().Set("Allow", join(methods))
	WriteError(w, http.StatusMethodNotAllowed, http.ErrNotSupported)
}
func join(xs []string) string {
	out := ""
	for i, x := range xs {
		if i > 0 {
			out += ", "
		}
		out += x
	}
	return out
}
