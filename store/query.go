package store

import (
	"bakery34/model"
	"encoding/json"
	"sort"
)

func (r *Repository) RecordsByUser(uid string) ([]model.Record, error) {
	all, e := r.AllRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, x := range all {
		if x.UserID == uid {
			out = append(out, x)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}
func (r *Repository) SaveMany(v []model.Record) error {
	for _, x := range v {
		if e := r.SaveRecord(x); e != nil {
			return e
		}
	}
	return nil
}
func (r *Repository) Export() ([]byte, error) {
	all, e := r.AllRecords()
	if e != nil {
		return nil, e
	}
	return json.Marshal(all)
}
func (r *Repository) Count() int {
	all, e := r.AllRecords()
	if e != nil {
		return 0
	}
	return len(all)
}
func (r *Repository) Has(id string) bool     { _, e := r.FindRecord(id); return e == nil }
func (r *Repository) Remove(id string) error { return r.db.Delete("records", id) }
