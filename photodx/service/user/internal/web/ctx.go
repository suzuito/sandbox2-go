package web

import (
	"github.com/gin-gonic/gin"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func CtxSetPhotoStudio(ctx *gin.Context, photoStudio *common_entity.PhotoStudio) {
	common_web.CtxSet(ctx, "PhotoStudio", photoStudio)
}
func CtxGetPhotoStudio(ctx *gin.Context) *common_entity.PhotoStudio {
	return common_web.CtxGet[*common_entity.PhotoStudio](ctx, "PhotoStudio")
}
