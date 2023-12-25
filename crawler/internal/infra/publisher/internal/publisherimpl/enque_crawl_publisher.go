package publisherimpl

import (
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type EnqueCrawlPublisher struct {
	TriggerCrawlerQueue queue.TriggerCrawlerQueue
	CrawlerID           crawler.CrawlerID
}

func (t *EnqueCrawlPublisher) ID() crawler.PublisherID {
	return "enqueuecrawl"
}

func (t *EnqueCrawlPublisher) Do(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	for _, d := range data {
		jsonBytes, err := json.Marshal(d)
		if err != nil {
			return terrors.Wrap(err)
		}
		input := crawler.CrawlerInputData{}
		if err := json.Unmarshal(jsonBytes, &input); err != nil {
			return terrors.Wrap(err)
		}
		if err := t.TriggerCrawlerQueue.PublishDispatchCrawlEvent(ctx, t.CrawlerID, input); err != nil {
			return terrors.Wrap(err)
		}
	}
	return nil
}
