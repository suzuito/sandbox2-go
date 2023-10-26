package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ArticleID string

type Article struct {
	ID          ArticleID `validate:"required"`
	Version     int32     `validate:"required"`
	Title       string    `validate:"required"`
	Description string
	// 記事に付与できるタグの最大数は10。
	// タグ数を無制限に付与できるようにしてしまうと、
	// DBにおける検索パフォーマンスが悪くなるため。
	Tags          []Tag `validate:"max=10"`
	ArticleSource ArticleSource
	Date          time.Time
	PublishedAt   *time.Time
	UpdatedAt     time.Time
}

func (t *Article) Validate() error {
	if err := validator.New().Struct(t); err != nil {
		return &ValidationError{err: err}
	}
	return nil
}

func (t *Article) DateString() string {
	return t.Date.Format("2006-01-02")
}

type ArticlePrimaryKey struct {
	ArticleID ArticleID
	Version   int32
}
