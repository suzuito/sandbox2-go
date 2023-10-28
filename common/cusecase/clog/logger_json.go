package clog

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type LoggerJSON struct {
	keys  map[string]any
	log   *log.Logger
	level Level
}

func (t *LoggerJSON) Debugf(ctx context.Context, format string, v ...any) {
	if t.level > LevelDebug {
		return
	}
	t.event(ctx, "DEBUG", format, v...)
}

func (t *LoggerJSON) Infof(ctx context.Context, format string, v ...any) {
	if t.level > LevelInfo {
		return
	}
	t.event(ctx, "INFO", format, v...)
}

func (t *LoggerJSON) Warnf(ctx context.Context, format string, v ...any) {
	if t.level > LevelWarn {
		return
	}
	t.event(ctx, "WARN", format, v...)
}

func (t *LoggerJSON) Errorf(ctx context.Context, format string, v ...any) {
	if t.level > LevelError {
		return
	}
	t.event(ctx, "ERROR", format, v...)
}

func (t *LoggerJSON) event(
	ctx context.Context,
	severity string,
	format string,
	v ...any,
) {
	ev := map[string]any{}
	for k := range t.keys {
		setMap(ctx, k, &ev)
	}
	ev["severity"] = severity
	ev["message"] = fmt.Sprintf(format, v...)
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		ev["file"] = filename
		ev["line"] = line
	}
	traceInfos := []terrors.TraceInfo{}
	for _, vv := range v {
		terr, ok := vv.(terrors.TraceableError)
		if ok {
			traceInfos = append(traceInfos, terr.StackTrace()...)
		}
	}
	ev["traceInfos"] = traceInfos
	b, err := json.Marshal(ev)
	if err != nil {
		b = []byte("{}")
	}

	t.log.Printf(string(b))
}

func (t *LoggerJSON) AddKey(keys ...string) {
	for _, k := range keys {
		t.keys[k] = nil
	}
}

func (t *LoggerJSON) SetLevel(l Level) {
	t.level = l
}

func setMap(ctx context.Context, key string, m *map[string]any) {
	v := ctx.Value(key)
	if v != nil {
		(*m)[key] = v
	}
}

func NewLoggerJSON(keys ...string) *LoggerJSON {
	l := log.Default()
	l.SetFlags(0)
	ljson := LoggerJSON{
		log:   l,
		level: LevelDebug,
		keys:  map[string]any{},
	}
	for _, k := range keys {
		ljson.keys[k] = nil
	}
	return &ljson
}
