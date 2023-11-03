package golangweekly

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

const CrawlerID crawler.CrawlerID = "golangweekly"

type Crawler struct {
	repository repository.Repository
	fetcher    fetcher.FetcherHTTP
	fp         *gofeed.Parser
}

func NewCrawler(
	repository repository.Repository,
	fetcher fetcher.FetcherHTTP,
) crawler.Crawler {
	return &Crawler{
		repository: repository,
		fetcher:    fetcher,
		fp:         gofeed.NewParser(),
	}
}

func (t *Crawler) ID() crawler.CrawlerID {
	return CrawlerID
}

func (t *Crawler) Name() string {
	return string(CrawlerID)
}

func (t *Crawler) Fetch(ctx context.Context, w io.Writer) error {
	request, _ := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://cprss.s3.amazonaws.com/golangweekly.com.xml", nil)
	return terrors.Wrap(t.fetcher.DoRequest(ctx, request, w))
}

func (t *Crawler) Parse(ctx context.Context, r io.Reader) ([]timeseriesdata.TimeSeriesData, error) {
	feed, err := t.fp.Parse(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	for _, item := range feed.Items {
		fmt.Println(item.Title, item.Link, item.Published)
		publishedAt, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.Published)
		if err != nil {
			clog.L.Errorf(ctx, "%+v\n", err)
			continue
		}
		data := timeseriesdata.TimeSeriesDataBlogFeed{
			ID: timeseriesdata.TimeSeriesDataID(
				strings.Replace(
					strings.Replace(item.GUID, ":", "-", -1),
					"/",
					"-",
					-1,
				),
			),
			PublishedAt: publishedAt,
			Title:       item.Title,
			URL:         item.Link,
		}
		returned = append(returned, &data)
	}
	return returned, nil
}

func (t *Crawler) Publish(ctx context.Context, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.repository.SetTimeSeriesData(ctx, CrawlerID, data...))
}
