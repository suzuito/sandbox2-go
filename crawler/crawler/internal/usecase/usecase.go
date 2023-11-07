package usecase

import (
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawlerfactory2"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
)

type UsecaseImpl struct {
	Repository      repository.Repository
	Queue           queue.Queue
	CrawlerFactory  crawlerfactory.CrawlerFactory
	CrawlerFactory2 crawlerfactory2.CrawlerFactory
	L               clog.Logger
}
