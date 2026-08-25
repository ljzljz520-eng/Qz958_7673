package model

import (
	"testing"
	"time"
)

func TestRecordValidation(t *testing.T) {
	r := NewRecord("r1", "u1", "flour", 2, "kg")
	if e := r.Valid(); e != nil {
		t.Fatal(e)
	}
	if r.DisplayStatus() != "pending" {
		t.Fatal(r.DisplayStatus())
	}
}
func TestFormatHelpers(t *testing.T) {
	r := NewRecord("r", "u", "whole flour", 1, "kg")
	if len(r.Tags()) != 2 {
		t.Fatal(r.Tags())
	}
	if r.Age(time.Now()) < 0 {
		t.Fatal("age")
	}
}
