package local

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/suzuito/sandbox2-go/blog/entity"
)

type RepositoryArticleHTML struct {
	DirPath string
}

func (t *RepositoryArticleHTML) SetArticle(
	ctx context.Context,
	article *entity.Article,
	html string,
) error {
	return ioutil.WriteFile(fmt.Sprintf("%s/%s.html", t.DirPath, article.ID), []byte(html), 0600)
}

func (t *RepositoryArticleHTML) GetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	html io.Writer,
) error {
	file, err := os.Open(fmt.Sprintf("%s/%s.html", t.DirPath, articleID))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(html, file)
	if err != nil {
		return err
	}
	return nil
}
