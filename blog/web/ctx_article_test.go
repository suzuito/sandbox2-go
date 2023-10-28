package web

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/blog/entity"
)

func TestCtxSetAndGetArticle(t *testing.T) {
	article := &entity.Article{
		ID:          "articleID",
		Version:     1,
		Title:       "Test Article",
		Description: "This is a test article",
		Tags:        []entity.Tag{{ID: "tag1"}, {ID: "tag2"}},
	}

	w := httptest.NewRecorder()

	// Create a new Gin context
	ctx, _ := gin.CreateTestContext(w)

	// Set the article in the context
	ctxSetArticle(ctx, article)

	// Retrieve the article from the context
	result := ctxGetArticle(ctx)

	// Assert that the retrieved article matches the original article
	assert.Equal(t, article, result)

	// Assert that a default article is returned when the key is not found in the context
	ctx2, _ := gin.CreateTestContext(w)
	result = ctxGetArticle(ctx2)
	assert.Equal(t, &entity.Article{}, result)

	// Assert that a default article is returned when the value in the context is of the wrong type
	ctx2.Set(ctxKeyArticle, "invalid")
	result = ctxGetArticle(ctx2)
	assert.Equal(t, &entity.Article{}, result)
}
