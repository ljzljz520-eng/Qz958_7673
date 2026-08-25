package audit

import (
	"bakery34/model"
	"sort"
	"strings"
)

func Filter(entries []model.Audit, actor string) []model.Audit {
	out := []model.Audit{}
	for _, e := range entries {
		if actor == "" || e.Actor == actor {
			out = append(out, e)
		}
	}
	return out
}
func Sort(entries []model.Audit) []model.Audit {
	out := append([]model.Audit(nil), entries...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func Search(entries []model.Audit, term string) []model.Audit {
	term = strings.ToLower(term)
	out := []model.Audit{}
	for _, e := range entries {
		if strings.Contains(strings.ToLower(e.Action), term) || strings.Contains(strings.ToLower(e.Target), term) {
			out = append(out, e)
		}
	}
	return out
}
func Actions(entries []model.Audit) map[string]int {
	m := map[string]int{}
	for _, e := range entries {
		m[e.Action]++
	}
	return m
}
