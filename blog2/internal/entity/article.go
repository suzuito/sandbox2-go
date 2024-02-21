package entity

import "time"

type ArticleID string

type Article struct {
	ID          ArticleID
	Title       string
	Summary     string
	Published   bool
	PublishedAt *time.Time
	Tags        []Tag
}
