package workflow

import (
	"bakery34/model"
	"bakery34/service"
	"bakery34/store"
	"os"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	p := "wf1.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	defer d.Close()
	e := New(service.New(d))
	if err := e.Intake(model.NewRecord("one", "u", "flour", 1, "kg")); err != nil {
		t.Fatal(err)
	}
}
func TestWorkflowTwo(t *testing.T) {
	p := "wf2.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	defer d.Close()
	e := New(service.New(d))
	e.Intake(model.NewRecord("two", "u", "salt", 1, "kg"))
	if _, err := e.Review("two", "staff"); err != nil {
		t.Fatal(err)
	}
	if _, err := e.Finalize("two", "staff"); err != nil {
		t.Fatal(err)
	}
}
func TestWorkflowThree(t *testing.T) {
	p := "wf3.db"
	defer os.Remove(p)
	d, _ := store.Open(p)
	e := New(service.New(d))
	e.Intake(model.NewRecord("three", "u", "butter", 1, "kg"))
	d.Close()
	d, _ = store.Open(p)
	defer d.Close()
	if got, err := New(service.New(d)).Recover("three"); err != nil || got != "normal" {
		t.Fatal(got, err)
	}
}
