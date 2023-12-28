package gcpimpl

import (
	"context"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/common/httpclientcache"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type Client struct {
	Cli      *storage.Client
	Bucket   string
	BasePath string
	NowFunc  func() time.Time
}

func (t *Client) getPath(key string) string {
	return filepath.Join(t.BasePath, key)
}

func (t *Client) Get(
	ctx context.Context,
	req *http.Request,
	w io.Writer,
	opt *httpclientcache.ClientOption,
) (bool, error) {
	key := opt.KeyGen.KeyStringHashHex(req)
	oh := t.Cli.Bucket(t.Bucket).Object(t.getPath(key))
	// Check expiration
	attrs, err := oh.Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, terrors.Wrap(err)
	}
	ttlInDays := 0
	ttlInDaysString, exists := attrs.Metadata["ttl_in_days"]
	if exists {
		ttlInDays, err = strconv.Atoi(ttlInDaysString)
		if err != nil {
			return false, terrors.Wrap(err)
		}
	}
	expirationTime := attrs.Created.AddDate(0, 0, ttlInDays)
	if t.NowFunc().After(expirationTime) {
		return false, nil
	}
	// Read body
	r, err := oh.NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, terrors.Wrap(err)
	}
	defer r.Close()
	if _, err := io.Copy(w, r); err != nil {
		return true, terrors.Wrap(err)
	}
	return true, nil
}

func (t *Client) Set(
	ctx context.Context,
	req *http.Request,
	contentType string,
	responseBody io.Reader,
	opt *httpclientcache.ClientOption,
) error {
	key := opt.KeyGen.KeyStringHashHex(req)
	oh := t.Cli.Bucket(t.Bucket).Object(t.getPath(key))
	w := oh.NewWriter(ctx)
	if _, err := io.Copy(w, responseBody); err != nil {
		return terrors.Wrap(err)
	}
	w.ContentType = contentType
	if err := w.Close(); err != nil {
		return terrors.Wrap(err)
	}
	attrsToUpdate := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"url":         req.URL.String(),
			"key":         opt.KeyGen.KeyString(req),
			"key_hash":    opt.KeyGen.KeyStringHashHex(req),
			"ttl_in_days": strconv.Itoa(opt.TTLInDays),
		},
	}
	if _, err := oh.Update(ctx, attrsToUpdate); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
