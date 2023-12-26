package fetcherimpl

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newMockLogger() (*slog.Logger, *bytes.Buffer) {
	w := bytes.NewBufferString("")
	h := slog.NewTextHandler(
		w,
		&slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch a.Key {
				case slog.SourceKey, slog.TimeKey:
					return slog.Attr{}
				}
				return a
			},
		},
	)
	l := slog.New(h)
	return l, w
}

func assertLogString(t *testing.T, expected []string, logString string) {
	if logString == "" && len(expected) <= 0 {
		return
	}
	assert.Equal(
		t,
		strings.Join(expected, "\n")+"\n",
		logString,
	)
}

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
