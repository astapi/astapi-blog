package model

import (
	"github.com/russross/blackfriday"
	"time"
)

type Entry struct {
	Title     string
	Body      string `datastore:",noindex"`
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e Entry) String() string {
	return e.Title
}

func (e *Entry) ToHtml() string {
	return string(blackfriday.MarkdownCommon([]byte(e.Body)))
}

func (e Entry) FormatCreatedAt() string {
	return e.CreatedAt.Format("2006.01.02")
}
