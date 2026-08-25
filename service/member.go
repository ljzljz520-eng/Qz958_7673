package service

import (
	"bakery34/model"
	"fmt"
)

func (s *Service) Member(id string) (model.User, error) { return s.repo.FindUser(id) }
func (s *Service) Reward(id string, points int) (model.User, error) {
	u, e := s.Member(id)
	if e != nil {
		return u, e
	}
	if points < 1 {
		return u, fmt.Errorf("points must be positive")
	}
	u.AddPoints(points)
	e = s.repo.SaveUser(u)
	return u, e
}
func (s *Service) MemberTier(points int) string {
	switch {
	case points >= 1000:
		return "gold"
	case points >= 500:
		return "silver"
	default:
		return "bronze"
	}
}
func (s *Service) RecalculateTier(id string) (model.User, error) {
	u, e := s.Member(id)
	if e != nil {
		return u, e
	}
	u.Tier = s.MemberTier(u.Points)
	e = s.repo.SaveUser(u)
	return u, e
}
