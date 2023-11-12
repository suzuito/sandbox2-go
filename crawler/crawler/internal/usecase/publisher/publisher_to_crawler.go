package publisher

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

type PublisherToCrawler struct {
	queue     queue.Queue
	crawlerID crawler.CrawlerID
}

func (t *PublisherToCrawler) Do(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	for _, d := range data {
		article := d.(*note.TimeSeriesDataNoteArticle)
		if err := t.queue.PublishCrawlEvent(ctx, t.crawlerID, crawler.CrawlerInputData{
			"URL": article.URL,
		}); err != nil {
			return terrors.Wrap(err)
		}
	}
	return nil
}

func NewPublisherToCrawler(
	queue queue.Queue,
	crawlerID crawler.CrawlerID,
) crawler.Publisher {
	return &PublisherToCrawler{
		queue:     queue,
		crawlerID: crawlerID,
	}
}
