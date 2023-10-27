package goconnpass

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/internal/common/terrors"
)

type Fetcher struct {
	CliHTTP *http.Client
}

func (t *Fetcher) Fetch(ctx context.Context, w io.Writer) error {
	u, _ := url.Parse("https://connpass.com/api/v1/event/")
	q := u.Query()
	q.Add("keyword_or", "go言語")
	q.Add("keyword_or", "golang")
	q.Add("keyword_or", "gopher")
	d := time.Now()
	for i := 0; i < 30; i++ {
		q.Add("ymd", d.Add(time.Duration(i)*time.Hour*24).Format("20060102"))
	}
	q.Add("count", "100")
	u.RawQuery = q.Encode()
	res, err := t.CliHTTP.Get(u.String())
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return terrors.Wrapf("HTTP error is occured code=%d", res.StatusCode)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
