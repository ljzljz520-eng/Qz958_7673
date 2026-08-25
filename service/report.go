package service

import (
	"bakery34/model"
	"bakery34/query"
	"sort"
	"time"
)

type Report struct {
	Summary     model.Summary
	ByUnit      map[string]float64
	GeneratedAt time.Time
}

func (s *Service) BuildReport() (Report, error) {
	all, e := s.ListRecords()
	if e != nil {
		return Report{}, e
	}
	sum, _ := s.Summary()
	return Report{Summary: sum, ByUnit: query.GroupByUnit(all), GeneratedAt: time.Now().UTC()}, nil
}
func (s *Service) Latest(n int) ([]model.Record, error) {
	all, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	sort.Slice(all, func(i, j int) bool { return all[i].UpdatedAt.After(all[j].UpdatedAt) })
	if n < 0 {
		n = 0
	}
	if n > len(all) {
		n = len(all)
	}
	return all[:n], nil
}
func (s *Service) StatusCounts() map[string]int {
	m := map[string]int{}
	all, e := s.ListRecords()
	if e != nil {
		return m
	}
	for _, r := range all {
		m[r.DisplayStatus()]++
	}
	return m
}
func (s *Service) ReadyForArchive() ([]model.Record, error) {
	all, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	return query.Filter(all, "approved"), nil
}
