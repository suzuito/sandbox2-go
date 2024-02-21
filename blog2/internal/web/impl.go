package web

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/internal/presenter"
)

type Impl struct {
	U          usecase.Usecase
	P          presenter.Presenter
	L          *slog.Logger
	AdminToken string
}

func NewPresenter() presenter.Presenter {
	return &presenter.Impl{}
}
