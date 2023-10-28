package usecase

import (
	"testing"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifierfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

type goMocks struct {
	Controller      *gomock.Controller
	Repository      *repository.MockRepository
	NotifierFactory *notifierfactory.MockNotifierFactory
	L               *clog.MockLogger
}

func (t *goMocks) Finish() {
	t.Controller.Finish()
}

func (t *goMocks) NewUsecase() *UsecaseImpl {
	return &UsecaseImpl{
		Repository:      t.Repository,
		NotifierFactory: t.NotifierFactory,
		L:               t.L,
	}
}

func newMocks(
	t *testing.T,
) *goMocks {
	ctrl := gomock.NewController(t)
	mocks := goMocks{
		Controller:      ctrl,
		Repository:      repository.NewMockRepository(ctrl),
		NotifierFactory: notifierfactory.NewMockNotifierFactory(ctrl),
		L:               clog.NewMockLogger(ctrl),
	}
	return &mocks
}
