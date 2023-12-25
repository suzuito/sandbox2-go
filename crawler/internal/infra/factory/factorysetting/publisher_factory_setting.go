package factorysetting

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
)

type PublisherFactorySetting struct {
	TriggerCrawlerQueue      queue.TriggerCrawlerQueue
	TimeSeriesDataRepository repository.TimeSeriesDataRepository
}
