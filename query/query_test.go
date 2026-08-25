package query

import (
	"bakery34/model"
	"testing"
)

func TestQueryFiltering(t *testing.T) {
	a := []model.Record{{ID: "1", Ingredient: "flour", Quantity: 2, Unit: "kg", Status: "pending"}, {ID: "2", Ingredient: "sugar", Quantity: 1, Unit: "kg", Status: "approved"}}
	if len(Filter(a, "pending")) != 1 {
		t.Fatal("filter")
	}
	if len(SearchIngredient(a, "SUG")) != 1 {
		t.Fatal("search")
	}
	if GroupByUnit(a)["kg"] != 3 {
		t.Fatal("group")
	}
}
