package service

import (
	"bakery34/model"
	"fmt"
	"time"
)

type Notification struct {
	RecordID  string
	Recipient string
	Kind      string
	Message   string
	SentAt    time.Time
}
type Notifier struct{ out []Notification }

func NewNotifier() *Notifier { return &Notifier{out: []Notification{}} }
func (n *Notifier) Send(r model.Record, recipient, kind string) Notification {
	x := Notification{RecordID: r.ID, Recipient: recipient, Kind: kind, Message: n.message(r, kind), SentAt: time.Now().UTC()}
	n.out = append(n.out, x)
	return x
}
func (n *Notifier) message(r model.Record, kind string) string {
	switch kind {
	case "received":
		return fmt.Sprintf("Record %s received", r.ID)
	case "processing":
		return fmt.Sprintf("Record %s is processing", r.ID)
	case "approved":
		return fmt.Sprintf("Record %s approved", r.ID)
	case "archived":
		return fmt.Sprintf("Record %s archived", r.ID)
	case "cancelled":
		return fmt.Sprintf("Record %s cancelled", r.ID)
	}
	return fmt.Sprintf("Record %s updated", r.ID)
}
func (n *Notifier) All() []Notification { return append([]Notification(nil), n.out...) }
func (n *Notifier) Count() int          { return len(n.out) }
func (n *Notifier) Last() Notification {
	if len(n.out) == 0 {
		return Notification{}
	}
	return n.out[len(n.out)-1]
}
func (n *Notifier) Clear() { n.out = nil }
func (s *Service) NotificationFor(id, kind, recipient string) (Notification, error) {
	r, e := s.GetRecord(id)
	if e != nil {
		return Notification{}, e
	}
	return NewNotifier().Send(r, recipient, kind), nil
}
