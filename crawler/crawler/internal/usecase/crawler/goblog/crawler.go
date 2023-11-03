package goblog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

const CrawlerID crawler.CrawlerID = "goblog"
const baseURLGoBlog = "https://go.dev"

type Crawler struct {
	repository repository.Repository
	fetcher    fetcher.FetcherHTTP
	cliHTTP    *http.Client
}

func NewCrawler(
	repository repository.Repository,
	fetcher fetcher.FetcherHTTP,
) crawler.Crawler {
	return &Crawler{
		repository: repository,
		fetcher:    fetcher,
		cliHTTP:    http.DefaultClient,
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
		ctx, http.MethodGet, baseURLGoBlog+"/blog", nil)
	return terrors.Wrap(t.fetcher.DoRequest(ctx, request, w))
}

func (t *Crawler) Parse(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	doc.Find(".blogtitle").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		if title == "" {
			return
		}
		href, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}
		dateString := s.Find(".date").Text()
		if dateString == "" {
			return
		}
		blogURL := fmt.Sprintf("%s%s", baseURLGoBlog, href)
		publishedAt, err := time.Parse("2 January 2006", dateString)
		if err != nil {
			clog.L.Errorf(ctx, "%+v", terrors.Wrap(err))
			return
		}
		data := timeseriesdata.TimeSeriesDataBlogFeed{
			ID:          timeseriesdata.TimeSeriesDataID(fmt.Sprintf("goblog-%s", publishedAt.Format("2006-01-02"))),
			PublishedAt: publishedAt,
			Title:       title,
			URL:         blogURL,
		}
		returned = append(returned, &data)
	})
	return returned, nil
}

func (t *Crawler) Publish(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.repository.SetTimeSeriesData(ctx, CrawlerID, data...))
}
