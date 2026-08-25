package query

import (
	"bakery34/model"
	"sort"
	"strings"
)

func Filter(records []model.Record, status string) []model.Record {
	out := []model.Record{}
	for _, r := range records {
		if status == "" || r.DisplayStatus() == status {
			out = append(out, r)
		}
	}
	return out
}
func SortByQuantity(records []model.Record) []model.Record {
	out := append([]model.Record(nil), records...)
	sort.Slice(out, func(i, j int) bool { return out[i].Quantity > out[j].Quantity })
	return out
}
func SearchIngredient(records []model.Record, term string) []model.Record {
	term = strings.ToLower(term)
	out := []model.Record{}
	for _, r := range records {
		if strings.Contains(strings.ToLower(r.Ingredient), term) {
			out = append(out, r)
		}
	}
	return out
}
func GroupByUnit(records []model.Record) map[string]float64 {
	m := map[string]float64{}
	for _, r := range records {
		m[r.Unit] += r.Quantity
	}
	return m
}
