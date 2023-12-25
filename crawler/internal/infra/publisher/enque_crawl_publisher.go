package publisher

import (
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/publisher/internal/publisherimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewEnqueCrawlPublisher(def *crawler.PublisherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Publisher, error) {
	publisher := publisherimpl.EnqueCrawlPublisher{
		TriggerCrawlerQueue: setting.PublisherFactorySetting.TriggerCrawlerQueue,
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
