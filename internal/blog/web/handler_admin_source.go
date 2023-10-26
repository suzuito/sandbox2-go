package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func (t *ControllerImpl) GetAdminSearchSource(ctx *gin.Context) {
	articleSources := []*entity.ArticleSource{}
	if err := t.RepositoryArticleSource.SearchArticleSources(
		ctx,
		ctx.Query("q"),
		func(as *entity.ArticleSource) error {
			articleSources = append(articleSources, as)
			return nil
		},
	); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_sources_search.html",
		Obj: gin.H{
			"ArticleSources": articleSources,
		},
	})
}

func (t *ControllerImpl) GetAdminSourcesByID(ctx *gin.Context) {
	articleSourceID := entity.ArticleSourceID(ctx.Param("articleSourceID"))
	branch := ctx.DefaultQuery("branch", "main")
	articleSources, err := t.RepositoryArticleSource.GetVersions(ctx, branch, articleSourceID)
	if err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	articles := []entity.Article{}
	mapArticleSourceVersionToArticle := map[string]*entity.Article{}
	{
		if err := t.RepositoryArticle.GetArticlesByArticleSourceID(ctx, articleSourceID, &articles); err != nil {
			t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
			return
		}
		for i := range articles {
			mapArticleSourceVersionToArticle[articles[i].ArticleSource.Version] = &articles[i]
		}
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_sources_by_id.html",
		Obj: gin.H{
			"ArticleSources":                   articleSources,
			"MapArticleSourceVersionToArticle": mapArticleSourceVersionToArticle,
		},
	})
}
