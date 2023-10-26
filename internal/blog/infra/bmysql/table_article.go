package bmysql

import (
	"time"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"gorm.io/gorm"
)

type tableArticle struct {
	ID                            entity.ArticleID `gorm:"primaryKey"`
	Version                       int32            `gorm:"primaryKey"`
	Title                         string
	Description                   string
	PublishedAt                   *time.Time
	Date                          time.Time
	Tags                          []tableTag                         `gorm:"many2many:mapping_articles_tags;joinForeignKey:article_id,article_version;joinReferences:tag_id;"`
	MappingArticlesSourceArticles tableMappingArticlesSourceArticles `gorm:"foreignKey:article_id,article_version;"`
	gorm.Model
}

func (t *tableArticle) TableName() string {
	return "articles"
}

func (t *tableArticle) ToEntity() *entity.Article {
	tags := []entity.Tag{}
	for _, t := range t.Tags {
		tags = append(tags, entity.Tag{ID: t.ID})
	}
	articleSource := entity.ArticleSource{
		ID:      t.MappingArticlesSourceArticles.ArticleSourceID,
		Version: t.MappingArticlesSourceArticles.ArticleSourceVersion,
		Meta:    t.MappingArticlesSourceArticles.Meta,
	}
	return &entity.Article{
		ID:            t.ID,
		Version:       t.Version,
		Title:         t.Title,
		Description:   t.Description,
		PublishedAt:   t.PublishedAt,
		UpdatedAt:     t.UpdatedAt,
		Date:          t.Date,
		Tags:          tags,
		ArticleSource: articleSource,
	}
}

func newTableArticle(article *entity.Article) *tableArticle {
	return &tableArticle{
		Model: gorm.Model{
			UpdatedAt: article.UpdatedAt,
		},
		MappingArticlesSourceArticles: tableMappingArticlesSourceArticles{
			ArticleID:            article.ID,
			ArticleVersion:       article.Version,
			ArticleSourceID:      article.ArticleSource.ID,
			ArticleSourceVersion: article.ArticleSource.Version,
			Meta:                 article.ArticleSource.Meta,
		},
		ID:          article.ID,
		Version:     article.Version,
		Title:       article.Title,
		Date:        article.Date,
		Description: article.Description,
		PublishedAt: article.PublishedAt,
	}
}
