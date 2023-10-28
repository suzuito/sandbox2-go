package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	gomock "go.uber.org/mock/gomock"
)

func TestGenerateSitemap(t *testing.T) {
	testCases := []struct {
		desc            string
		inputSiteOrigin string
		expectedURLs    XMLURLSet
		expectedError   string
		setUp           func(repositoryArticle *MockRepositoryArticle)
	}{
		{
			desc:            `Successful generation of sitemap. Loop 0`,
			inputSiteOrigin: "https://example.com",
			expectedURLs: XMLURLSet{
				URLs: []XMLURL{
					{
						Loc:     "https://example.com/",
						Lastmod: "2023-02-20",
					},
				},
				XMLNSXsi:          "http://www.w3.org/2001/XMLSchema-instance",
				XMLNS:             "http://www.sitemaps.org/schemas/sitemap/0.9",
				XsiSchemaLocation: "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd",
			},
			setUp: func(repositoryArticle *MockRepositoryArticle) {
				primaryKeys := []entity.ArticlePrimaryKey{}
				hasNext := false

				repositoryArticle.EXPECT().SearchArticles(
					gomock.Any(),
					SearchArticlesQuery{
						Offset:    0,
						Limit:     7,
						SortField: SearchArticlesQuerySortFieldDate,
						SortOrder: SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).SetArg(2, primaryKeys).SetArg(3, hasNext).Return(nil).Times(1)
			},
		},
		{
			desc:            `Successful generation of sitemap. Loop 1`,
			inputSiteOrigin: "https://example.com",
			expectedURLs: XMLURLSet{
				URLs: []XMLURL{
					{
						Loc:     "https://example.com/",
						Lastmod: "2023-02-20",
					},
					{
						Loc:     "https://example.com/articles/articleID1",
						Lastmod: "2023-05-01",
					},
					{
						Loc:     "https://example.com/articles/articleID2",
						Lastmod: "2023-04-29",
					},
				},
				XMLNSXsi:          "http://www.w3.org/2001/XMLSchema-instance",
				XMLNS:             "http://www.sitemaps.org/schemas/sitemap/0.9",
				XsiSchemaLocation: "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd",
			},
			setUp: func(repositoryArticle *MockRepositoryArticle) {
				primaryKeys := []entity.ArticlePrimaryKey{
					{ArticleID: entity.ArticleID("articleID1"), Version: 1},
					{ArticleID: entity.ArticleID("articleID2"), Version: 1},
				}
				hasNext := false

				repositoryArticle.EXPECT().SearchArticles(
					gomock.Any(),
					SearchArticlesQuery{
						Offset:    0,
						Limit:     7,
						SortField: SearchArticlesQuerySortFieldDate,
						SortOrder: SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).SetArg(2, primaryKeys).SetArg(3, hasNext).Return(nil).Times(1)

				articles := []entity.Article{
					{
						ID:        "articleID1",
						Version:   1,
						UpdatedAt: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "articleID2",
						Version:   1,
						UpdatedAt: time.Date(2023, 4, 29, 0, 0, 0, 0, time.UTC),
					},
				}
				repositoryArticle.EXPECT().GetArticlesByPrimaryKey(
					gomock.Any(),
					primaryKeys,
					SearchArticlesQuerySortFieldVersion,
					SortOrderDesc,
					gomock.Any(),
				).SetArg(4, articles).Return(nil).Times(1)
			},
		},
		{
			desc:            `Successful generation of sitemap. Loop 2`,
			inputSiteOrigin: "https://example.com",
			expectedURLs: XMLURLSet{
				URLs: []XMLURL{
					{
						Loc:     "https://example.com/",
						Lastmod: "2023-02-20",
					},
					{
						Loc:     "https://example.com/articles/articleID1",
						Lastmod: "2023-05-01",
					},
					{
						Loc:     "https://example.com/articles/articleID2",
						Lastmod: "2023-04-29",
					},
					{
						Loc:     "https://example.com/articles/articleID3",
						Lastmod: "2023-04-28",
					},
					{
						Loc:     "https://example.com/articles/articleID4",
						Lastmod: "2023-04-27",
					},
				},
				XMLNSXsi:          "http://www.w3.org/2001/XMLSchema-instance",
				XMLNS:             "http://www.sitemaps.org/schemas/sitemap/0.9",
				XsiSchemaLocation: "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd",
			},
			setUp: func(repositoryArticle *MockRepositoryArticle) {
				// Loop1
				primaryKeys1 := []entity.ArticlePrimaryKey{
					{ArticleID: entity.ArticleID("articleID1"), Version: 1},
					{ArticleID: entity.ArticleID("articleID2"), Version: 1},
				}
				hasNext1 := true
				repositoryArticle.EXPECT().SearchArticles(
					gomock.Any(),
					SearchArticlesQuery{
						Offset:    0,
						Limit:     7,
						SortField: SearchArticlesQuerySortFieldDate,
						SortOrder: SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).SetArg(2, primaryKeys1).SetArg(3, hasNext1).Return(nil).Times(1)
				articles1 := []entity.Article{
					{
						ID:        "articleID1",
						Version:   1,
						UpdatedAt: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "articleID2",
						Version:   1,
						UpdatedAt: time.Date(2023, 4, 29, 0, 0, 0, 0, time.UTC),
					},
				}
				repositoryArticle.EXPECT().GetArticlesByPrimaryKey(
					gomock.Any(),
					primaryKeys1,
					SearchArticlesQuerySortFieldVersion,
					SortOrderDesc,
					gomock.Any(),
				).SetArg(4, articles1).Return(nil).Times(1)
				// Loop2
				primaryKeys2 := []entity.ArticlePrimaryKey{
					{ArticleID: entity.ArticleID("articleID3"), Version: 1},
					{ArticleID: entity.ArticleID("articleID4"), Version: 1},
				}
				hasNext2 := false
				repositoryArticle.EXPECT().SearchArticles(
					gomock.Any(),
					SearchArticlesQuery{
						Offset:    7,
						Limit:     7,
						SortField: SearchArticlesQuerySortFieldDate,
						SortOrder: SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).SetArg(2, primaryKeys2).SetArg(3, hasNext2).Return(nil).Times(1)
				articles2 := []entity.Article{
					{
						ID:        "articleID3",
						Version:   1,
						UpdatedAt: time.Date(2023, 4, 28, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "articleID4",
						Version:   1,
						UpdatedAt: time.Date(2023, 4, 27, 0, 0, 0, 0, time.UTC),
					},
				}
				repositoryArticle.EXPECT().GetArticlesByPrimaryKey(
					gomock.Any(),
					primaryKeys2,
					SearchArticlesQuerySortFieldVersion,
					SortOrderDesc,
					gomock.Any(),
				).SetArg(4, articles2).Return(nil).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepositoryArticle := NewMockRepositoryArticle(ctrl)
			tc.setUp(mockRepositoryArticle)

			usecase := UsecaseImpl{
				RepositoryArticle: mockRepositoryArticle,
			}

			urls := XMLURLSet{}
			err := usecase.GenerateSitemap(ctx, tc.inputSiteOrigin, &urls)
			test_helper.AssertErrorAs(t, tc.expectedError, err)

			if err == nil {
				assert.Equal(t, tc.expectedURLs, urls)
			}
		})
	}
}
