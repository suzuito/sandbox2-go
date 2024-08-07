package web

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

type WebResponseOption struct {
	SuccessStatusCode   int
	ErrorHandlers       []*WebResponseOptionErrorHandler
	DefaultErrorHandler *WebResponseOptionErrorHandler
}

type WebResponseOptionErrorHandler struct {
	AsErrorTarget  any
	IsErrorTarget  error
	FailStatusCode int
	LogError       bool
	Response       *ResponseError
}

var ptrNoEntryError *common_repository.NoEntryError
var ptrDuplicateEntryError *common_repository.DuplicateEntryError

var DefaultWebResponseOption = WebResponseOption{
	SuccessStatusCode: http.StatusOK,
	ErrorHandlers: []*WebResponseOptionErrorHandler{
		&Handler_NoEntryError_404,
		&Handler_NoEntryError_400,
	},
	DefaultErrorHandler: &Handler_Error_500,
}

var (
	Handler_NoEntryError_404 = WebResponseOptionErrorHandler{
		AsErrorTarget:  &ptrNoEntryError,
		FailStatusCode: http.StatusNotFound,
		LogError:       false,
		Response: &ResponseError{
			Message: "not found",
		},
	}
	Handler_NoEntryError_400 = WebResponseOptionErrorHandler{
		AsErrorTarget:  &ptrDuplicateEntryError,
		FailStatusCode: http.StatusBadRequest,
		LogError:       false,
		Response: &ResponseError{
			Message: "bad request",
		},
	}
	Handler_Error_500 = WebResponseOptionErrorHandler{
		FailStatusCode: http.StatusInternalServerError,
		LogError:       true,
		Response: &ResponseError{
			Message: "internal server error",
		},
	}
)

func Response(
	ctx *gin.Context,
	l *slog.Logger,
	p presenter.Presenter,
	res any,
	err error,
	opt *WebResponseOption,
) error {
	if err == nil {
		p.JSON(ctx, opt.SuccessStatusCode, res)
		return nil
	}
	for _, handler := range opt.ErrorHandlers {
		handleError := false
		if handler.AsErrorTarget != nil && errors.As(err, handler.AsErrorTarget) {
			handleError = true
		}
		if handler.IsErrorTarget != nil && errors.Is(err, handler.IsErrorTarget) {
			handleError = true
		}
		if !handleError {
			continue
		}
		if handler.LogError {
			l.Error(err.Error(), "err", err)
		}
		p.JSON(ctx, handler.FailStatusCode, handler.Response)
		return nil
	}
	if opt.DefaultErrorHandler != nil {
		if opt.DefaultErrorHandler.LogError {
			l.Error("", "err", err)
		}
		p.JSON(ctx, opt.DefaultErrorHandler.FailStatusCode, opt.DefaultErrorHandler.Response)
		return nil
	}
	return err
}
