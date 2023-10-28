package usecase

import (
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/repository"
)

type UsecaseImpl struct {
	Repository     repository.Repository
	Queue          queue.Queue
	CrawlerFactory crawlerfactory.CrawlerFactory
	L              clog.Logger
}
