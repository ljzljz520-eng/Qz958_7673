package store

import (
	"bakery34/model"
	"fmt"
	"time"
)

type Maintenance struct {
	Started time.Time
	Closed  time.Time
	Records int
	Errors  []string
}

func (r *Repository) Health() Maintenance {
	x := Maintenance{Started: time.Now().UTC()}
	x.Records = r.Count()
	return x
}
func (r *Repository) ValidateAll() []string {
	all, e := r.AllRecords()
	if e != nil {
		return []string{e.Error()}
	}
	out := []string{}
	for _, x := range all {
		if e := x.Valid(); e != nil {
			out = append(out, x.ID+": "+e.Error())
		}
		if !model.IsStatus(x.DisplayStatus()) {
			out = append(out, x.ID+": invalid status")
		}
	}
	return out
}
func (r *Repository) RepairStatuses() int {
	all, e := r.AllRecords()
	if e != nil {
		return 0
	}
	n := 0
	for _, x := range all {
		norm := model.NormalizeStatus(x.Status)
		if norm != x.Status {
			x.Status = norm
			_ = r.SaveRecord(x)
			n++
		}
	}
	return n
}
func (r *Repository) PurgeArchived(before time.Time) int {
	all, e := r.AllRecords()
	if e != nil {
		return 0
	}
	n := 0
	for _, x := range all {
		if x.DisplayStatus() == "archived" && x.UpdatedAt.Before(before) {
			if r.Remove(x.ID) == nil {
				n++
			}
		}
	}
	return n
}
func (r *Repository) RequireOpen() error {
	if r.db == nil || r.db.raw == nil {
		return fmt.Errorf("repository closed")
	}
	return nil
}
