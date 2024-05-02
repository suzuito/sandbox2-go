package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageAdminFilesImage struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
}

func (t *Impl) GetAdminFilesImage(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_files_image.html",
		PageAdminFilesImage{},
	)
}

// func (t *Impl) PostAdminFilesImage(ctx *gin.Context) {
// 	fileHeader, err := ctx.FormFile("inputImage")
// 	if err != nil {
// 		t.P.RenderHTML(
// 			ctx,
// 			http.StatusBadRequest,
// 			"page_error.html",
// 			PageError{
// 				ComponentCommonHead: ComponentCommonHead{},
// 				Message:             err.Error(),
// 			},
// 		)
// 		return
// 	}
// 	f, err := fileHeader.Open()
// 	if err != nil {
// 		t.P.RenderHTML(
// 			ctx,
// 			http.StatusBadRequest,
// 			"page_error.html",
// 			PageError{
// 				ComponentCommonHead: ComponentCommonHead{},
// 				Message:             err.Error(),
// 			},
// 		)
// 		return
// 	}
// 	defer f.Close()
// 	_, err = t.U.PostAdminFiles(ctx, "", entity.FileTypeImage, f)
// 	if err != nil {
// 		t.L.Error("", "err", err)
// 		t.RenderUnknownError(ctx)
// 		return
// 	}
// 	ctx.Redirect(http.StatusFound, "/admin/files/image")
// }
