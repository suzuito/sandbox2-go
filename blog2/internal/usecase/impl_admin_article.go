package usecase

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticle struct {
	MarkdownBody string
}

func (t *Impl) GetAdminArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*DTOGetAdminArticle, error) {
	markdownBodyBuffer := bytes.NewBufferString("")
	if err := t.StorageArticle.GetArticle(ctx, articleID, markdownBodyBuffer); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticle{
		MarkdownBody: markdownBodyBuffer.String(),
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
	article, err := t.RepositoryArticle.UpdateArticle(ctx, articleID, title, nil, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPutAdminArticle{
		Article: article,
	}, nil
}

func (t *Impl) PutAdminArticleMarkdown(ctx context.Context, articleID entity.ArticleID, markdownBodyReader io.Reader) error {
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		return terrors.Wrapf("Document %s is not found", articleID)
	}
	if err := t.StorageArticle.PutArticle(ctx, articleID, markdownBodyReader); err != nil {
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
	_, err := t.RepositoryArticle.UpdateArticle(
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
	_, err := t.RepositoryArticle.UpdateArticle(
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
	_, err := t.RepositoryArticle.UpdateArticleTags(ctx, articleID, add, delete)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
