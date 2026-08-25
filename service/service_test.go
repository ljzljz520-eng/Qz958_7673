package service

import (
	"bakery34/model"
	"bakery34/store"
	"os"
	"testing"
)

func TestServiceLifecycle(t *testing.T) {
	p := "svc.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	defer d.Close()
	s := New(d)
	if e := s.RegisterRecord(model.NewRecord("r", "u", "yeast", 2, "g")); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Process("r", "staff"); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Approve("r", "staff"); e != nil {
		t.Fatal(e)
	}
}
