package web

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
	"github.com/suzuito/sandbox2-go/photodx/service/user/internal/usecase"
)

type Impl struct {
	U usecase.Usecase
	L *slog.Logger
	P presenter.Presenter
}
