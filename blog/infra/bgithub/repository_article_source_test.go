package bgithub

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/v50/github"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

func TestGetBranches(t *testing.T) {
	repo := RepositoryArticleSource{
		Client:  github.NewClient(http.DefaultClient),
		Owner:   "dummy_owner",
		Repo:    "dummy_repo",
		BaseDir: "dummy_dir",
	}
	testCases := []struct {
		desc          string
		setUp         func()
		expected      []string
		expectedError string
	}{
		{
			desc: "Success",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/branches",
					httpmock.NewStringResponder(http.StatusOK, `[
					{ "name": "br1" },
					{ "name": "br2" }
				  ]
				`),
				)
			},
			expected: []string{"br1", "br2"},
		},
		{
			desc: "Failed",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/branches",
					httpmock.NewStringResponder(http.StatusInternalServerError, `
					{ "message": "dummy_error" }
				`),
				)
			},
			expectedError: "500 dummy_error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			tC.setUp()
			branches, err := repo.GetBranches(ctx)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				assert.Equal(t, tC.expected, branches)
			}
		})
	}
}

func TestSearchArticleSources(t *testing.T) {
	ctx := context.Background()
	repo := RepositoryArticleSource{
		Client:  github.NewClient(http.DefaultClient),
		Owner:   "dummy_owner",
		Repo:    "dummy_repo",
		BaseDir: "dummy_dir",
	}

	testCases := []struct {
		desc                  string
		inputQueryString      string
		setUp                 func()
		expectedError         string
		expectedArticleSource entity.ArticleSource
	}{
		{
			desc:             "Success",
			inputQueryString: "testquery",
			setUp: func() {
				searchResult := `
				{
				  "total_count": 1,
				  "incomplete_results": false,
				  "items": [
					{
					  "name": "test.md",
					  "sha": "testsha"
					}
				  ]
				}
				`
				httpmock.RegisterResponder("GET", "https://api.github.com/search/code?q=repo%3Adummy_owner%2Fdummy_repo+path%3Adummy_dir+testquery",
					httpmock.NewStringResponder(200, searchResult))
			},
			expectedArticleSource: entity.ArticleSource{
				ID:      "test",
				Version: "testsha",
			},
		},
		{
			desc:             "Failed",
			inputQueryString: "testquery",
			setUp: func() {
				searchResult := `dummy error`
				httpmock.RegisterResponder("GET", "https://api.github.com/search/code?q=repo%3Adummy_owner%2Fdummy_repo+path%3Adummy_dir+testquery",
					httpmock.NewStringResponder(500, searchResult))
			},
			expectedError: "Cannot search code in GitHub",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			tC.setUp()
			err := repo.SearchArticleSources(ctx, tC.inputQueryString, func(articleSource *entity.ArticleSource) error {
				if articleSource.ID != tC.expectedArticleSource.ID || articleSource.Version != tC.expectedArticleSource.Version {
					return fmt.Errorf("unexpected article source: want %+v, got %+v", tC.expectedArticleSource, *articleSource)
				}
				return nil
			})
			test_helper.AssertError(t, tC.expectedError, err)
		})
	}
}

func TestGetVersions(t *testing.T) {
	repo := RepositoryArticleSource{
		Client:  github.NewClient(http.DefaultClient),
		Owner:   "dummy_owner",
		Repo:    "dummy_repo",
		BaseDir: "dummy_dir",
	}

	testCases := []struct {
		desc                 string
		setUp                func()
		inputBranch          string
		inputArticleSourceID entity.ArticleSourceID
		expected             []entity.ArticleSource
		expectedError        string
	}{
		{
			desc:                 "Success",
			inputBranch:          "br1",
			inputArticleSourceID: "src1",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/commits?path=dummy_dir%2Fsrc1.md&per_page=100&sha=br1",
					httpmock.NewStringResponder(http.StatusOK, `[
						{ "sha": "commit1", "commit": { "message": "message1" } },
						{ "sha": "commit2", "commit": { "message": "message2" } }
					]`),
				)
			},
			expected: []entity.ArticleSource{
				{ID: "src1", Version: "commit1", Meta: entity.ArticleSourceMeta{URL: "https://github.com/dummy_owner/dummy_repo/blob/commit1/dummy_dir/src1.md"}},
				{ID: "src1", Version: "commit2", Meta: entity.ArticleSourceMeta{URL: "https://github.com/dummy_owner/dummy_repo/blob/commit2/dummy_dir/src1.md"}},
			},
		},
		{
			desc:                 "Failed",
			inputBranch:          "br1",
			inputArticleSourceID: "src1",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/commits?path=dummy_dir%2Fsrc1.md&per_page=100&sha=br1",
					httpmock.NewStringResponder(http.StatusInternalServerError, `
						{ "message": "dummy_error" }
					`),
				)
			},
			expectedError: "500 dummy_error",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			tC.setUp()
			versions, err := repo.GetVersions(ctx, tC.inputBranch, tC.inputArticleSourceID)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				assert.Equal(t, tC.expected, versions)
			}
		})
	}
}

func TestGetArticleSource(t *testing.T) {
	repo := RepositoryArticleSource{
		Client:  github.NewClient(http.DefaultClient),
		Owner:   "dummy_owner",
		Repo:    "dummy_repo",
		BaseDir: "dummy_dir",
	}

	testCases := []struct {
		desc                  string
		setUp                 func()
		inputArticleSourceID  entity.ArticleSourceID
		inputVersion          string
		expectedArticleSource entity.ArticleSource
		expectedContent       string
		expectedError         string
	}{
		{
			desc:                 "Success",
			inputArticleSourceID: "src1",
			inputVersion:         "dummy_sha",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `{
						"sha": "dummy_sha"
					}`),
				)
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir",
					httpmock.NewStringResponder(http.StatusOK, `
					    [
							{
								"name": "src1.md",
								"download_url": "https://www.example.com/src1.md"
							}
						]
					`),
				)
				httpmock.RegisterResponder("GET", "https://www.example.com/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `dummy_content`),
				)
			},
			expectedArticleSource: entity.ArticleSource{
				ID:      "src1",
				Version: "dummy_sha",
				Meta: entity.ArticleSourceMeta{
					URL: "https://github.com/dummy_owner/dummy_repo/blob/dummy_sha/dummy_dir/src1.md",
				},
			},
			expectedContent: "dummy_content",
		},
		{
			desc:                 "Failed - HTTP error 404 from GitHub API on 1st GetContent request",
			inputArticleSourceID: "src2",
			inputVersion:         "dummy_sha",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src2.md",
					httpmock.NewStringResponder(http.StatusNotFound, `{}`),
				)
			},
			expectedError: "Not found dummy_repo/dummy_dir/src2.md",
		},
		{
			desc:                 "Failed - HTTP error 5XX from GitHub API on 1st GetContent request",
			inputArticleSourceID: "src2",
			inputVersion:         "dummy_sha",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src2.md",
					httpmock.NewStringResponder(http.StatusInternalServerError, `{}`),
				)
			},
			expectedError: "500",
		},
		{
			desc:                 "Failed - Not HTTP error on 1st GetContent request",
			inputArticleSourceID: "src2",
			inputVersion:         "dummy_sha",
			setUp:                func() {},
			expectedError:        "no responder found",
		},
		{
			desc:                 "Failed - HTTP error 404 from GitHub API on GetDownloads request",
			inputArticleSourceID: "src1",
			inputVersion:         "dummy_sha",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `{
						"sha": "dummy_sha"
					}`),
				)
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir",
					httpmock.NewStringResponder(http.StatusNotFound, ``),
				)
			},
			expectedError: "Not found dummy_repo/dummy_dir/src1.md",
		},
		{
			desc:                 "Failed - HTTP error 5XX from GitHub API on GetDownloads request",
			inputArticleSourceID: "src1",
			inputVersion:         "dummy_sha",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `{
						"sha": "dummy_sha"
					}`),
				)
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir",
					httpmock.NewStringResponder(http.StatusInternalServerError, ``),
				)
			},
			expectedError: "500",
		},
		{
			desc:                 "Failed - Not HTTP error on GetDownloads request",
			inputArticleSourceID: "src1",
			inputVersion:         "dummy_sha",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `{
						"sha": "dummy_sha"
					}`),
				)
			},
			expectedError: "no responder found",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			tC.setUp()
			articleSrouce, content, err := repo.GetArticleSource(ctx, tC.inputArticleSourceID, tC.inputVersion)
			test_helper.AssertError(t, tC.expectedError, err)
			if tC.expectedError == "" {
				assert.Equal(t, tC.expectedArticleSource, *articleSrouce)
				assert.Equal(t, tC.expectedContent, string(content))
			}
		})
	}
}

func TestGetArticleSources(t *testing.T) {
	repo := RepositoryArticleSource{
		Client:  github.NewClient(http.DefaultClient),
		Owner:   "dummy_owner",
		Repo:    "dummy_repo",
		BaseDir: "dummy_dir",
	}

	testCases := []struct {
		desc                  string
		setUp                 func()
		inputRef              string
		expectedArticleSource entity.ArticleSource
		expectedContent       string
		expectedError         string
	}{
		{
			desc:     "Success",
			inputRef: "ref1",
			setUp: func() {
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `{
						"sha": "dummy_sha"
					}`),
				)
				httpmock.RegisterResponder("GET", "https://api.github.com/repos/dummy_owner/dummy_repo/contents/dummy_dir",
					httpmock.NewStringResponder(http.StatusOK, `
					    [
							{
								"name": "src1.md",
								"download_url": "https://www.example.com/src1.md"
							}
						]
					`),
				)
				httpmock.RegisterResponder("GET", "https://www.example.com/src1.md",
					httpmock.NewStringResponder(http.StatusOK, `dummy_content`),
				)
			},
			expectedArticleSource: entity.ArticleSource{
				ID:      "src1",
				Version: "dummy_sha",
				Meta: entity.ArticleSourceMeta{
					URL: "https://github.com/dummy_owner/dummy_repo/blob/ref1/dummy_dir/src1.md",
				},
			},
			expectedContent: "dummy_content",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			tC.setUp()
			err := repo.GetArticleSources(ctx, tC.inputRef, func(articleSrouce *entity.ArticleSource, content []byte) error {
				if tC.expectedError == "" {
					assert.Equal(t, tC.expectedArticleSource, *articleSrouce)
					assert.Equal(t, tC.expectedContent, string(content))
				}
				return nil
			})
			test_helper.AssertError(t, tC.expectedError, err)

		})
	}
}
