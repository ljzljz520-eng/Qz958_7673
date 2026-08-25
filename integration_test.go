package bakery34

import (
	"bakery34/api"
	"bakery34/model"
	"bakery34/service"
	"bakery34/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHTTPRecords(t *testing.T) {
	p := "http.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	defer d.Close()
	s := service.New(d)
	h := api.New(s)
	req := httptest.NewRequest(http.MethodPost, "/records", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Fatal(w.Code)
	}
}
func TestRecordFlow34(t *testing.T) {
	p := "bug.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	s := service.New(d)
	s.RegisterRecord(model.NewRecord("bug", "u", "flour", 1, "kg"))
	d.Close()
	d, _ = store.Open(p)
	defer d.Close()
	got, err := workflowRecover(service.New(d), "bug")
	if err != nil {
		t.Fatal(err)
	}
	if got != "normal" {
		t.Fatalf("expected normal status, got %s", got)
	}
}
func workflowRecover(s *service.Service, id string) (string, error) {
	r, e := s.GetRecord(id)
	if e != nil {
		return "", e
	}
	return r.DisplayStatus(), nil
}
