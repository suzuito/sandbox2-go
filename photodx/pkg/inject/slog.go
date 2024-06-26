package inject

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
	"github.com/suzuito/sandbox2-go/photodx/pkg/environment"
)

func NewSlogHandler(
	env *environment.Environment,
) slog.Handler {
	var level slog.Level
	if err := level.UnmarshalText([]byte(env.LogLevel)); err != nil {
		fmt.Printf("use LogLevel 'DEBUG' because cannot parse LOG_LEVEL : %s", env.LogLevel)
		level = slog.LevelDebug
	}
	if env.Env == "local" {
		return devslog.NewHandler(os.Stdout, &devslog.Options{
			HandlerOptions: &slog.HandlerOptions{
				Level:     level,
				AddSource: true,
			},
		})
	}
	return slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     level,
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
