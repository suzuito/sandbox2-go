package web

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web/presenter"
)

type Impl struct {
	U usecase.Usecase
	L *slog.Logger
	P presenter.Presenter
}
