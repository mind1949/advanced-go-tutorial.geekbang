package biz

import (
	"time"
)

type Article struct {
	ID          int64
	Title       string
	Content     string
	CreatedTime time.Time
	UpdatedTime time.Time
}
