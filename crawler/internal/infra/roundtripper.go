package infra

import (
	"log/slog"
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/roundtripperimpl/requestlog"
)

func NewRequestLogRoundTripper(
	logger *slog.Logger,
) http.RoundTripper {
	return &requestlog.RoundTripper{
		L:                logger,
		DefaultTransport: http.DefaultTransport,
	}
}
