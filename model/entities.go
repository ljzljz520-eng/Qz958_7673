package model

import "time"

type Record struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Ingredient string    `json:"ingredient"`
	Quantity   float64   `json:"quantity"`
	Unit       string    `json:"unit"`
	Status     string    `json:"status"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Tier   string `json:"tier"`
	Points int    `json:"points"`
	Active bool   `json:"active"`
}
type Event struct {
	ID       string    `json:"id"`
	RecordID string    `json:"record_id"`
	Kind     string    `json:"kind"`
	Message  string    `json:"message"`
	At       time.Time `json:"at"`
}
type Audit struct {
	ID      string    `json:"id"`
	Actor   string    `json:"actor"`
	Action  string    `json:"action"`
	Target  string    `json:"target"`
	At      time.Time `json:"at"`
	Details string    `json:"details"`
}

func (r Record) Valid() error {
	if r.ID == "" {
		return ErrInvalid("id required")
	}
	if r.Ingredient == "" {
		return ErrInvalid("ingredient required")
	}
	if r.Quantity <= 0 {
		return ErrInvalid("quantity must be positive")
	}
	if r.Unit == "" {
		return ErrInvalid("unit required")
	}
	return nil
}
func (r Record) IsOpen() bool { return r.Status == "pending" || r.Status == "processing" }
func (r Record) DisplayStatus() string {
	if r.Status == "" {
		return "pending"
	}
	return r.Status
}
func (u User) Eligible() bool { return u.Active && u.Email != "" }
func (u *User) AddPoints(n int) {
	if n > 0 {
		u.Points += n
	}
}
func NewRecord(id, user, ingredient string, q float64, unit string) Record {
	now := time.Now().UTC()
	return Record{ID: id, UserID: user, Ingredient: ingredient, Quantity: q, Unit: unit, Status: "pending", CreatedAt: now, UpdatedAt: now}
}

type InvalidError string

func (e InvalidError) Error() string { return string(e) }
func ErrInvalid(s string) error      { return InvalidError(s) }
