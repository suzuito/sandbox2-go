package noterss

import (
	"context"
	"io"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

type Parser struct {
	fp *gofeed.Parser
}

func (t *Parser) Do(ctx context.Context, r io.Reader, input crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
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

func NewParser() crawler.Parser {
	return &Parser{
		fp: gofeed.NewParser(),
	}
}
