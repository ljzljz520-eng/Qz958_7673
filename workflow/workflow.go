package workflow

import (
	"bakery34/model"
	"bakery34/service"
	"fmt"
)

type Engine struct{ s *service.Service }

func New(s *service.Service) *Engine { return &Engine{s: s} }
func (e *Engine) Intake(r model.Record) error {
	if err := r.Valid(); err != nil {
		return err
	}
	return e.s.RegisterRecord(r)
}
func (e *Engine) Review(id, actor string) (model.Record, error) { return e.s.Process(id, actor) }
func (e *Engine) Finalize(id, actor string) (model.Record, error) {
	r, err := e.s.Approve(id, actor)
	if err != nil {
		return r, err
	}
	return e.s.Archive(id, actor)
}
func (e *Engine) Recover(id string) (string, error) {
	r, err := e.s.GetRecord(id)
	if err != nil {
		return "", err
	}
	if r.DisplayStatus() == "pending" {
		return "normal", nil
	}
	if r.DisplayStatus() == "archived" {
		return "normal", nil
	}
	return fmt.Sprintf("%s", r.DisplayStatus()), nil
}
