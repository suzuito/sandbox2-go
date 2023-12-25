package httpclientwrapperimpl

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"slices"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository"
)

type HTTPClientWrapperImpl struct {
	Cli      *http.Client
	CliCache repository.HTTPClientCacheRepository
}

func (t *HTTPClientWrapperImpl) Do(
	ctx context.Context,
	logger *slog.Logger,
	req *http.Request,
	w io.Writer,
	statusCodesSuccess []int,
) error {
	hit, cacheBytes, err := t.CliCache.Get(ctx, req)
	if err != nil {
		logger.WarnContext(ctx, "Failed to get cache: %+v", "err", err)
	} else {
		if hit {
			buf := bytes.NewBuffer(cacheBytes)
			_, err := io.Copy(w, buf)
			return terrors.Wrap(err)
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
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	if err := t.CliCache.Set(ctx, req, resBody); err != nil {
		logger.WarnContext(ctx, "Failed to set cache: %+v", "err", err)
	}
	if _, err := w.Write(resBody); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
