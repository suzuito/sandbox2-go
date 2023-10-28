package usecase

import (
	"testing"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

type goMocks struct {
	Controller     *gomock.Controller
	Repository     *repository.MockRepository
	Queue          *queue.MockQueue
	CrawlerFactory *crawlerfactory.MockCrawlerFactory
	L              *clog.MockLogger
}

func (t *goMocks) Finish() {
	t.Controller.Finish()
}

func (t *goMocks) NewUsecase() *UsecaseImpl {
	return &UsecaseImpl{
		Repository:     t.Repository,
		CrawlerFactory: t.CrawlerFactory,
		Queue:          t.Queue,
		L:              t.L,
	}
}

func newMocks(
	t *testing.T,
) *goMocks {
	ctrl := gomock.NewController(t)
	mocks := goMocks{
		Controller:     ctrl,
		Repository:     repository.NewMockRepository(ctrl),
		Queue:          queue.NewMockQueue(ctrl),
		CrawlerFactory: crawlerfactory.NewMockCrawlerFactory(ctrl),
		L:              clog.NewMockLogger(ctrl),
	}
	return &mocks
}
