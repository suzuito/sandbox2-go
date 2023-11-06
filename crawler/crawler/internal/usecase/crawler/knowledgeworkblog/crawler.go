package knowledgeworkblog

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

const CrawlerID crawler.CrawlerID = "knowledgeworkblog"

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

func (t *Crawler) Fetch(ctx context.Context, w io.Writer, input crawler.CrawlerInputData) error {
	urlString, exists := input["URL"]
	if !exists {
		return terrors.Wrapf("input[\"URL\"] not found in input")
	}
	u, err := url.Parse(urlString.(string))
	if err != nil {
		return terrors.Wrap(err)
	}
	request, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	return terrors.Wrap(t.fetcher.DoRequest(ctx, request, w))
}

func (t *Crawler) Parse(ctx context.Context, r io.Reader, input crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	parser := note.Parser{}
	article, err := parser.Parse(ctx, r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	hasGolangTag := slices.ContainsFunc(article.Tags, func(tag note.TimeSeriesDataNoteArticleTag) bool {
		name := strings.ToLower(tag.Name)
		if name == "go" || name == "golang" || name == "go言語" {
			return true
		}
		return false
	})
	returned := []timeseriesdata.TimeSeriesData{}
	if hasGolangTag {
		returned = append(returned, article)
	}
	return returned, nil
}

func (t *Crawler) Publish(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.repository.SetTimeSeriesData(ctx, CrawlerID, data...))
}
