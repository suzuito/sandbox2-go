package bgcp

import (
	"bytes"
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/blog/entity"
	"google.golang.org/api/iterator"
)

var testbucket = "suzuito-minilla-ut"

func TestSetArticle(t *testing.T) {
	ctx := context.Background()
	cli, err := storage.NewClient(ctx)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	defer func() {
		it := cli.Bucket(testbucket).Objects(ctx, &storage.Query{
			Prefix: "RepositoryArticleHTML.TestSetArticle",
		})
		for {
			attr, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				t.Errorf("%+v\n", err)
				break
			}
			assert.Nil(t, cli.Bucket(testbucket).Object(attr.Name).Delete(ctx))
		}
	}()
	repo := RepositoryArticleHTML{
		Cli:      cli,
		Bucket:   testbucket,
		RootPath: "RepositoryArticleHTML.TestSetArticle",
	}
	inputArticle := entity.Article{
		ID:      "a001",
		Version: 123,
	}
	inputHTML := "hoge"

	// Set
	err = repo.SetArticle(ctx, &inputArticle, inputHTML)
	assert.Nil(t, err)
	_, err = cli.
		Bucket(testbucket).
		Object("RepositoryArticleHTML.TestSetArticle/a001/123.html").
		Attrs(ctx)
	assert.Nil(t, err)

	// Get
	exactHTML := bytes.NewBufferString("")
	err = repo.GetArticle(ctx, "a001", 123, exactHTML)
	assert.Nil(t, err)
}
