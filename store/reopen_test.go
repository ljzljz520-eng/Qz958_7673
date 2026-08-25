package store

import (
	"bakery34/model"
	"os"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "reopen.db"
	defer os.Remove(p)
	d, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := NewRepository(d)
	if e = r.SaveRecord(model.NewRecord("persist", "u", "cocoa", 4, "kg")); e != nil {
		t.Fatal(e)
	}
	if e = d.Close(); e != nil {
		t.Fatal(e)
	}
	d, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	x, e := NewRepository(d).FindRecord("persist")
	if e != nil || x.Ingredient != "cocoa" {
		t.Fatal(x, e)
	}
}
