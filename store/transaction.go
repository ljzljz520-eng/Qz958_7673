package store

import (
	"bakery34/model"
	"fmt"
	"time"
)

func (r *Repository) Transition(id, to, actor string) (model.Record, error) {
	v, e := r.FindRecord(id)
	if e != nil {
		return v, e
	}
	if !model.CanTransition(v.DisplayStatus(), to) {
		return v, fmt.Errorf("invalid transition")
	}
	v.Status = to
	v.UpdatedAt = time.Now().UTC()
	if e = r.SaveRecord(v); e != nil {
		return v, e
	}
	_ = r.SaveEvent(model.Event{ID: id + "-" + to, RecordID: id, Kind: "status", Message: to, At: v.UpdatedAt})
	_ = r.SaveAudit(model.Audit{ID: id + "-audit-" + to, Actor: actor, Action: "transition", Target: id, At: v.UpdatedAt, Details: to})
	return v, nil
}
func (r *Repository) Archive(id, actor string) (model.Record, error) {
	return r.Transition(id, "archived", actor)
}
func (r *Repository) Cancel(id, actor string) (model.Record, error) {
	return r.Transition(id, "cancelled", actor)
}
func (r *Repository) Advance(id, actor string) (model.Record, error) {
	v, e := r.FindRecord(id)
	if e != nil {
		return v, e
	}
	return r.Transition(id, model.NextStatus(v.DisplayStatus()), actor)
}
