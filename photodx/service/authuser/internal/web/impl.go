package web

import (
	"log/slog"
	"net/url"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

type Impl struct {
	L                 *slog.Logger
	U                 usecase.Usecase
	P                 presenter.Presenter
	OAuth2RedirectURL url.URL
	FrontBaseURL      url.URL
}
