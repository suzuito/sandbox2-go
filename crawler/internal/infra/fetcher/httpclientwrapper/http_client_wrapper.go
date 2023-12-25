package httpclientwrapper

import (
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper/internal/httpclientwrapperimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository"
)

type HTTPClientWrapper interface {
	Do(
		ctx context.Context,
		logger *slog.Logger,
		req *http.Request,
		w io.Writer,
		statusCodesSuccess []int,
	) error
}

func NewHTTPClientWrapper(cli *http.Client, cliCache repository.HTTPClientCacheRepository) HTTPClientWrapper {
	return &httpclientwrapperimpl.HTTPClientWrapperImpl{
		Cli:      cli,
		CliCache: cliCache,
	}
}
