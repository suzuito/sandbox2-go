package usecase

import (
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
)

type UsecaseImpl struct {
	L                        clog.Logger
	TriggerCrawlerQueue      queue.TriggerCrawlerQueue
	CrawlerRepository        repository.CrawlerRepository
	CrawlerFactory           factory.CrawlerFactory
	TimeSeriesDataRepository repository.TimeSeriesDataRepository
}
