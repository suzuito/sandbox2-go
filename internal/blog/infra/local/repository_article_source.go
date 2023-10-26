package local

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type RepositoryArticleSource struct {
	DirPath string
}

func (t *RepositoryArticleSource) GetArticleSource(
	ctx context.Context,
	articleSourceID entity.ArticleSourceID,
) ([]byte, error) {
	return ioutil.ReadFile(path.Join(t.DirPath, fmt.Sprintf("%s.md", string(articleSourceID))))
}

func (t *RepositoryArticleSource) GetArticleSources(
	ctx context.Context,
	proc func(entity.ArticleSourceID, []byte) error,
) error {
	entries, err := os.ReadDir(t.DirPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		articleSourceID := strings.ReplaceAll(entry.Name(), path.Ext(entry.Name()), "")
		b, err := ioutil.ReadFile(path.Join(t.DirPath, fmt.Sprintf("%s.md", string(articleSourceID))))
		if err != nil {
			return err
		}
		if err := proc(entity.ArticleSourceID(articleSourceID), b); err != nil {
			return err
		}
	}
	return nil
}
