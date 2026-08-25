package model

import "testing"

func TestStatusRules(t *testing.T) {
	if !CanTransition("pending", "processing") {
		t.Fatal("transition")
	}
	if CanTransition("archived", "pending") {
		t.Fatal("bad transition")
	}
	if NormalizeStatus("weird") != "pending" {
		t.Fatal("normalize")
	}
}
