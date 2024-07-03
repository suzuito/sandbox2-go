package web

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

type Impl struct {
	L                 *slog.Logger
	P                 presenter.Presenter
	CorsAllowOrigins  []string
	CorsAllowMethods  []string
	CorsAllowHeaders  []string
	CorsExposeHeaders []string
}
