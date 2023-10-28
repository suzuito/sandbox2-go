package bmysql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
	"gorm.io/gorm"
)

func setUpTestDataForSearch(
	t *testing.T,
	db *gorm.DB,
	indices []tableArticleSearchIndex,
) {
	for _, index := range indices {
		mustCreate(t, db, tableArticle{
			ID:      index.ArticleID,
			Version: index.CurrentArticleVersion,
			Date:    time.Now(),
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
		mustCreate(t, db, index)
	}
}

func TestSearchArticles(t *testing.T) {
	indices := []tableArticleSearchIndex{
		{
			ArticleID:             "article01",
			CurrentArticleVersion: 1,
			Date:                  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			Tags:                  "tag01 tag02",
		},
		{
			ArticleID:             "article02",
			CurrentArticleVersion: 1,
			Date:                  time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
			Tags:                  "tag02 tag03",
		},
		{
			ArticleID:             "article03",
			CurrentArticleVersion: 1,
			Date:                  time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
			Tags:                  "tag03 tag04",
		},
		{
			ArticleID:             "article04",
			CurrentArticleVersion: 1,
			Date:                  time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			Tags:                  "tag04 tag05",
		},
		{
			ArticleID:             "article05",
			CurrentArticleVersion: 1,
			Date:                  time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC),
			Tags:                  "tag05 tag06",
		},
	}
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc                       string
		inputQuery                 usecase.SearchArticlesQuery
		expectedArticlePrimaryKeys []entity.ArticlePrimaryKey
		expectedHasNext            bool
		expectedError              string
		setUp                      func(t *testing.T, db *gorm.DB)
	}{
		{
			desc: `Success check pager. Check 1st`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderDesc,
				Offset:    0,
				Limit:     2,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{
				{ArticleID: "article05", Version: 1},
				{ArticleID: "article04", Version: 1},
			},
			expectedHasNext: true,
		},
		{
			desc: `Success check pager. Check 2nd`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderDesc,
				Offset:    2,
				Limit:     2,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{
				{ArticleID: "article03", Version: 1},
				{ArticleID: "article02", Version: 1},
			},
			expectedHasNext: true,
		},
		{
			desc: `Success check pager. Check last`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderDesc,
				Offset:    4,
				Limit:     2,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{
				{ArticleID: "article01", Version: 1},
			},
			expectedHasNext: false,
		},
		{
			desc: `Success check pager. empty`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderDesc,
				Offset:    999,
				Limit:     2,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{},
			expectedHasNext:            false,
		},
		{
			desc: `Success check order. date asc`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderAsc,
				Offset:    2,
				Limit:     2,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{
				{ArticleID: "article03", Version: 1},
				{ArticleID: "article04", Version: 1},
			},
			expectedHasNext: true,
		},
		{
			desc: `Success. cond tags.`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderDesc,
				Tags: []string{
					"tag02",
				},
				Offset: 0,
				Limit:  100,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{
				{ArticleID: "article02", Version: 1},
				{ArticleID: "article01", Version: 1},
			},
			expectedHasNext: false,
		},
		{
			desc: `Success. multiple cond tags.`,
			inputQuery: usecase.SearchArticlesQuery{
				SortField: usecase.SearchArticlesQuerySortFieldDate,
				SortOrder: usecase.SortOrderDesc,
				Tags: []string{
					"tag02",
					"tag04",
				},
				Offset: 0,
				Limit:  100,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestDataForSearch(
					t,
					db,
					indices,
				)
			},
			expectedArticlePrimaryKeys: []entity.ArticlePrimaryKey{
				{ArticleID: "article04", Version: 1},
				{ArticleID: "article03", Version: 1},
				{ArticleID: "article02", Version: 1},
				{ArticleID: "article01", Version: 1},
			},
			expectedHasNext: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			articlePrimaryKeys := []entity.ArticlePrimaryKey{}
			hasNext := false
			err := repo.SearchArticles(
				ctx,
				tC.inputQuery,
				&articlePrimaryKeys,
				&hasNext,
			)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticlePrimaryKeys, articlePrimaryKeys)
				assert.Equal(t, tC.expectedHasNext, hasNext)
			}
		})
	}
}

func TestSetArticleSearchIndex(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc          string
		inputArticle  entity.Article
		expectedError string
		setUp         func(t *testing.T, db *gorm.DB)
		assert        func(t *testing.T, db *gorm.DB)
	}{
		{
			desc: "Success",
			inputArticle: entity.Article{
				ID:          "article01",
				Version:     1,
				Title:       "title01",
				Description: "desc01",
				Date:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				Tags: []entity.Tag{
					{ID: "tag01"},
					{ID: "tag02"},
				},
				ArticleSource: entity.ArticleSource{
					ID:      "source01",
					Version: "v01",
					Meta: entity.ArticleSourceMeta{
						URL: "https://www.example.com",
					},
				},
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				mustCreate(t, db, tableArticle{
					ID:      "article01",
					Version: 1,
					Date:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					Model: gorm.Model{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				})
			},
			assert: func(t *testing.T, db *gorm.DB) {

			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			err := repo.SetArticleSearchIndex(
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
