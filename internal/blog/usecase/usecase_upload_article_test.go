package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase/markdown2html"
	gomock "go.uber.org/mock/gomock"
)

func TestUploadArticle(t *testing.T) {
	testCases := []struct {
		desc                      string
		inputArticleSourceID      entity.ArticleSourceID
		inputArticleSourceVersion string
		expectedArticle           entity.Article
		expectedError             string
		setUp                     func(
			repositoryArticleSource *MockRepositoryArticleSource,
			repositoryArticle *MockRepositoryArticle,
			repositoryArticleHTML *MockRepositoryArticleHTML,
			md2HTML *markdown2html.MockMarkdown2HTML,
		)
	}{
		{
			desc:                      `Successful. empty primary keys.`,
			inputArticleSourceID:      "s1",
			inputArticleSourceVersion: "v1",
			expectedArticle: entity.Article{
				ID:      "a1",
				Version: 1,
				Title:   "t1",
				ArticleSource: entity.ArticleSource{
					ID:      entity.ArticleSourceID("s1_returned"),
					Version: "v1_returned",
				},
			},
			setUp: func(
				repositoryArticleSource *MockRepositoryArticleSource,
				repositoryArticle *MockRepositoryArticle,
				repositoryArticleHTML *MockRepositoryArticleHTML,
				md2HTML *markdown2html.MockMarkdown2HTML,
			) {
				returnedArticleSource := entity.ArticleSource{
					ID:      entity.ArticleSourceID("s1_returned"),
					Version: "v1_returned",
				}
				returnedArticle := entity.Article{
					ID:            "a1",
					Version:       1,
					Title:         "t1",
					ArticleSource: returnedArticleSource,
				}
				returnedHTML := "dummy_html"
				repositoryArticleSource.EXPECT().GetArticleSource(
					gomock.Any(),
					entity.ArticleSourceID("s1"),
					"v1",
				).Return(
					&returnedArticleSource,
					[]byte("dummy_source_content"),
					nil,
				).Times(1)
				md2HTML.EXPECT().Generate(
					gomock.Any(),
					"dummy_source_content",
					gomock.Any(),
					gomock.Any(),
				).
					SetArg(2, returnedHTML).
					SetArg(3, returnedArticle).
					Return(nil).
					Times(1)
				repositoryArticleHTML.EXPECT().SetArticle(
					gomock.Any(),
					&returnedArticle,
					returnedHTML,
				).
					Return(nil).
					Times(1)
				repositoryArticle.EXPECT().SetArticle(
					gomock.Any(),
					&returnedArticle,
				).
					Return(nil).
					Times(1)
				repositoryArticle.EXPECT().SetArticleSearchIndex(
					gomock.Any(),
					&returnedArticle,
				).
					Return(nil).
					Times(1)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepositoryArticle := NewMockRepositoryArticle(ctrl)
			mockRepositoryArticleSource := NewMockRepositoryArticleSource(ctrl)
			mockRepositoryArticleHTML := NewMockRepositoryArticleHTML(ctrl)
			mockMd2HTML := markdown2html.NewMockMarkdown2HTML(ctrl)
			tc.setUp(
				mockRepositoryArticleSource,
				mockRepositoryArticle,
				mockRepositoryArticleHTML,
				mockMd2HTML,
			)
			usecase := UsecaseImpl{
				RepositoryArticle:       mockRepositoryArticle,
				RepositoryArticleSource: mockRepositoryArticleSource,
				RepositoryArticleHTML:   mockRepositoryArticleHTML,
				Markdown2HTML:           mockMd2HTML,
			}
			article, err := usecase.UploadArticle(
				ctx,
				tc.inputArticleSourceID,
				tc.inputArticleSourceVersion,
			)
			test_helper.AssertErrorAs(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.expectedArticle, *article)
			}
		})
	}
}
