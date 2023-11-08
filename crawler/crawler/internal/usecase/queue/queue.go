package queue

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

// TODO インターフェース名を変えたい
// 実質的に、CrawlerのトリガーイベントとしてのQueueに過ぎない
// 名前の候補は TriggerQueueCrawler など
type Queue interface {
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
