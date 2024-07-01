package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service"
)

func (t *Impl) MiddlewareAccessTokenAutho(
	requiredPermissions []*rbac.Permission,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := t.U.MiddlewareAccessTokenAutho(
			ctx,
			ctxGet[entity.Principal](ctx, ctxPrincipal),
			requiredPermissions,
		)
		var forbiddenError *service.ForbiddenError
		if errors.As(err, &forbiddenError) {
			t.P.JSON(
				ctx,
				http.StatusForbidden,
				ResponseError{
					Message: forbiddenError.Error(),
				},
			)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
