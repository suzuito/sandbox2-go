package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
)

func (t *ControllerImpl) GetAdminImport(ctx *gin.Context) {
	articleSourceID := entity.ArticleSourceID(ctx.Param("articleSourceID"))
	articleSourceVersion := ctx.Param("articleSourceVersion")
	html := ""
	var article *entity.Article
	article, html, err := t.UC.GenerateArticleHTML(
		ctx,
		articleSourceID,
		articleSourceVersion,
	)
	if err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_import.html",
		Obj: gin.H{
			"ArticleSourceID": articleSourceID,
			"Version":         articleSourceVersion,
			"Article":         article,
			"ArticleHTML":     template.HTML(html),
		},
	})
}

func (t *ControllerImpl) PostAdminImport(ctx *gin.Context) {
	articleSourceID := entity.ArticleSourceID(ctx.Param("articleSourceID"))
	articleSourceVersion := ctx.Param("articleSourceVersion")
	article, err := t.UC.UploadArticle(ctx, entity.ArticleSourceID(articleSourceID), articleSourceVersion)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		t.Presenters.Response(
			ctx,
			PresenterArgRedirect{
				Code: http.StatusFound,
				Location: fmt.Sprintf(
					"/admin/import/%s/%s/error",
					articleSourceID,
					articleSourceVersion,
				),
			},
		)
		return
	}
	t.Presenters.Response(
		ctx,
		PresenterArgRedirect{
			Code: http.StatusFound,
			Location: fmt.Sprintf(
				"/admin/import/%s/%s/success?articleID=%s&articleVersion=%d",
				articleSourceID,
				articleSourceVersion,
				article.ID,
				article.Version,
			),
		},
	)
}

func (t *ControllerImpl) GetAdminImportSuccess(ctx *gin.Context) {
	articleSourceID := entity.ArticleSourceID(ctx.Param("articleSourceID"))
	articleSourceVersion := ctx.Param("articleSourceVersion")
	articleID := entity.ArticleID(ctx.Query("articleID"))
	articleVersion := ctx.Query("articleVersion")
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_result_success.html",
		Obj: gin.H{
			"ArticleID":            articleID,
			"ArticleVersion":       articleVersion,
			"ArticleSourceID":      articleSourceID,
			"ArticleSourceVersion": articleSourceVersion,
		},
	})
}

func (t *ControllerImpl) GetAdminImportError(ctx *gin.Context) {
	articleSourceID := entity.ArticleSourceID(ctx.Param("articleSourceID"))
	articleSourceVersion := ctx.Param("articleSourceVersion")
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "admin_result_error.html",
		Obj: gin.H{
			"ArticleSourceID":      articleSourceID,
			"ArticleSourceVersion": articleSourceVersion,
		},
	})
}
