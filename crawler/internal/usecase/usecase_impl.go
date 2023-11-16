package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
)

type UsecaseImpl struct {
	L                        *slog.Logger
	TriggerCrawlerQueue      queue.TriggerCrawlerQueue
	CrawlerRepository        repository.CrawlerRepository
	CrawlerFactory           factory.CrawlerFactory
	TimeSeriesDataRepository repository.TimeSeriesDataRepository
}
