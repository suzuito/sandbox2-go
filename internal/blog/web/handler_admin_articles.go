package web

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func (t *ControllerImpl) GetAdminArticlesByID(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	html := bytes.NewBufferString("")
	{
		if err := t.RepositoryArticleHTML.GetArticle(ctx, article.ID, article.Version, html); err != nil {
			t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
			return
		}
	}
	articles := []entity.Article{}
	MapArticleSourceToArticleSourceVersions := map[entity.ArticleSourceID][]entity.ArticleSource{}
	MapArticleSourceVersionToArticleVersion := map[string]int32{}
	{
		if err := t.RepositoryArticle.GetArticlesByID(ctx, article.ID, &articles); err != nil {
			t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
			return
		}
		articlesSourceIDMap := map[entity.ArticleSourceID]string{}
		for _, article := range articles {
			articlesSourceIDMap[article.ArticleSource.ID] = ""
		}
		for articleSourceID := range articlesSourceIDMap {
			articleSourceVersions, err := t.RepositoryArticleSource.GetVersions(ctx, "main", articleSourceID)
			if err != nil {
				t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
				return
			}
			MapArticleSourceToArticleSourceVersions[articleSourceID] = articleSourceVersions
			for _, articleSourceVersion := range articleSourceVersions {
				for _, article := range articles {
					if article.ArticleSource.Version == articleSourceVersion.Version {
						MapArticleSourceVersionToArticleVersion[articleSourceVersion.Version] = article.Version
					}
				}
			}
		}
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_articles_by_id.html",
		Obj: responseGetAdminArticlesByID{
			responseCommon: responseCommon{
				Header: struct{}{},
				Meta:   siteMetaData{},
			},
			Article:                                 *ctxGetArticle(ctx),
			ArticleHTML:                             template.HTML(html.String()),
			ArticlesByID:                            articles,
			MapArticleSourceToArticleSourceVersions: MapArticleSourceToArticleSourceVersions,
			MapArticleSourceVersionToArticleVersion: MapArticleSourceVersionToArticleVersion,
		},
	})
}

type responseGetAdminArticlesByID struct {
	responseCommon
	Article                                 entity.Article
	ArticleHTML                             template.HTML
	ArticlesByID                            []entity.Article
	MapArticleSourceToArticleSourceVersions map[entity.ArticleSourceID][]entity.ArticleSource
	MapArticleSourceVersionToArticleVersion map[string]int32
}
