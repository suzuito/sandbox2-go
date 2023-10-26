package bmysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryArticle struct {
	DB *gorm.DB
}

func (t *RepositoryArticle) GetArticlesByArticleSourceID(
	ctx context.Context,
	articleSourceID entity.ArticleSourceID,
	articles *[]entity.Article,
) error {
	tMappingArticlesSourceArticles := []tableMappingArticlesSourceArticles{}
	if err := t.DB.
		Where("article_source_id = ?", articleSourceID).
		Find(&tMappingArticlesSourceArticles).
		Error; err != nil {
		return err
	}
	primaryKeys := []entity.ArticlePrimaryKey{}
	for _, t := range tMappingArticlesSourceArticles {
		primaryKeys = append(primaryKeys, entity.ArticlePrimaryKey{
			ArticleID: t.ArticleID,
			Version:   t.ArticleVersion,
		})
	}
	if len(primaryKeys) <= 0 {
		return nil
	}
	return t.GetArticlesByPrimaryKey(ctx, primaryKeys, usecase.SearchArticlesQuerySortFieldVersion, usecase.SortOrderDesc, articles)
}

func (t *RepositoryArticle) GetArticlesByPrimaryKey(
	ctx context.Context,
	primaryKeys []entity.ArticlePrimaryKey,
	sortField usecase.SearchArticlesQuerySortField,
	sortOrder usecase.SortOrder,
	articles *[]entity.Article,
) error {
	if sortField == "" {
		return fmt.Errorf("sortField must not be empty")
	}
	if sortOrder == "" {
		return fmt.Errorf("sortOrder must not be empty")
	}
	if len(primaryKeys) <= 0 {
		return nil
	}
	tArticles := []tableArticle{}
	db := t.DB.WithContext(ctx)
	for _, primaryKey := range primaryKeys {
		db = db.Or("id = ? AND version = ?", primaryKey.ArticleID, primaryKey.Version)
	}
	if err := db.Order(fmt.Sprintf("%s %s", sortField, sortOrder)).
		Order("id desc").
		Preload("Tags").
		Preload("MappingArticlesSourceArticles").
		Find(&tArticles).Error; err != nil {
		return err
	}
	for _, tArticle := range tArticles {
		*articles = append(*articles, *tArticle.ToEntity())
	}
	return nil
}

func (t *RepositoryArticle) GetArticleByPrimaryKey(
	ctx context.Context,
	primaryKey entity.ArticlePrimaryKey,
	article *entity.Article,
) error {
	tArticle := tableArticle{}
	if err := t.DB.
		Where("id = ? AND version = ?", primaryKey.ArticleID, primaryKey.Version).
		Preload("Tags").
		Preload("MappingArticlesSourceArticles").
		First(&tArticle).
		Error; err != nil {
		return err
	}
	*article = *tArticle.ToEntity()
	return nil
}

func (t *RepositoryArticle) GetLatestArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	article *entity.Article,
) error {
	tArticle := tableArticle{}
	db := t.DB.WithContext(ctx)
	if err := db.
		Where("id = ?", articleID).
		Order("version desc").
		Order("id desc").
		Preload("Tags").
		Preload("MappingArticlesSourceArticles").
		First(&tArticle).Error; err != nil {
		return err
	}
	*article = *tArticle.ToEntity()
	return nil
}

func (t *RepositoryArticle) GetArticlesByID(
	ctx context.Context,
	articleID entity.ArticleID,
	articles *[]entity.Article,
) error {
	tArticles := []tableArticle{}
	if err := t.DB.
		Where("id = ?", articleID).
		Order("version desc").
		Order("id desc").
		Preload("Tags").
		Preload("MappingArticlesSourceArticles").
		Find(&tArticles).Error; err != nil {
		return err
	}
	for _, t := range tArticles {
		*articles = append(*articles, *t.ToEntity())
	}
	return nil
}

func (t *RepositoryArticle) SearchArticles(
	ctx context.Context,
	query usecase.SearchArticlesQuery,
	articlePrimaryKeys *[]entity.ArticlePrimaryKey,
	hasNext *bool,
) error {
	if query.SortField == "" {
		return fmt.Errorf("sortField must not be empty")
	}
	if query.SortOrder == "" {
		return fmt.Errorf("sortOrder must not be empty")
	}
	indices := []tableArticleSearchIndex{}
	db := t.DB.Order(fmt.Sprintf("%s %s", query.SortField, query.SortOrder))
	if len(query.Tags) > 0 {
		db = db.Where("match(tags) against(? in boolean mode)", strings.Join(query.Tags, " "))
	}
	db = db.Order("article_id desc")
	db = db.Limit(query.Limit + 1)
	db = db.Offset(query.Offset)
	if err := db.
		Find(&indices).
		Error; err != nil {
		return err
	}
	*hasNext = len(indices) > query.Limit
	for i, index := range indices {
		if i >= query.Limit {
			continue
		}
		*articlePrimaryKeys = append(*articlePrimaryKeys, entity.ArticlePrimaryKey{
			ArticleID: index.ArticleID,
			Version:   index.CurrentArticleVersion,
		})
	}
	return nil
}

func (t *RepositoryArticle) SetArticle(
	ctx context.Context,
	article *entity.Article,
) error {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		tArticle := newTableArticle(article)
		if err := tx.Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).Create(tArticle).Error; err != nil {
			return err
		}
		if err := tx.Where("article_id", article.ID).Delete(&tableMappingArticlesTags{}).Error; err != nil {
			return err
		}
		for _, tag := range article.Tags {
			tTag := newTableTag(&tag)
			if err := tx.Clauses(
				clause.OnConflict{
					DoNothing: true,
				},
			).Create(tTag).Error; err != nil {
				return err
			}
		}
		for _, tag := range article.Tags {
			tMap := tableMappingArticlesTags{
				ArticleID:      article.ID,
				ArticleVersion: article.Version,
				TagID:          tag.ID,
			}
			if err := tx.Clauses(
				clause.OnConflict{
					DoNothing: true,
				},
			).Create(tMap).Error; err != nil {
				return err
			}
		}
		{
			if err := tx.Clauses(
				clause.OnConflict{
					UpdateAll: true,
				},
			).Create(tArticle.MappingArticlesSourceArticles).Error; err != nil {
				return err
			}

		}
		return nil
	})
}

func (t *RepositoryArticle) SetArticleSearchIndex(
	ctx context.Context,
	article *entity.Article,
) error {
	tagStrings := []string{}
	for _, t := range article.Tags {
		tagStrings = append(
			tagStrings,
			fmt.Sprintf("%s", t.ID),
		)
	}
	index := tableArticleSearchIndex{
		ArticleID:             article.ID,
		CurrentArticleVersion: article.Version,
		Date:                  article.Date,
		Tags:                  strings.Join(tagStrings, " "),
	}
	if err := t.DB.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		},
	).Create(index).Error; err != nil {
		return err
	}
	return nil
}
