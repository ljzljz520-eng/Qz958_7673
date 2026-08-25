package query

import (
	"bakery34/model"
	"sort"
)

type Page struct {
	Items   []model.Record
	Offset  int
	Limit   int
	Total   int
	HasNext bool
}

func Paginate(records []model.Record, offset, limit int) Page {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	total := len(records)
	if offset > total {
		offset = total
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return Page{Items: append([]model.Record(nil), records[offset:end]...), Offset: offset, Limit: limit, Total: total, HasNext: end < total}
}
func SortByCreated(records []model.Record, ascending bool) []model.Record {
	out := append([]model.Record(nil), records...)
	sort.SliceStable(out, func(i, j int) bool {
		if ascending {
			return out[i].CreatedAt.Before(out[j].CreatedAt)
		}
		return out[i].CreatedAt.After(out[j].CreatedAt)
	})
	return out
}
func Quantities(records []model.Record) []float64 {
	out := make([]float64, len(records))
	for i, r := range records {
		out[i] = r.Quantity
	}
	return out
}
func EmptyPage() Page { return Page{Items: []model.Record{}, Limit: 20} }
