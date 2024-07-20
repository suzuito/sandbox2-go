package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) SuperPostInit(ctx *gin.Context) {
	dto, err := t.U.SuperPostInit(ctx)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
			Message: err.Error(),
		})
		return
	}
	t.P.JSON(ctx, http.StatusCreated, struct {
		PhotoStudio     *entity.PhotoStudio       `json:"photoStudio"`
		SuperMember     *entity.PhotoStudioMember `json:"superMember"`
		InitialPassword string                    `json:"initialPassword"`
	}{
		PhotoStudio:     dto.PhotoStudio,
		SuperMember:     dto.SuperMember,
		InitialPassword: dto.InitialPassword,
	})
}
