package usecase

import (
	"context"
	"io"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticle struct {
	MarkdownBody string
	HTMLBody     string
}

func (t *Impl) GetAdminArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*DTOGetAdminArticle, error) {
	markdownBody, htmlBody, err := t.S.GetArticleBody(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticle{
		MarkdownBody: markdownBody,
		HTMLBody:     htmlBody,
	}, nil
}

type DTOPutAdminArticle struct {
	Article *entity.Article
}

func (t *Impl) PutAdminArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	title *string,
) (*DTOPutAdminArticle, error) {
	article, err := t.S.PutArticle(ctx, articleID, title, nil, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPutAdminArticle{
		Article: article,
	}, nil
}

func (t *Impl) PutAdminArticleMarkdown(ctx context.Context, articleID entity.ArticleID, markdownBodyReader io.Reader) error {
	if err := t.S.PutArticleMarkdown(ctx, articleID, markdownBodyReader); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) PostAdminArticlePublish(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	valueTrue := true
	valueNow := time.Now()
	_, err := t.S.PutArticle(
		ctx,
		articleID,
		nil,
		&valueTrue,
		&valueNow,
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) DeleteAdminArticlePublish(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	valueFalse := false
	_, err := t.S.PutArticle(
		ctx,
		articleID,
		nil,
		&valueFalse,
		nil,
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) PostAdminArticleEditTags(
	ctx context.Context,
	articleID entity.ArticleID,
	add []entity.TagID,
	delete []entity.TagID,
) error {
	if len(add) <= 0 && len(delete) <= 0 {
		return nil
	}
	_, err := t.S.UpdateArticleTags(ctx, articleID, add, delete)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
