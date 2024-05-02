package entity

import (
	"time"
)

type ListQuery struct {
	Offset *int
	Limit  *int
}

type ArticleSearchQuery struct {
	ListQuery
	TagID            *TagID
	Published        *bool
	PublishedAtStart *time.Time
	PublishedAtEnd   *time.Time
	CreatedAtStart   *time.Time
	CreatedAtEnd     *time.Time
}
