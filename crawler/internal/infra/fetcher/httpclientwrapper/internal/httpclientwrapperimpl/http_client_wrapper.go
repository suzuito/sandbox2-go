package httpclientwrapperimpl

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"slices"

	"github.com/suzuito/sandbox2-go/common/httpclientcache"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type HTTPClientWrapperImpl struct {
	Cli         *http.Client
	UseCache    bool
	Cache       httpclientcache.Client
	CacheOption *httpclientcache.ClientOption
}

func (t *HTTPClientWrapperImpl) Do(
	ctx context.Context,
	logger *slog.Logger,
	req *http.Request,
	w io.Writer,
	statusCodesSuccess []int,
) error {
	if t.UseCache {
		hit, err := t.Cache.Get(ctx, req, w, t.CacheOption)
		if err != nil {
			logger.WarnContext(ctx, "Failed to get cache", "err", err)
		} else if hit {
			return nil
		}
	}
	LogRequest(logger, req)
	res, err := t.Cli.Do(req)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if !slices.Contains(statusCodesSuccess, res.StatusCode) {
		return terrors.Wrapf("HTTP error : status=%d", res.StatusCode)
	}
	// resBody 不要？
	// resBody, err := io.Copy(w, res.Body)
	// t.Cache.Set(ctx, req, res.Header.Get("Content-Type"), w, t.CacheOption)
	// にしてもいい？
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	if _, err := w.Write(resBody); err != nil {
		return terrors.Wrap(err)
	}
	if t.UseCache {
		if err := t.Cache.Set(ctx, req, res.Header.Get("Content-Type"), bytes.NewBuffer(resBody), t.CacheOption); err != nil {
			logger.WarnContext(ctx, "Failed to set cache", "err", err)
		}
	}
	return nil
}
