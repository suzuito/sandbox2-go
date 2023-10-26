package clog

import "context"

type Logger interface {
	Debugf(ctx context.Context, format string, v ...any)
	Infof(ctx context.Context, format string, v ...any)
	Errorf(ctx context.Context, format string, v ...any)
	AddKey(k ...string)
}

// L is default logger
var L Logger = NewLoggerJSON()
