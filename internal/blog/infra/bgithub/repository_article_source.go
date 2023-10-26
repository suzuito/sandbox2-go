package bgithub

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

func getFilePathForArticleSource(
	baseDir string,
	articleSourceID entity.ArticleSourceID,
) string {
	return fmt.Sprintf("%s/%s.md", baseDir, articleSourceID)
}

type RepositoryArticleSource struct {
	Client  *github.Client
	Owner   string
	Repo    string
	BaseDir string
}

func (t *RepositoryArticleSource) newArticleSourceMetaURL(
	articleSourceID entity.ArticleSourceID,
	version string,
) string {
	return fmt.Sprintf(
		"https://github.com/%s/%s/blob/%s/%s/%s.md",
		t.Owner,
		t.Repo,
		version,
		t.BaseDir,
		articleSourceID,
	)
}

func (t *RepositoryArticleSource) GetArticleSources(
	ctx context.Context,
	ref string,
	proc func(*entity.ArticleSource, []byte) error,
) error {
	_, directoryContents, _, err := t.Client.Repositories.GetContents(
		ctx,
		t.Owner,
		t.Repo,
		t.BaseDir,
		&github.RepositoryContentGetOptions{
			Ref: ref,
		},
	)
	if err != nil {
		return fmt.Errorf("Cannot search code in GitHub : %w", err)
	}
	for _, content := range directoryContents {
		name := strings.ReplaceAll(content.GetName(), path.Ext(content.GetName()), "")
		articleSourceID := entity.ArticleSourceID(name)
		articleSource, b, err := t.getArticleSource(
			ctx,
			articleSourceID,
			ref,
		)
		if err != nil {
			return err
		}
		if err := proc(articleSource, b); err != nil {
			return err
		}
	}
	return nil
}

func (t *RepositoryArticleSource) GetArticleSource(
	ctx context.Context,
	articleSourceID entity.ArticleSourceID,
	version string,
) (*entity.ArticleSource, []byte, error) {
	return t.getArticleSource(ctx, articleSourceID, version)
}

func (t *RepositoryArticleSource) getArticleSource(
	ctx context.Context,
	articleSourceID entity.ArticleSourceID,
	version string,
) (*entity.ArticleSource, []byte, error) {
	filePath := getFilePathForArticleSource(t.BaseDir, articleSourceID)
	fileContent, _, resp, err := t.Client.Repositories.GetContents(
		ctx,
		t.Owner,
		t.Repo,
		filePath,
		&github.RepositoryContentGetOptions{
			Ref: version,
		},
	)
	if err != nil {
		if resp == nil || resp.Response == nil {
			return nil, nil, err
		}
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil, &usecase.RepositoryError{
				EntityURL: fmt.Sprintf("%s/%s", t.Repo, filePath),
				Code:      usecase.RepositoryErrorCodeNotFound,
			}
		}
		return nil, nil, fmt.Errorf(": %w", err)
	}
	reader, resp, err := t.Client.Repositories.DownloadContents(
		ctx,
		t.Owner,
		t.Repo,
		filePath,
		&github.RepositoryContentGetOptions{
			Ref: version,
		},
	)
	if err != nil {
		if resp == nil || resp.Response == nil {
			return nil, nil, err
		}
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil, &usecase.RepositoryError{
				EntityURL: fmt.Sprintf("%s/%s", t.Repo, filePath),
				Code:      usecase.RepositoryErrorCodeNotFound,
			}
		}
		return nil, nil, fmt.Errorf(": %w", err)
	}
	defer reader.Close()
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}
	return &entity.ArticleSource{
		ID:      articleSourceID,
		Version: fileContent.GetSHA(),
		Meta: entity.ArticleSourceMeta{
			URL: t.newArticleSourceMetaURL(articleSourceID, version),
		},
	}, content, nil
}

func (t *RepositoryArticleSource) GetBranches(ctx context.Context) ([]string, error) {
	branches, _, err := t.Client.Repositories.ListBranches(
		ctx,
		t.Owner,
		t.Repo,
		&github.BranchListOptions{},
	)
	if err != nil {
		return nil, err
	}
	ret := []string{}
	for _, b := range branches {
		ret = append(ret, b.GetName())
	}
	return ret, nil
}

func (t *RepositoryArticleSource) GetVersions(
	ctx context.Context,
	branch string,
	articleSourceID entity.ArticleSourceID,
) ([]entity.ArticleSource, error) {
	commits, _, err := t.Client.Repositories.ListCommits(
		ctx,
		t.Owner,
		t.Repo,
		&github.CommitsListOptions{
			SHA: branch,
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
			Path: getFilePathForArticleSource(t.BaseDir, articleSourceID),
		},
	)
	if err != nil {
		return nil, err
	}
	ret := []entity.ArticleSource{}
	for _, c := range commits {
		ret = append(
			ret,
			entity.ArticleSource{
				ID:      articleSourceID,
				Version: c.GetSHA(),
				Meta: entity.ArticleSourceMeta{
					URL: t.newArticleSourceMetaURL(articleSourceID, c.GetSHA()),
				},
			},
		)
	}
	return ret, nil
}

func (t *RepositoryArticleSource) SearchArticleSources(
	ctx context.Context,
	queryString string,
	proc func(*entity.ArticleSource) error,
) error {
	results, _, err := t.Client.Search.Code(
		ctx,
		fmt.Sprintf("repo:%s/%s path:%s %s", t.Owner, t.Repo, t.BaseDir, queryString),
		&github.SearchOptions{},
	)
	if err != nil {
		return fmt.Errorf("Cannot search code in GitHub : %w", err)
	}
	for _, result := range results.CodeResults {
		name := strings.ReplaceAll(result.GetName(), path.Ext(result.GetName()), "")
		articleSource := entity.ArticleSource{
			ID:      entity.ArticleSourceID(name),
			Version: result.GetSHA(),
		}
		proc(&articleSource)
	}
	return nil
}
