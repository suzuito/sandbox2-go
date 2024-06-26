package web

import (
	"context"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

var ctxKeyAuth0ValidatedClaims = "auth0ValidatedClaims"

func ctxSetAuth0ValidatedClaims(ctx *gin.Context, v *validator.ValidatedClaims) {
	ctx.Set(ctxKeyAuth0ValidatedClaims, v)
}

func ctxGetAuth0ValidatedClaims(ctx *gin.Context) *validator.ValidatedClaims {
	v, ok := ctx.Get(ctxKeyAuth0ValidatedClaims)
	if !ok {
		return nil
	}
	vv, ok := v.(*validator.ValidatedClaims)
	if !ok {
		return nil
	}
	return vv
}

type Auth0CustomClaims struct {
	Scope string `json:"scope"`
}

func (t *Auth0CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (t *Impl) MiddlewareAuth0Authe(ctx *gin.Context) {
	accessToken, err := jwtmiddleware.AuthHeaderTokenExtractor(ctx.Request)
	if err != nil {
		t.L.Debug("AuthHeaderTokenExtractor is failed", "err", err)
		t.P.JSON(ctx, http.StatusBadRequest, ResponseError{
			Message: "cannot extract access token in authorization header",
		})
		ctx.Abort()
		return
	}
	if accessToken == "" {
		t.P.JSON(ctx, http.StatusUnauthorized, ResponseError{
			Message: "no access token",
		})
		ctx.Abort()
		return
	}
	v, err := t.Auth0Validator.ValidateToken(ctx, accessToken)
	if err != nil {
		t.L.Debug("ValidateToken is failed", "err", err)
		t.P.JSON(ctx, http.StatusUnauthorized, ResponseError{
			Message: "unauthorized",
		})
		ctx.Abort()
		return
	}
	validatedClaims, ok := v.(*validator.ValidatedClaims)
	if !ok {
		t.L.Error("ValidateToken returns value that is not *validator.ValidatedClaims")
		t.P.JSON(ctx, http.StatusUnauthorized, ResponseError{
			Message: "unauthorized",
		})
		ctx.Abort()
		return
	}
	ctxSetAuth0ValidatedClaims(ctx, validatedClaims)
	fmt.Printf("%+v\n", validatedClaims.CustomClaims)
	ctx.Next()
}
