package store

import (
	"bakery34/model"
	"encoding/json"
	"fmt"
)

type Repository struct{ db *DB }

func NewRepository(db *DB) *Repository                { return &Repository{db: db} }
func (r *Repository) SaveRecord(v model.Record) error { return r.db.Put("records", v.ID, v) }
func (r *Repository) FindRecord(id string) (model.Record, error) {
	var v model.Record
	e := r.db.Get("records", id, &v)
	return v, e
}
func (r *Repository) AllRecords() ([]model.Record, error) {
	raw, e := r.db.List("records")
	if e != nil {
		return nil, e
	}
	out := make([]model.Record, 0, len(raw))
	for _, b := range raw {
		var v model.Record
		if json.Unmarshal(b, &v) == nil {
			out = append(out, v)
		}
	}
	return out, nil
}
func (r *Repository) SaveUser(v model.User) error { return r.db.Put("users", v.ID, v) }
func (r *Repository) FindUser(id string) (model.User, error) {
	var v model.User
	e := r.db.Get("users", id, &v)
	return v, e
}
func (r *Repository) SaveEvent(v model.Event) error {
	if v.ID == "" {
		return fmt.Errorf("event id required")
	}
	return r.db.Put("events", v.ID, v)
}
func (r *Repository) SaveAudit(v model.Audit) error { return r.db.Put("audits", v.ID, v) }
