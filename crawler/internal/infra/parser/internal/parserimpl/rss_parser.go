package parserimpl

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type RSSParser struct {
	FP *gofeed.Parser
}

func (t *RSSParser) ID() crawler.ParserID {
	return "rss"
}

func (t *RSSParser) Do(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	feed, err := t.FP.Parse(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	for _, item := range feed.Items {
		publishedAt, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.Published)
		if err != nil {
			clog.L.Errorf(ctx, "%+v\n", err)
			continue
		}
		data := timeseriesdata.TimeSeriesDataBlogFeed{
			ID: timeseriesdata.TimeSeriesDataID(strings.Replace(
				strings.Replace(item.GUID, ":", "-", -1),
				"/",
				"-",
				-1,
			)),
			PublishedAt: publishedAt,
			Title:       item.Title,
			URL:         item.Link,
			Author: &timeseriesdata.TimeSeriesDataBlogFeedAuthor{
				Name:     "Golang Weekly",
				URL:      "https://golangweekly.com/",
				ImageURL: "https://golangweekly.com/images/gopher-keith-57.png",
			},
		}
		returned = append(returned, &data)
	}
	return returned, nil
}
