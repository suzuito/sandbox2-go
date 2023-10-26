package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase/markdown2html"
	"github.com/suzuito/sandbox2-go/internal/common/test_helper"
	gomock "go.uber.org/mock/gomock"
)

func TestGenerateArticleHTMLFromMarkdown(t *testing.T) {
	testCases := []struct {
		desc                string
		inputArticleSource  *entity.ArticleSource
		inputMD             []byte
		expectedArticle     *entity.Article
		expectedArticleHTML string
		expectedError       string
		setUp               func(mockMarkdown2html *markdown2html.MockMarkdown2HTML)
	}{
		{
			desc:               "Test Case 1: Your case description here",
			inputArticleSource: &entity.ArticleSource{}, // your input here
			inputMD:            []byte("Your Markdown here"),
			expectedArticle: &entity.Article{
				ID:      "article01",
				Version: 1,
				Title:   "title01",
				ArticleSource: entity.ArticleSource{
					ID:      "src01",
					Version: "v1",
				},
			}, // your expected output here
			expectedArticleHTML: "Your expected HTML here",
			expectedError:       "", // if you expect an error, specify it here
			setUp: func(mockMarkdown2html *markdown2html.MockMarkdown2HTML) {
				mockMarkdown2html.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					SetArg(2, "Your expected HTML here").
					SetArg(3, entity.Article{
						ID:      "article01",
						Version: 1,
						Title:   "title01",
						ArticleSource: entity.ArticleSource{
							ID:      "src01",
							Version: "v1",
						},
					}).
					Return(nil).Times(1)
			},
		},
		{
			desc:                "Test Case 2: Error in Markdown to HTML conversion",
			inputArticleSource:  &entity.ArticleSource{},
			inputMD:             []byte("Your Markdown here"),
			expectedArticle:     nil,
			expectedArticleHTML: "",
			expectedError:       "Markdown to HTML conversion failed",
			setUp: func(mockMarkdown2html *markdown2html.MockMarkdown2HTML) {
				mockMarkdown2html.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(errors.New("Markdown to HTML conversion failed")).Times(1)
			},
		},
		{
			desc:                "Test Case 3: Invalid article",
			inputArticleSource:  &entity.ArticleSource{},
			inputMD:             []byte("Your Markdown here"),
			expectedArticle:     nil,
			expectedArticleHTML: "",
			expectedError:       "Invalid article : *",
			setUp: func(mockMarkdown2html *markdown2html.MockMarkdown2HTML) {
				mockMarkdown2html.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					SetArg(2, "Your expected HTML here").
					SetArg(3, entity.Article{
						ID:      "article01",
						Version: 1,
						Title:   "",
						ArticleSource: entity.ArticleSource{
							ID:      "src01",
							Version: "v1",
						},
					}).
					Return(nil).Times(1)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockMarkdown2html := markdown2html.NewMockMarkdown2HTML(ctrl)
			tC.setUp(mockMarkdown2html)
			u := UsecaseImpl{
				Markdown2HTML: mockMarkdown2html,
			}
			article, articleHTML, err := u.GenerateArticleHTMLFromMarkdown(ctx, tC.inputArticleSource, tC.inputMD)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tC.expectedArticle, article)
				assert.Equal(t, tC.expectedArticleHTML, articleHTML)
			}
		})
	}
}

func TestGenerateArticleHTML(t *testing.T) {
	testCases := []struct {
		desc                      string
		inputArticleSourceID      entity.ArticleSourceID
		inputArticleSourceVersion string
		expectedArticle           *entity.Article
		expectedArticleHTML       string
		expectedError             string
		setUp                     func(mockRepoArticleSource *MockRepositoryArticleSource, mockMarkdown2HTML *markdown2html.MockMarkdown2HTML)
	}{
		{
			desc:                      "Test Case 1: Successful generation of article HTML",
			inputArticleSourceID:      "articleSourceID",
			inputArticleSourceVersion: "v1",
			expectedArticle: &entity.Article{
				ID:      "article01",
				Version: 1,
				Title:   "title01",
				ArticleSource: entity.ArticleSource{
					ID:      "src01",
					Version: "v1",
				},
			}, // your expected output here
			expectedArticleHTML: "Your expected HTML here",
			expectedError:       "",
			setUp: func(mockRepoArticleSource *MockRepositoryArticleSource, mockMarkdown2HTML *markdown2html.MockMarkdown2HTML) {
				mockRepoArticleSource.EXPECT().GetArticleSource(
					gomock.Any(),
					gomock.Eq(entity.ArticleSourceID("articleSourceID")),
					gomock.Eq("v1"),
				).Return(&entity.ArticleSource{}, []byte("Your Markdown here"), nil).Times(1)
				mockMarkdown2HTML.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					SetArg(2, "Your expected HTML here").
					SetArg(3, entity.Article{
						ID:      "article01",
						Version: 1,
						Title:   "title01",
						ArticleSource: entity.ArticleSource{
							ID:      "src01",
							Version: "v1",
						},
					}).
					Return(nil).Times(1)
			},
		},
		{
			desc:                      "Test Case 2: Error",
			inputArticleSourceID:      "articleSourceID",
			inputArticleSourceVersion: "v1",
			expectedArticle:           nil,
			expectedArticleHTML:       "",
			expectedError:             "Dummy error",
			setUp: func(mockRepoArticleSource *MockRepositoryArticleSource, mockMarkdown2HTML *markdown2html.MockMarkdown2HTML) {
				mockRepoArticleSource.EXPECT().GetArticleSource(
					gomock.Any(),
					gomock.Eq(entity.ArticleSourceID("articleSourceID")),
					gomock.Eq("v1"),
				).Return(nil, nil, errors.New("Dummy error")).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepoArticleSource := NewMockRepositoryArticleSource(ctrl)
			mockMarkdown2HTML := markdown2html.NewMockMarkdown2HTML(ctrl)
			tc.setUp(mockRepoArticleSource, mockMarkdown2HTML)

			usecase := UsecaseImpl{
				RepositoryArticleSource: mockRepoArticleSource,
				Markdown2HTML:           mockMarkdown2HTML,
			}

			article, articleHTML, err := usecase.GenerateArticleHTML(ctx, tc.inputArticleSourceID, tc.inputArticleSourceVersion)
			test_helper.AssertError(t, tc.expectedError, err)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedArticle, article)
				assert.Equal(t, tc.expectedArticleHTML, articleHTML)
			}
		})
	}
}
