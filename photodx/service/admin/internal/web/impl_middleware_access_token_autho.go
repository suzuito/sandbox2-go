package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) MiddlewareAccessTokenAutho(policyString string) gin.HandlerFunc {
	policy := rbac.NewPolicy(policyString)
	return func(ctx *gin.Context) {
		principal := common_web.CtxGet[entity.AdminPrincipal](ctx, common_web.CtxPrincipal)
		if principal == nil {
			t.P.JSON(ctx, http.StatusForbidden, common_web.ResponseError{
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		result, err := policy.EvalGinContext(
			principal.GetPermissions(),
			string(principal.GetPhotoStudioMemberID()),
			string(principal.GetPhotoStudioID()),
			ctx,
		)
		if err != nil {
			t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
				Message: "internal server error",
			})
			ctx.Abort()
			return
		}
		if !result {
			t.P.JSON(ctx, http.StatusForbidden, common_web.ResponseError{
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
