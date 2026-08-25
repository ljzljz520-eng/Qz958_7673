package service

import (
	"bakery34/model"
	"bakery34/store"
	"fmt"
	"strings"
	"time"
)

type Service struct{ repo *store.Repository }

func New(db *store.DB) *Service { return &Service{repo: store.NewRepository(db)} }
func (s *Service) RegisterRecord(r model.Record) error {
	if e := r.Valid(); e != nil {
		return e
	}
	r.Status = model.NormalizeStatus(r.Status)
	return s.repo.SaveRecord(r)
}
func (s *Service) GetRecord(id string) (model.Record, error)      { return s.repo.FindRecord(id) }
func (s *Service) ListRecords() ([]model.Record, error)           { return s.repo.AllRecords() }
func (s *Service) Process(id, actor string) (model.Record, error) { return s.repo.Advance(id, actor) }
func (s *Service) Approve(id, actor string) (model.Record, error) {
	return s.repo.Transition(id, "approved", actor)
}
func (s *Service) Archive(id, actor string) (model.Record, error) { return s.repo.Archive(id, actor) }
func (s *Service) Cancel(id, actor string) (model.Record, error)  { return s.repo.Cancel(id, actor) }
func (s *Service) Search(term string) ([]model.Record, error) {
	all, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(term)
	out := []model.Record{}
	for _, r := range all {
		if strings.Contains(strings.ToLower(r.Ingredient), term) || strings.Contains(strings.ToLower(r.Notes), term) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Service) Summary() (model.Summary, error) {
	all, e := s.ListRecords()
	if e != nil {
		return model.Summary{}, e
	}
	var x model.Summary
	for _, r := range all {
		x.Add(r)
	}
	return x, nil
}
func (s *Service) SeedUser(u model.User) error {
	if u.ID == "" || !u.Eligible() {
		return fmt.Errorf("invalid user")
	}
	return s.repo.SaveUser(u)
}
func (s *Service) Touch(r model.Record) model.Record { r.UpdatedAt = time.Now().UTC(); return r }
