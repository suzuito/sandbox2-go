package entity

import (
	"time"
)

type ArticleID string

type Article struct {
	ID          ArticleID  `json:"id"`
	Title       string     `json:"title"`
	Published   bool       `json:"published"`
	PublishedAt *time.Time `json:"publishedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	Tags        []Tag      `json:"tags"`
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

func (t *Article) GetTagIDs() []TagID {
	tagIDs := make([]TagID, len(t.Tags))
	for i, tag := range t.Tags {
		tagIDs[i] = tag.ID
	}
	return tagIDs
}
