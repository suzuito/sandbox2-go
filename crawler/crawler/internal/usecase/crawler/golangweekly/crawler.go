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
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

const CrawlerID crawler.CrawlerID = "golangweekly"

type Crawler struct {
	repository repository.Repository
	cliHTTP    *http.Client
	fp         *gofeed.Parser
}

func NewCrawler(repository repository.Repository) crawler.Crawler {
	return &Crawler{
		repository: repository,
		cliHTTP:    http.DefaultClient,
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
	res, err := t.cliHTTP.Get("https://cprss.s3.amazonaws.com/golangweekly.com.xml")
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return terrors.Wrapf("HTTP error is occured code=%d", res.StatusCode)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrap(err)
	}
	return nil
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
