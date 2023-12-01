package requestlog

import (
	"log/slog"
	"net/http"
)

type RoundTripper struct {
	L                *slog.Logger
	DefaultTransport http.RoundTripper
}

func (t *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	t.L.InfoContext(
		req.Context(),
		"",
		"fetcher",
		map[string]any{
			"request": map[string]any{
				"host": req.URL.Host,
			},
		},
	)
	return t.DefaultTransport.RoundTrip(req)
}
