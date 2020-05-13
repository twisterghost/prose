package prose

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	Title    string            `json:"title"`
	Body     string            `json:"body"`
	Id       string            `json:"id"`
	Author   string            `json:"author"`
	Time     time.Time         `json:"time"`
	Metadata map[string]string `json:"metadata"`
}

func (e *Entry) Format() {
	if e.Time.IsZero() {
		e.Time = time.Now()
	}

	if e.Id == "" {
		e.Id = uuid.New().String()
	}
}

func NewBasicEntry(content string, author string) Entry {
	return Entry{
		Title:    "",
		Body:     content,
		Id:       uuid.New().String(),
		Author:   author,
		Time:     time.Now(),
		Metadata: make(map[string]string),
	}
}
