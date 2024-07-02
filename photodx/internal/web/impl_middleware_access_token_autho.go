package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

func (t *Impl) MiddlewareAccessTokenAutho(policyString string) gin.HandlerFunc {
	policy := rbac.NewPolicy(policyString)
	return func(ctx *gin.Context) {
		principal := ctxGet[entity.Principal](ctx, ctxPrincipal)
		if principal == nil {
			t.P.JSON(ctx, http.StatusForbidden, ResponseError{
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
			t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{
				Message: "internal server error",
			})
			ctx.Abort()
			return
		}
		if !result {
			t.P.JSON(ctx, http.StatusForbidden, ResponseError{
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
