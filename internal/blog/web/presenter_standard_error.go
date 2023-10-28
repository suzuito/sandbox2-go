package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

type PresenterStandardError struct {
}

func (t *PresenterStandardError) Response(
	ctx *gin.Context,
	arg *PresenterArgStandardError,
) {
	defer ctx.Abort()
	if arg.ArgHTML != nil {
		ctx.HTML(arg.ArgHTML.Code, arg.ArgHTML.Name, arg.ArgHTML.Obj)
		return
	}
	err := arg.Err
	var errRepository *usecase.RepositoryError
	if errors.As(err, &errRepository) {
		if errRepository.Code == usecase.RepositoryErrorCodeNotFound {
			ctx.Set("code", http.StatusNotFound)
			ctx.HTML(
				http.StatusNotFound,
				"pc_error.html",
				responseError{
					Title:   "404 - エンティティが見つかりません",
					Message: "Not found",
				},
			)
			return
		}
	}
	defer clog.L.Errorf(ctx, "%+v", err)
	ctx.Set("code", http.StatusInternalServerError)
	ctx.Set("error", err)
	ctx.HTML(
		http.StatusInternalServerError,
		"pc_error.html",
		responseError{
			Title:   "500 - サーバーエラー",
			Message: "Internal server error",
		},
	)
}

type PresenterArgStandardError struct {
	Err     error
	ArgHTML *PresenterArgHTML
}
