package store

import (
	"bakery34/model"
	"os"
	"testing"
)

func TestRepositoryRoundTrip(t *testing.T) {
	p := "store-test.db"
	defer os.Remove(p)
	d, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := NewRepository(d)
	x := model.NewRecord("r", "u", "sugar", 3, "kg")
	if e = r.SaveRecord(x); e != nil {
		t.Fatal(e)
	}
	y, e := r.FindRecord("r")
	if e != nil || y.Ingredient != "sugar" {
		t.Fatal(y, e)
	}
	d.Close()
}
