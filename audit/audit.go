package audit

import (
	"bakery34/model"
	"fmt"
	"time"
)

type Logger struct{ entries []model.Audit }

func New() *Logger { return &Logger{entries: []model.Audit{}} }
func (l *Logger) Add(actor, action, target string) model.Audit {
	a := model.Audit{ID: fmt.Sprintf("%d", len(l.entries)+1), Actor: actor, Action: action, Target: target, At: time.Now().UTC()}
	l.entries = append(l.entries, a)
	return a
}
func (l *Logger) All() []model.Audit { return append([]model.Audit(nil), l.entries...) }
func (l *Logger) Count() int         { return len(l.entries) }
