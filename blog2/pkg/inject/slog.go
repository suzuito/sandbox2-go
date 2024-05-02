package inject

import (
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
)

func newSlogHandlerText(
	l slog.Leveler,
) slog.Handler {
	return devslog.NewHandler(os.Stdout, &devslog.Options{
		HandlerOptions: &slog.HandlerOptions{
			Level:     l,
			AddSource: true,
		},
	})
	// return slog.NewTextHandler(
	// 	os.Stdout,
	// 	&slog.HandlerOptions{
	// 		Level:     l,
	// 		AddSource: true,
	// 	},
	// )
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
