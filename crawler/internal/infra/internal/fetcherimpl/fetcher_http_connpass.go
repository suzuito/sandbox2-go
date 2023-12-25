package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/fetchercache"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPConnpass struct {
	Cli                *http.Client
	FetcherCacheClient fetchercache.Client
	TimeNowFunc        func() time.Time
	Query              url.Values
	Days               int
	UseCache           bool
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
	var r io.Reader
	if t.UseCache {
		r, err = t.FetcherCacheClient.Do(ctx, req)
		if err != nil {
			return terrors.Wrap(err)
		}
	} else {
		res, err := t.Cli.Do(req)
		if err != nil {
			return terrors.Wrap(err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			status := res.StatusCode
			return terrors.Wrapf("HTTP error : status=%d", status)
		}
		r = res.Body
	}
	if _, err := io.Copy(w, r); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}

func NewFetcherHTTPConnpass(def *crawler.FetcherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
	f := FetcherHTTPConnpass{}
	if f.ID() != def.ID {
		return nil, factory.ErrNoMatchedFetcherID
	}
	days, err := argument.GetFromArgumentDefinition[int](def.Argument, "Days")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	query, err := argument.GetFromArgumentDefinition[url.Values](def.Argument, "Query")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	useCache, err := argument.GetFromArgumentDefinition[bool](def.Argument, "UseCache")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	f.TimeNowFunc = time.Now
	f.Days = days
	f.Query = query
	f.Cli = setting.FetcherFactorySetting.HTTPClient
	f.FetcherCacheClient = setting.FetcherFactorySetting.FetcherCacheClient
	f.UseCache = useCache
	return &f, nil
}
