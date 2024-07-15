package web

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

func MiddlewareAdminAccessTokenAutho(
	logger *slog.Logger,
	policyString string,
	presenter presenter.Presenter,
) gin.HandlerFunc {
	policy := entity.NewPolicyAdminPrincipalAccessToken(policyString)
	return func(ctx *gin.Context) {
		principal := CtxGetAdminPrincipalAccessToken(ctx)
		if principal == nil {
			presenter.JSON(ctx, http.StatusForbidden, ResponseError{
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
			logger.Error("", "err", err)
			presenter.JSON(ctx, http.StatusInternalServerError, ResponseError{
				Message: "internal server error",
			})
			ctx.Abort()
			return
		}
		if !result {
			presenter.JSON(ctx, http.StatusForbidden, ResponseError{
				Message: "forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
