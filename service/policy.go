package service

import (
	"bakery34/model"
	"fmt"
	"strings"
)

type Policy struct {
	AllowedUnits  map[string]bool
	MaxQuantity   float64
	RequireMember bool
}

func DefaultPolicy() Policy {
	return Policy{AllowedUnits: map[string]bool{"g": true, "kg": true, "ml": true, "l": true}, MaxQuantity: 10000, RequireMember: true}
}
func (p Policy) Validate(r model.Record) error {
	if p.RequireMember && strings.TrimSpace(r.UserID) == "" {
		return fmt.Errorf("member required")
	}
	if !p.AllowedUnits[r.Unit] {
		return fmt.Errorf("unsupported unit")
	}
	if r.Quantity > p.MaxQuantity {
		return fmt.Errorf("quantity exceeds policy")
	}
	return nil
}
func (s *Service) RegisterWithPolicy(r model.Record, p Policy) error {
	if e := p.Validate(r); e != nil {
		return e
	}
	if e := model.ValidateAgainstCatalog(r); e != nil {
		return e
	}
	return s.RegisterRecord(r)
}
func (s *Service) CanEdit(id string) bool {
	r, e := s.GetRecord(id)
	if e != nil {
		return false
	}
	return r.DisplayStatus() == "pending" || r.DisplayStatus() == "processing"
}
func (s *Service) EditNotes(id, notes string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if !s.CanEdit(id) {
		return fmt.Errorf("record locked")
	}
	r.Notes = strings.TrimSpace(notes)
	return s.repo.SaveRecord(r)
}
