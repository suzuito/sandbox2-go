package knowledgeworks

import (
	"context"
	"io"
	"net/http"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

const CrawlerID crawler.CrawlerID = "knowledgeworks"

type Crawler struct {
	fetcher fetcher.FetcherHTTP
	queue   queue.Queue
	fp      *gofeed.Parser
}

func NewCrawler(
	queue queue.Queue,
	fetcher fetcher.FetcherHTTP,
) crawler.Crawler {
	return &Crawler{
		fetcher: fetcher,
		queue:   queue,
		fp:      gofeed.NewParser(),
	}
}

func (t *Crawler) ID() crawler.CrawlerID {
	return CrawlerID
}

func (t *Crawler) Name() string {
	return string(CrawlerID)
}

func (t *Crawler) Fetch(ctx context.Context, w io.Writer, _ crawler.CrawlerInputData) error {
	request, _ := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://note.com/knowledgework/rss", nil)
	return terrors.Wrap(t.fetcher.DoRequest(ctx, request, w))
}

func (t *Crawler) Parse(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	feed, err := t.fp.Parse(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	for _, item := range feed.Items {
		returned = append(returned, &note.TimeSeriesDataNoteArticle{
			URL: item.Link,
		})
	}
	return returned, nil
}

func (t *Crawler) Publish(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	for _, d := range data {
		article := d.(*note.TimeSeriesDataNoteArticle)
		if err := t.queue.PublishCrawlEvent(ctx, "knowledgeworks", crawler.CrawlerInputData{
			"URL": article.URL,
		}); err != nil {
			return terrors.Wrap(err)
		}
	}
	return nil
}
