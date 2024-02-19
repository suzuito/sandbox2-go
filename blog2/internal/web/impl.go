package web

import (
	"github.com/suzuito/sandbox2-go/blog2/internal/environment"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/internal/presenter"
)

type Impl struct {
	U usecase.Usecase
	P presenter.Presenter
}

func NewImpl(u usecase.Usecase, env *environment.Environment) *Impl {
	return &Impl{
		U: u,
		P: &presenter.Impl{},
	}
}
