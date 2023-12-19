package enqueuecrawl

import (
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Publisher struct {
	TriggerCrawlerQueue queue.TriggerCrawlerQueue
	CrawlerID           crawler.CrawlerID
}

func (t *Publisher) ID() crawler.PublisherID {
	return "enqueuecrawl"
}

func (t *Publisher) Do(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
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

func New(def *crawler.PublisherDefinition, arg *factory.NewFuncPublisherArgument) (crawler.Publisher, error) {
	publisher := Publisher{
		TriggerCrawlerQueue: arg.TriggerCrawlerQueue,
	}
	if def.ID != publisher.ID() {
		return nil, factory.ErrNoMatchedPublisherID
	}
	var err error
	publisher.CrawlerID, err = argument.GetFromArgumentDefinition[crawler.CrawlerID](def.Argument, "CrawlerID")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &publisher, nil
}
