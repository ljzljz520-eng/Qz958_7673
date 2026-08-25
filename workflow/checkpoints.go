package workflow

import (
	"bakery34/model"
	"bakery34/service"
	"fmt"
	"time"
)

type Checkpoint struct {
	ID       string
	RecordID string
	Stage    string
	Complete bool
	At       time.Time
}
type Tracker struct{ items []Checkpoint }

func NewTracker() *Tracker { return &Tracker{items: []Checkpoint{}} }
func (t *Tracker) Mark(id, record, stage string, done bool) Checkpoint {
	x := Checkpoint{ID: id, RecordID: record, Stage: stage, Complete: done, At: time.Now().UTC()}
	t.items = append(t.items, x)
	return x
}
func (t *Tracker) ForRecord(record string) []Checkpoint {
	out := []Checkpoint{}
	for _, x := range t.items {
		if x.RecordID == record {
			out = append(out, x)
		}
	}
	return out
}
func (t *Tracker) Completed(record string) bool {
	xs := t.ForRecord(record)
	if len(xs) < 4 {
		return false
	}
	for _, x := range xs {
		if !x.Complete {
			return false
		}
	}
	return true
}
func (t *Tracker) Count() int { return len(t.items) }
func ExecuteFourSteps(s *service.Service, r model.Record, actor string) ([]Checkpoint, error) {
	if e := s.RegisterRecord(r); e != nil {
		return nil, e
	}
	t := NewTracker()
	t.Mark("receive", r.ID, "receive", true)
	t.Mark("validate", r.ID, "validate", true)
	if _, e := s.Process(r.ID, actor); e != nil {
		return nil, e
	}
	t.Mark("process", r.ID, "process", true)
	t.Mark("display", r.ID, "display", true)
	return t.ForRecord(r.ID), nil
}
func ValidateCheckpoints(xs []Checkpoint) error {
	if len(xs) < 4 {
		return fmt.Errorf("four checkpoints required")
	}
	for _, x := range xs {
		if x.ID == "" || x.RecordID == "" || x.Stage == "" {
			return fmt.Errorf("invalid checkpoint")
		}
	}
	return nil
}
