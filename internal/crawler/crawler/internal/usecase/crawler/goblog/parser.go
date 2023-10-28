package goblog

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type Parser struct {
}

func (t *Parser) Parse(ctx context.Context, r io.Reader) ([]timeseriesdata.TimeSeriesData, error) {
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
