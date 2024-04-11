package service

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreateTestData(ctx context.Context) error {
	// Create tags
	tagsCreated := []*entity.Tag{}
	for i := range 30 {
		name := fmt.Sprintf("タグ-%d", i)
		tag, err := t.RepositoryArticle.CreateTag(ctx, name)
		if err != nil {
			return terrors.Wrap(err)
		}
		tagsCreated = append(tagsCreated, tag)
	}
	// Create articles
	nowBase := time.Now()
	articlesCreated := []*entity.Article{}
	for i := range 50 {
		articleID := entity.ArticleID(fmt.Sprintf("article-%d", i))
		if err := t.StorageArticle.PutArticle(ctx, articleID, bytes.NewBuffer([]byte{})); err != nil {
			return terrors.Wrap(err)
		}
		article, err := t.RepositoryArticle.CreateArticle(ctx, articleID, nowBase.Add(time.Duration(i*-24)*time.Hour))
		if err != nil {
			return terrors.Wrap(err)
		}
		articlesCreated = append(articlesCreated, article)
	}
	trueValue := true
	for i, article := range articlesCreated {
		title := fmt.Sprintf("これはテスト記事ですよん-%d", i)
		var published *bool
		var publishedAt *time.Time
		if i%5 == 0 {
			published = &trueValue
			publishedAtValue := nowBase.Add(time.Duration(i*-24) * time.Hour)
			publishedAt = &publishedAtValue
		}
		_, err := t.RepositoryArticle.UpdateArticle(
			ctx,
			article.ID,
			&title,
			published,
			publishedAt,
		)
		if err != nil {
			return terrors.Wrap(err)
		}
		attachedTagIDs := []entity.TagID{}
		for _, tag := range tagsCreated {
			if rand.Int()%2 != 0 {
				continue
			}
			attachedTagIDs = append(attachedTagIDs, tag.ID)
		}
		_, err = t.RepositoryArticle.UpdateArticleTags(
			ctx,
			article.ID,
			attachedTagIDs,
			[]entity.TagID{},
		)
		if err != nil {
			return terrors.Wrap(err)
		}
	}
	return nil
}
