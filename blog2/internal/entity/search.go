package entity

import (
	"time"

	"github.com/suzuito/sandbox2-go/blog/entity"
)

type ListQuery struct {
	Offset *int
	Limit  *int
}

type ArticleSearchQuery struct {
	ListQuery
	TagID            *entity.TagID
	Published        *bool
	PublishedAtStart *time.Time
	PublishedAtEnd   *time.Time
	CreatedAtStart   *time.Time
	CreatedAtEnd     *time.Time
}
