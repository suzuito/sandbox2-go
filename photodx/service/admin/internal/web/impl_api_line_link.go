package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) APIGetLINELink(ctx *gin.Context) {
	principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
	dto, err := t.U.APIGetLINELink(ctx, principal)
	if err != nil {
		var errNoEntry *common_repository.NoEntryError
		if errors.As(err, &errNoEntry) {
			t.P.JSON(
				ctx,
				http.StatusNotFound,
				common_web.ResponseError{
					Message: "not found",
				},
			)
			return
		}
		t.P.JSON(
			ctx,
			http.StatusInternalServerError,
			common_web.ResponseError{
				Message: "",
			},
		)
		return
	}
	t.P.JSON(ctx, http.StatusOK, dto)
}

func (t *Impl) APIPostLINELink(ctx *gin.Context) {
	principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
	dto, err := t.U.APIPostLINELink(ctx, principal)
	if err != nil {
		var errDupEntry *common_repository.DuplicateEntryError
		if errors.As(err, &errDupEntry) {
			t.P.JSON(
				ctx,
				http.StatusBadRequest,
				common_web.ResponseError{
					Message: "not found",
				},
			)
			return
		}
		t.P.JSON(
			ctx,
			http.StatusInternalServerError,
			common_web.ResponseError{
				Message: "",
			},
		)
		return
	}
	t.P.JSON(ctx, http.StatusOK, dto)
}

func (t *Impl) APIDeleteLINELink(ctx *gin.Context) {
	principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
	if err := t.U.APIDeleteLINELink(ctx, principal); err != nil {
		t.P.JSON(
			ctx,
			http.StatusInternalServerError,
			common_web.ResponseError{
				Message: "",
			},
		)
		return
	}
	t.P.JSON(ctx, http.StatusOK, struct{}{})
}

func (t *Impl) APIPutLINELinkMessagingAPIChannelSecret(ctx *gin.Context) {
	message := struct {
		Secret string `json:"secret"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		t.P.JSON(
			ctx,
			http.StatusBadRequest,
			common_web.ResponseError{
				Message: err.Error(),
			},
		)
		return
	}
	principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
	dto, err := t.U.APIPutLINELinkMessagingAPIChannelSecret(ctx, principal, message.Secret)
	if err != nil {
		var errNoEntry *common_repository.NoEntryError
		if errors.As(err, &errNoEntry) {
			t.P.JSON(
				ctx,
				http.StatusNotFound,
				common_web.ResponseError{
					Message: "not found",
				},
			)
			return
		}
		t.P.JSON(
			ctx,
			http.StatusInternalServerError,
			common_web.ResponseError{
				Message: "",
			},
		)
		return
	}
	t.P.JSON(ctx, http.StatusOK, dto)
}
