package web

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func MiddlewareUserAccessTokenAutho(
	l *slog.Logger,
	policyString string,
	p presenter.Presenter,
) gin.HandlerFunc {
	policy := entity.NewPolicyUserPrincipalAccessToken(policyString)
	return func(ctx *gin.Context) {
		principal := CtxGetUserPrincipalAccessToken(ctx)
		if principal == nil {
			p.JSON(ctx, http.StatusForbidden, ResponseError{
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		result, err := policy.EvalGinContext(
			principal,
			ctx,
		)
		if err != nil {
			p.JSON(ctx, http.StatusInternalServerError, ResponseError{
				Message: "internal server error",
			})
			ctx.Abort()
			return
		}
		if !result {
			p.JSON(ctx, http.StatusForbidden, ResponseError{
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
