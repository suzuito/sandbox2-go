package web

import (
	"context"
	"io"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type FetcherHTTP struct {
	cli *http.Client
}

func (t *FetcherHTTP) DoRequest(
	ctx context.Context,
	request *http.Request,
	w io.Writer,
) error {
	res, err := t.cli.Do(request)
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

func NewFetcherHTTP(cli *http.Client) *FetcherHTTP {
	return &FetcherHTTP{
		cli: cli,
	}
}
