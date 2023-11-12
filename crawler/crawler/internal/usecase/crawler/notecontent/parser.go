package notecontent

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

type Parser struct {
	filter func(article *note.TimeSeriesDataNoteArticle) bool
}

func (t *Parser) Do(ctx context.Context, r io.Reader, input crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	parser := note.Parser{}
	article, err := parser.Parse(ctx, r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	if t.filter(article) {
		returned = append(returned, article)
	}
	return returned, nil
}

func NewParser(
	filter func(article *note.TimeSeriesDataNoteArticle) bool,
) *Parser {
	return &Parser{
		filter: filter,
	}
}
