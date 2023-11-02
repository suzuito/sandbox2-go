package knowledgeworkblogs

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Parser struct {
	fp *gofeed.Parser
}

func (t *Parser) Parse(ctx context.Context, r io.Reader) ([]timeseriesdata.TimeSeriesData, error) {
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
