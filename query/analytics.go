package query

import (
	"bakery34/model"
	"math"
	"sort"
)

type Analytics struct {
	Mean                float64
	Median              float64
	Min                 float64
	Max                 float64
	DistinctIngredients int
}

func Analyze(records []model.Record) Analytics {
	if len(records) == 0 {
		return Analytics{}
	}
	vals := Quantities(records)
	sort.Float64s(vals)
	sum := 0.0
	seen := map[string]bool{}
	for _, r := range records {
		sum += r.Quantity
		seen[r.Ingredient] = true
	}
	median := vals[len(vals)/2]
	if len(vals)%2 == 0 {
		median = (vals[len(vals)/2-1] + vals[len(vals)/2]) / 2
	}
	return Analytics{Mean: sum / float64(len(vals)), Median: median, Min: vals[0], Max: vals[len(vals)-1], DistinctIngredients: len(seen)}
}
func Round(v float64, places int) float64 { p := math.Pow10(places); return math.Round(v*p) / p }
func StatusRatio(records []model.Record, status string) float64 {
	if len(records) == 0 {
		return 0
	}
	n := len(Filter(records, status))
	return float64(n) / float64(len(records))
}
func TopIngredients(records []model.Record, n int) []string {
	m := map[string]float64{}
	for _, r := range records {
		m[r.Ingredient] += r.Quantity
	}
	type pair struct {
		name  string
		value float64
	}
	ps := []pair{}
	for k, v := range m {
		ps = append(ps, pair{k, v})
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].value > ps[j].value })
	if n > len(ps) {
		n = len(ps)
	}
	out := []string{}
	for _, p := range ps[:n] {
		out = append(out, p.name)
	}
	return out
}
