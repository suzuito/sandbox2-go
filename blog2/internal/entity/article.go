package entity

import (
	"time"
)

type ArticleID string

type Article struct {
	ID          ArticleID
	Title       string
	Published   bool
	PublishedAt *time.Time
	Tags        []Tag
	Images      []ArticleImage
}

func (t *Article) StateString() string {
	if t.Published {
		return "公開中"
	}
	return "ドラフト"
}

func (t *Article) HasTag(tagID TagID) bool {
	for _, t := range t.Tags {
		if t.ID == tagID {
			return true
		}
	}
	return false
}
