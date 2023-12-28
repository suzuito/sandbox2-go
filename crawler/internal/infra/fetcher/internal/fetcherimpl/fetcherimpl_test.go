package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPClientWrapper struct {
	ExpectedMethod string
	ExpectedURL    *url.URL
	RealMethod     string
	RealURL        *url.URL
}

func (t *MockHTTPClientWrapper) Do(
	ctx context.Context,
	logger *slog.Logger,
	req *http.Request,
	w io.Writer,
	statusCodesSuccess []int,
) error {
	t.RealMethod = req.Method
	t.RealURL = req.URL
	return nil
}

func (t *MockHTTPClientWrapper) Assert(tt *testing.T) {
	assert.Equal(tt, t.ExpectedMethod, t.RealMethod)
	assert.Equal(tt, t.ExpectedURL.String(), t.RealURL.String())
}

func mustURL(raw string) *url.URL {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		panic(err)
	}
	return u
}
