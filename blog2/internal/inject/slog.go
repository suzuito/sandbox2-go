package inject

import (
	"log/slog"
	"os"
)

func newSlogHandlerText(
	l slog.Leveler,
) slog.Handler {
	return slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     l,
			AddSource: true,
		},
	)
}

func newSlogHandlerJSON(
	l slog.Leveler,
) slog.Handler {
	return slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     l,
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					a.Key = "severity"
				}
				return a
			},
		},
	)
}
