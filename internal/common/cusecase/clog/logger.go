package clog

import "context"

type Logger interface {
	Debugf(ctx context.Context, format string, v ...any)
	Infof(ctx context.Context, format string, v ...any)
	Errorf(ctx context.Context, format string, v ...any)
	AddKey(k ...string)
	SetLevel(l Level)
}

// A Level is the importance or severity of a log event. The higher the level, the more important or severe the event.
type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 1
	LevelError Level = 2
)

// L is default logger
var L Logger = NewLoggerJSON()
