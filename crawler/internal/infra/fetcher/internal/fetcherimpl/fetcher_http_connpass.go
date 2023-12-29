package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPConnpass struct {
	Cli         httpclientwrapper.HTTPClientWrapper
	TimeNowFunc func() time.Time
	Query       url.Values
	Days        int
}

func (t *FetcherHTTPConnpass) ID() crawler.FetcherID {
	return "fetcher_http_connpass"
}

func (t *FetcherHTTPConnpass) Do(ctx context.Context, logger *slog.Logger, w io.Writer, _ crawler.CrawlerInputData) error {
	u, _ := url.Parse("https://connpass.com/api/v1/event/")
	q := u.Query()
	for k, v := range t.Query {
		for _, vv := range v {
			q.Add(k, vv)
		}
	}
	d := t.TimeNowFunc()
	for i := 0; i < t.Days; i++ {
		q.Add("ymd", d.Add(time.Duration(i)*time.Hour*24).Format("20060102"))
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return terrors.Wrap(err)
	}
	return terrors.Wrap(t.Cli.Do(ctx, logger, req, w, []int{http.StatusOK}))
}
