package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageAdminArticleImages struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
	JsEnv               PageAdminArticleImagesJsEnv
	Article             *entity.Article
}

type PageAdminArticleImagesJsEnv struct {
}

func (t *Impl) PageAdminArticleImages(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article_images.html",
		PageAdminArticleImages{
			ComponentCommonHead: ComponentCommonHead{},
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			JsEnv:   PageAdminArticleImagesJsEnv{},
			Article: article,
		},
	)
}

func (t *Impl) PostAdminArticleImages(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	fileHeader, err := ctx.FormFile("inputImage")
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusBadRequest,
			"page_error.html",
			PageError{
				ComponentCommonHead: ComponentCommonHead{},
				Message:             err.Error(),
			},
		)
		return
	}
	f, err := fileHeader.Open()
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusBadRequest,
			"page_error.html",
			PageError{
				ComponentCommonHead: ComponentCommonHead{},
				Message:             err.Error(),
			},
		)
		return
	}
	defer f.Close()
	dto, err := t.U.PostAdminArticleImages(ctx, article, f)
	if err != nil {
		t.L.Error("", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	fmt.Printf("%+v\n", dto)
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/admin/articles/%s/images", article.ID))
}
