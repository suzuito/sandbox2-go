package web

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPConnpass struct {
	Cli   *http.Client
	Query url.Values
	Days  int
}

func (t *FetcherHTTPConnpass) ID() crawler.FetcherID {
	return "fetcher_http_connpass"
}

func (t *FetcherHTTPConnpass) Do(ctx context.Context, w io.Writer, _ crawler.CrawlerInputData) error {
	u, _ := url.Parse("https://connpass.com/api/v1/event/")
	q := u.Query()
	for k, v := range t.Query {
		for _, vv := range v {
			q.Add(k, vv)
		}
	}
	d := time.Now()
	for i := 0; i < t.Days; i++ {
		q.Add("ymd", d.Add(time.Duration(i)*time.Hour*24).Format("20060102"))
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return terrors.Wrap(err)
	}
	res, err := t.Cli.Do(req)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		status := res.StatusCode
		return terrors.Wrapf("HTTP error : status=%d body=%s", status, body)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}

func NewFetcherHTTPConnpass(def *crawler.FetcherDefinition, args *factory.NewFuncFetcherArgument) (crawler.Fetcher, error) {
	days, err := crawler.GetFromArgumentDefinition[int](def.Argument, "Days")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	query, err := crawler.GetFromArgumentDefinition[url.Values](def.Argument, "Query")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &FetcherHTTPConnpass{
		Cli:   args.HTTPClient,
		Days:  days,
		Query: query,
	}, nil
}
