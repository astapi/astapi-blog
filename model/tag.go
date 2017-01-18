package model

import (
	"time"
)

type Tag struct {
	TagName    string
	EntryCount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
