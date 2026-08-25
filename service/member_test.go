package service

import (
	"bakery34/model"
	"bakery34/store"
	"os"
	"testing"
)

func TestMemberRewards(t *testing.T) {
	p := "mem.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	defer d.Close()
	s := New(d)
	if e := s.SeedUser(model.User{ID: "u", Name: "Baker", Email: "b@x", Active: true}); e != nil {
		t.Fatal(e)
	}
	u, e := s.Reward("u", 600)
	if e != nil || u.Points != 600 {
		t.Fatal(u, e)
	}
	u, e = s.RecalculateTier("u")
	if e != nil || u.Tier != "silver" {
		t.Fatal(u, e)
	}
}
