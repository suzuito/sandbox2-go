package queue

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type TriggerCrawlerQueue interface {
	PublishCrawlEvent(
		ctx context.Context,
		crawlerID crawler.CrawlerID,
		crawlerInputData crawler.CrawlerInputData,
	) error
	RecieveCrawlEvent(
		ctx context.Context,
		rawBytes []byte,
	) (crawler.CrawlerID, crawler.CrawlerInputData, error)
}
