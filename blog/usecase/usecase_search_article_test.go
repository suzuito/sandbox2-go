package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	gomock "go.uber.org/mock/gomock"
)

func TestSearchArticles(t *testing.T) {
	testCases := []struct {
		desc             string
		inputQuery       SearchArticlesQuery
		expectedArticles []entity.Article
		expectedHasNext  bool
		expectedError    string
		setUp            func(repositoryArticle *MockRepositoryArticle)
	}{
		{
			desc: `Successful. empty primary keys.`,
			inputQuery: SearchArticlesQuery{
				Offset:    0,
				Limit:     3,
				Tags:      []string{},
				SortField: SearchArticlesQuerySortFieldDate,
				SortOrder: SortOrderAsc,
			},
			expectedArticles: []entity.Article{},
			expectedHasNext:  false,
			setUp: func(repositoryArticle *MockRepositoryArticle) {
				primaryKeys := []entity.ArticlePrimaryKey{}
				hasNext := false
				repositoryArticle.EXPECT().SearchArticles(
					gomock.Any(),
					SearchArticlesQuery{
						Offset:    0,
						Limit:     3,
						SortField: SearchArticlesQuerySortFieldDate,
						SortOrder: SortOrderAsc,
						Tags:      []string{},
					},
					gomock.Any(),
					gomock.Any(),
				).SetArg(2, primaryKeys).SetArg(3, hasNext).Return(nil).Times(1)
			},
		},
		{
			desc: `Successful`,
			inputQuery: SearchArticlesQuery{
				Offset:    0,
				Limit:     3,
				Tags:      []string{},
				SortField: SearchArticlesQuerySortFieldDate,
				SortOrder: SortOrderAsc,
			},
			expectedArticles: []entity.Article{
				{
					ID: "id1",
				},
				{
					ID: "id2",
				},
				{
					ID: "id3",
				},
			},
			expectedHasNext: true,
			setUp: func(repositoryArticle *MockRepositoryArticle) {
				primaryKeys := []entity.ArticlePrimaryKey{
					{
						ArticleID: entity.ArticleID("id1"),
						Version:   1,
					},
					{
						ArticleID: entity.ArticleID("id2"),
						Version:   2,
					},
					{
						ArticleID: entity.ArticleID("id3"),
						Version:   3,
					},
				}
				articles := []entity.Article{
					{
						ID: "id1",
					},
					{
						ID: "id2",
					},
					{
						ID: "id3",
					},
				}
				hasNext := true
				repositoryArticle.EXPECT().SearchArticles(
					gomock.Any(),
					SearchArticlesQuery{
						Offset:    0,
						Limit:     3,
						SortField: SearchArticlesQuerySortFieldDate,
						SortOrder: SortOrderAsc,
						Tags:      []string{},
					},
					gomock.Any(),
					gomock.Any(),
				).SetArg(2, primaryKeys).SetArg(3, hasNext).Return(nil).Times(1)
				repositoryArticle.EXPECT().GetArticlesByPrimaryKey(
					gomock.Any(),
					primaryKeys,
					SearchArticlesQuerySortFieldDate,
					SortOrderAsc,
					gomock.Any(),
				).SetArg(4, articles).Return(nil).Times(1)
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

			articles := []entity.Article{}
			hasNext := false
			err := usecase.SearchArticles(
				ctx,
				tc.inputQuery,
				&articles,
				&hasNext,
			)
			test_helper.AssertErrorAs(t, tc.expectedError, err)

			if err == nil {
				assert.Equal(t, tc.expectedArticles, articles)
				assert.Equal(t, tc.expectedHasNext, hasNext)
			}
		})
	}
}
