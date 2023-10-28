package bmysql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"gorm.io/gorm"
)

func TestSetArticle(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc             string
		inputArticle     entity.Article
		expectedArticles []entity.Article
		expectedError    string
		setUp            func(t *testing.T, db *gorm.DB)
		assert           func(t *testing.T, db *gorm.DB)
	}{
		{
			desc: "Success create new entity",
			inputArticle: entity.Article{
				ID:          "article101",
				Version:     1,
				Title:       "articleTitle101",
				Description: "articleDescription101",
				Tags: []entity.Tag{
					{ID: "tag101"},
				},
				ArticleSource: entity.ArticleSource{ID: "src101", Version: "sha101", Meta: entity.ArticleSourceMeta{URL: "https://www.example.com"}},
				Date:          time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				PublishedAt:   timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			assert: func(t *testing.T, db *gorm.DB) {
				article := tableArticle{ID: "article101", Version: 1}
				db.First(&article)
				assert.Equal(t, "article101", string(article.ID))
				assert.Equal(t, int32(1), article.Version)
				assert.Equal(t, "articleTitle101", article.Title)
				assert.Equal(t, "articleDescription101", article.Description)
				assert.Equal(t, time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), article.Date)
				assert.Equal(t, timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC), article.PublishedAt)
				tag := tableTag{}
				db.Where("id = ?", "tag101").First(&tag)
				assert.Equal(t, "tag101", string(tag.ID))
				tags := []tableMappingArticlesTags{}
				db.Where("article_id = ?", "article101").Find(&tags)
				assert.Equal(t, 1, len(tags))
				if len(tags) > 0 {
					assert.Equal(t, "tag101", string(tags[0].TagID))
					assert.Equal(t, "article101", string(tags[0].ArticleID))
					assert.Equal(t, int32(1), tags[0].ArticleVersion)
				}
				sourceArticle := tableMappingArticlesSourceArticles{}
				db.Where("article_id = ?", "article101").First(&sourceArticle)
				assert.Equal(t, "article101", string(sourceArticle.ArticleID))
				assert.Equal(t, int32(1), sourceArticle.ArticleVersion)
				assert.Equal(t, "src101", string(sourceArticle.ArticleSourceID))
				assert.Equal(t, "sha101", sourceArticle.ArticleSourceVersion)
			},
		},
		{
			desc: "Success update existing entity",
			inputArticle: entity.Article{
				ID:          "article01",
				Version:     1,
				Title:       "articleTitle01Updated",
				Description: "articleDescription01Updated",
				Tags: []entity.Tag{
					{ID: "tag01Updated"},
				},
				ArticleSource: entity.ArticleSource{ID: "src01Updated", Version: "sha01Updated", Meta: entity.ArticleSourceMeta{URL: "https://www.example.com/updated"}},
				Date:          time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				PublishedAt:   timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			assert: func(t *testing.T, db *gorm.DB) {
				article := tableArticle{ID: "article01", Version: 1}
				db.First(&article)
				assert.Equal(t, "article01", string(article.ID))
				assert.Equal(t, int32(1), article.Version)
				assert.Equal(t, "articleTitle01Updated", article.Title)
				assert.Equal(t, "articleDescription01Updated", article.Description)
				assert.Equal(t, time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), article.Date)
				assert.Equal(t, timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC), article.PublishedAt)
				tag := tableTag{}
				db.Where("id = ?", "tag01Updated").First(&tag)
				assert.Equal(t, "tag01Updated", string(tag.ID))
				tags := []tableMappingArticlesTags{}
				db.Where("article_id = ?", "article01").Find(&tags)
				assert.Equal(t, 1, len(tags))
				if len(tags) > 0 {
					assert.Equal(t, "tag01Updated", string(tags[0].TagID))
					assert.Equal(t, "article01", string(tags[0].ArticleID))
					assert.Equal(t, int32(1), tags[0].ArticleVersion)
				}
				sourceArticle := tableMappingArticlesSourceArticles{}
				db.Where("article_id = ?", "article01").First(&sourceArticle)
				assert.Equal(t, "article01", string(sourceArticle.ArticleID))
				assert.Equal(t, int32(1), sourceArticle.ArticleVersion)
				assert.Equal(t, "src01Updated", string(sourceArticle.ArticleSourceID))
				assert.Equal(t, "sha01Updated", sourceArticle.ArticleSourceVersion)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			err := repo.SetArticle(
				ctx,
				&tC.inputArticle,
			)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				tC.assert(t, db)
			}
		})
	}
}
