package knowledgeworkblogs

import (
	"context"
	"io"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type Fetcher struct {
	cliHTTP *http.Client
}

func (t *Fetcher) Fetch(ctx context.Context, w io.Writer) error {
	res, err := t.cliHTTP.Get("https://note.com/knowledgework/rss")
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
