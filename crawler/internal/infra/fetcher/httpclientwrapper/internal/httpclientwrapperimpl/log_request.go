package httpclientwrapperimpl

import (
	"log/slog"
	"net/http"
)

func LogRequest(logger *slog.Logger, request *http.Request) {
	logger.InfoContext(
		request.Context(),
		"",
		"fetcher",
		map[string]any{
			"request": map[string]any{
				"host":  request.URL.Host,
				"path":  request.URL.Path,
				"query": request.URL.Query(),
			},
		},
	)
}
