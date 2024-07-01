package web

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/internal/web/presenter"
)

type Impl struct {
	U                 usecase.Usecase
	L                 *slog.Logger
	P                 presenter.Presenter
	CorsAllowOrigins  []string
	CorsAllowMethods  []string
	CorsAllowHeaders  []string
	CorsExposeHeaders []string
}
