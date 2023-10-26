package bmysql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
	"github.com/suzuito/sandbox2-go/internal/common/test_helper"
	"gorm.io/gorm"
)

func TestGetArticlesByArticleSourceID(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc                 string
		inputArticleSourceID entity.ArticleSourceID
		expectedArticles     []entity.Article
		expectedError        string
		setUp                func(t *testing.T, db *gorm.DB)
	}{
		{
			desc:                 "Suceess",
			inputArticleSourceID: "src01",
			setUp:                setUpTestData,
			expectedArticles: []entity.Article{
				{
					ID:          "article01",
					Version:     2,
					Title:       "articleTitle01",
					Description: "articleDescription01",
					Tags: []entity.Tag{
						{ID: "tag02"},
					},
					ArticleSource: entity.ArticleSource{
						ID:      "src01",
						Version: "sha02",
						Meta: entity.ArticleSourceMeta{
							URL: "https://www.example.com",
						},
					},
					Date:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
				{
					ID:          "article01",
					Version:     1,
					Title:       "articleTitle01",
					Description: "articleDescription01",
					Tags: []entity.Tag{
						{ID: "tag01"},
					},
					ArticleSource: entity.ArticleSource{
						ID:      "src01",
						Version: "sha01",
						Meta: entity.ArticleSourceMeta{
							URL: "https://www.example.com",
						},
					},
					Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			articles := []entity.Article{}
			err := repo.GetArticlesByArticleSourceID(ctx, tC.inputArticleSourceID, &articles)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticles, articles)
			}
		})
	}
}

func TestGetArticleByPrimaryKey(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc                        string
		inputPrimaryKey             entity.ArticlePrimaryKey
		expectedArticle             entity.Article
		expectedError               string
		expectedErrorRecordNotFound bool
		setUp                       func(t *testing.T, db *gorm.DB)
	}{
		{
			desc: "Suceess 1",
			inputPrimaryKey: entity.ArticlePrimaryKey{
				ArticleID: "article01",
				Version:   2,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticle: entity.Article{
				ID:          "article01",
				Version:     2,
				Title:       "articleTitle01",
				Description: "articleDescription01",
				Tags: []entity.Tag{
					{ID: "tag02"},
				},
				ArticleSource: entity.ArticleSource{
					ID:      "src01",
					Version: "sha02",
					Meta: entity.ArticleSourceMeta{
						URL: "https://www.example.com",
					},
				},
				Date:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
		{
			desc: "Suceess 2",
			inputPrimaryKey: entity.ArticlePrimaryKey{
				ArticleID: "article02",
				Version:   1,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticle: entity.Article{
				ID:            "article02",
				Version:       1,
				Title:         "articleTitle02",
				Description:   "articleDescription02",
				Tags:          []entity.Tag{},
				ArticleSource: entity.ArticleSource{},
				Date:          time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
				PublishedAt:   timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				UpdatedAt:     time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
		{
			desc: "Failed: record not found",
			inputPrimaryKey: entity.ArticlePrimaryKey{
				ArticleID: "article01",
				Version:   99,
			},
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedError:               "record not found",
			expectedErrorRecordNotFound: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			article := entity.Article{}
			err := repo.GetArticleByPrimaryKey(ctx, tC.inputPrimaryKey, &article)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedErrorRecordNotFound {
				assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
			}
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticle, article)
			}
		})
	}
}

func TestGetArticlesByPrimaryKey(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc                        string
		inputPrimaryKey             []entity.ArticlePrimaryKey
		inputSortField              usecase.SearchArticlesQuerySortField
		inputSortOrder              usecase.SortOrder
		expectedArticles            []entity.Article
		expectedError               string
		expectedErrorRecordNotFound bool
		setUp                       func(t *testing.T, db *gorm.DB)
	}{
		{
			desc:            "Success: return empty record",
			inputPrimaryKey: []entity.ArticlePrimaryKey{},
			inputSortField:  usecase.SearchArticlesQuerySortFieldDate,
			inputSortOrder:  usecase.SortOrderDesc,
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticles: []entity.Article{},
		},
		{
			desc: "Suceess: return multiple record",
			inputPrimaryKey: []entity.ArticlePrimaryKey{
				{ArticleID: "article01", Version: 1},
				{ArticleID: "article01", Version: 2},
				{ArticleID: "article02", Version: 1},
			},
			inputSortField: usecase.SearchArticlesQuerySortFieldDate,
			inputSortOrder: usecase.SortOrderDesc,
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticles: []entity.Article{
				{
					ID:            "article02",
					Version:       1,
					Title:         "articleTitle02",
					Description:   "articleDescription02",
					Tags:          []entity.Tag{},
					ArticleSource: entity.ArticleSource{},
					Date:          time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
					PublishedAt:   timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:     time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
				{
					ID:          "article01",
					Version:     2,
					Title:       "articleTitle01",
					Description: "articleDescription01",
					Tags: []entity.Tag{
						{ID: "tag02"},
					},
					ArticleSource: entity.ArticleSource{
						ID:      "src01",
						Version: "sha02",
						Meta: entity.ArticleSourceMeta{
							URL: "https://www.example.com",
						},
					},
					Date:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
				{
					ID:          "article01",
					Version:     1,
					Title:       "articleTitle01",
					Description: "articleDescription01",
					Tags: []entity.Tag{
						{ID: "tag01"},
					},
					ArticleSource: entity.ArticleSource{
						ID:      "src01",
						Version: "sha01",
						Meta: entity.ArticleSourceMeta{
							URL: "https://www.example.com",
						},
					},
					Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
			},
		},
		{
			desc: "Suceess: return empty record",
			inputPrimaryKey: []entity.ArticlePrimaryKey{
				{ArticleID: "foo", Version: 2},
				{ArticleID: "bar", Version: 3},
				{ArticleID: "boo", Version: 1},
			},
			inputSortField: usecase.SearchArticlesQuerySortFieldDate,
			inputSortOrder: usecase.SortOrderAsc,
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticles: []entity.Article{},
		},
		{
			desc:             "Failed: sortField is empty",
			inputSortField:   "",
			setUp:            func(t *testing.T, db *gorm.DB) {},
			expectedArticles: []entity.Article{},
			expectedError:    "sortField must not be empty",
		},
		{
			desc:             "Failed: sortOrder is empty",
			inputSortField:   usecase.SearchArticlesQuerySortFieldDate,
			inputSortOrder:   "",
			setUp:            func(t *testing.T, db *gorm.DB) {},
			expectedArticles: []entity.Article{},
			expectedError:    "sortOrder must not be empty",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			articles := []entity.Article{}
			err := repo.GetArticlesByPrimaryKey(
				ctx,
				tC.inputPrimaryKey,
				tC.inputSortField,
				tC.inputSortOrder,
				&articles,
			)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedErrorRecordNotFound {
				assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
			}
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticles, articles)
			}
		})
	}
}

func TestGetLatestArticle(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc                        string
		inputArticleID              entity.ArticleID
		expectedArticle             entity.Article
		expectedError               string
		expectedErrorRecordNotFound bool
		setUp                       func(t *testing.T, db *gorm.DB)
	}{
		{
			desc:           "Suceess 1",
			inputArticleID: "article01",
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticle: entity.Article{
				ID:          "article01",
				Version:     2,
				Title:       "articleTitle01",
				Description: "articleDescription01",
				Tags: []entity.Tag{
					{ID: "tag02"},
				},
				ArticleSource: entity.ArticleSource{
					ID:      "src01",
					Version: "sha02",
					Meta: entity.ArticleSourceMeta{
						URL: "https://www.example.com",
					},
				},
				Date:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
		{
			desc:           "Failed: record not found",
			inputArticleID: "article99",
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedError:               "record not found",
			expectedErrorRecordNotFound: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			article := entity.Article{}
			err := repo.GetLatestArticle(ctx, tC.inputArticleID, &article)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedErrorRecordNotFound {
				assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
			}
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticle, article)
			}
		})
	}
}

func TestGetArticlesByID(t *testing.T) {
	db, err := newDB()
	if err != nil {
		t.Errorf("newDB is failed : %+v", err)
		return
	}
	repo := RepositoryArticle{DB: db}
	testCases := []struct {
		desc                        string
		inputArticleID              entity.ArticleID
		expectedArticles            []entity.Article
		expectedError               string
		expectedErrorRecordNotFound bool
		setUp                       func(t *testing.T, db *gorm.DB)
	}{
		{
			desc:           "Suceess: return multiple record",
			inputArticleID: "article01",
			setUp: func(t *testing.T, db *gorm.DB) {
				setUpTestData(t, db)
			},
			expectedArticles: []entity.Article{
				{
					ID:          "article01",
					Version:     2,
					Title:       "articleTitle01",
					Description: "articleDescription01",
					Tags: []entity.Tag{
						{ID: "tag02"},
					},
					ArticleSource: entity.ArticleSource{
						ID:      "src01",
						Version: "sha02",
						Meta: entity.ArticleSourceMeta{
							URL: "https://www.example.com",
						},
					},
					Date:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
				{
					ID:          "article01",
					Version:     1,
					Title:       "articleTitle01",
					Description: "articleDescription01",
					Tags: []entity.Tag{
						{ID: "tag01"},
					},
					ArticleSource: entity.ArticleSource{
						ID:      "src01",
						Version: "sha01",
						Meta: entity.ArticleSourceMeta{
							URL: "https://www.example.com",
						},
					},
					Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					PublishedAt: timeDate(2023, 1, 2, 3, 4, 5, 0, time.UTC),
					UpdatedAt:   time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC),
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer cleanUpDB(t, db)
			ctx := context.Background()
			tC.setUp(t, db)
			articles := []entity.Article{}
			err := repo.GetArticlesByID(
				ctx,
				tC.inputArticleID,
				&articles,
			)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedErrorRecordNotFound {
				assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
			}
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticles, articles)
			}
		})
	}
}
