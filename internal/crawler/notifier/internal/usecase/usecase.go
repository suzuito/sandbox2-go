package usecase

import (
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifierfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/repository"
)

type UsecaseImpl struct {
	NotifierFactory notifierfactory.NotifierFactory
	Repository      repository.Repository
	L               clog.Logger
}
