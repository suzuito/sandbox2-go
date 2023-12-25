package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPConnpass struct {
	Cli         *http.Client
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
	LogRequest(logger, req)
	res, err := t.Cli.Do(req)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		status := res.StatusCode
		return terrors.Wrapf("HTTP error : status=%d", status)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}
