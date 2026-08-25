package model

import (
	"fmt"
	"strings"
	"time"
)

func (r Record) Label() string {
	return fmt.Sprintf("%s: %.2f %s %s", r.ID, r.Quantity, r.Unit, r.Ingredient)
}
func (r Record) Age(now time.Time) time.Duration {
	if now.Before(r.CreatedAt) {
		return 0
	}
	return now.Sub(r.CreatedAt)
}
func (r Record) Tags() []string {
	parts := strings.Fields(strings.ToLower(r.Ingredient))
	out := []string{}
	for _, p := range parts {
		if len(p) > 2 {
			out = append(out, p)
		}
	}
	return out
}
func ParseQuantity(s string) (float64, error) {
	var q float64
	_, e := fmt.Sscanf(strings.TrimSpace(s), "%f", &q)
	if e != nil || q <= 0 {
		return 0, ErrInvalid("quantity")
	}
	return q, nil
}
func NormalizeUnit(s string) string {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "g", "gram", "grams":
		return "g"
	case "kg", "kilogram", "kilograms":
		return "kg"
	case "ml", "milliliter", "milliliters":
		return "ml"
	}
	return strings.ToLower(strings.TrimSpace(s))
}
func (u User) Label() string    { return fmt.Sprintf("%s (%s)", u.Name, u.Tier) }
func (e Event) Summary() string { return e.Kind + ": " + e.Message }
func (a Audit) Summary() string { return a.Actor + " " + a.Action + " " + a.Target }
