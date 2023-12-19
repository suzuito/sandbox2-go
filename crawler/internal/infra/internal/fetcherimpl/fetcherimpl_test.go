package fetcherimpl

import (
	"bytes"
	"log/slog"
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
