package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/cweb"
)

func (t *Impl) APIPostAdminFiles(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		t.P.RenderJSON(ctx, http.StatusBadRequest, APIErrorResponse{})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		t.P.RenderJSON(ctx, http.StatusBadRequest, APIErrorResponse{})
		return
	}
	defer file.Close()
	fileName := fileHeader.Filename
	dto, err := t.U.APIPostAdminFiles(
		ctx,
		fileName,
		file,
	)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.RenderJSON(ctx, http.StatusInternalServerError, APIErrorResponse{})
		return
	}
	t.P.RenderJSON(ctx, http.StatusOK, struct {
		File          *entity.File          `json:"file"`
		FileThumbnail *entity.FileThumbnail `json:"fileThumbnail"`
	}{
		File:          dto.File,
		FileThumbnail: dto.FileThumbnail,
	})
}

func (t *Impl) APIGetAdminFiles(ctx *gin.Context) {
	page := cweb.DefaultQueryAsInt(ctx, "page", 0)
	size := cweb.DefaultQueryAsInt(ctx, "size", 10)
	query := ctx.DefaultQuery("q", "")
	dto, err := t.U.APIGetAdminFiles(
		ctx,
		query,
		page,
		size,
	)
	if err != nil {
		t.P.RenderJSON(ctx, http.StatusInternalServerError, APIErrorResponse{})
		return
	}
	t.P.RenderJSON(ctx, http.StatusOK, struct {
		Files []*entity.FileAndThumbnail `json:"files"`
		Page  int                        `json:"page"`
		Size  int                        `json:"size"`
	}{
		Files: dto.Files,
		Page:  page,
		Size:  size,
	})
}
